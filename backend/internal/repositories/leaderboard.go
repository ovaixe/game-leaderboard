package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ovaixe/game-leaderboard/internal/models"
)

type LeaderboardRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewLeaderboardRepository(db *sql.DB, redisClient *redis.Client) *LeaderboardRepository {
	return &LeaderboardRepository{
		db:    db,
		redis: redisClient,
	}
}

func (r *LeaderboardRepository) SubmitScore(userID, score int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO leaderboard(user_id, total_score)
		VALUES($1, $2)
		ON CONFLICT (user_id) DO UPDATE
		SET total_score = leaderboard.total_score + EXCLUDED.total_score
	`, userID, score)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	// Invalidate Redis caches
	ctx := context.Background()
	r.redis.Del(ctx, "top_players") // Invalidate top players cache
	r.redis.Del(ctx, fmt.Sprintf("player_rank:%d", userID)) // Invalidate specific player rank

	return nil
}

func (r *LeaderboardRepository) UpdateRanks() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE leaderboard
		SET rank = ranked.rank
		FROM (
			SELECT user_id, RANK() OVER (ORDER BY total_score DESC) as rank
			FROM leaderboard
		) ranked
		WHERE leaderboard.user_id = ranked.user_id
	`)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	// Invalidate Redis caches after ranks are updated
	ctx := context.Background()
	r.redis.Del(ctx, "top_players") // Invalidate top players cache
	// No need to invalidate individual player ranks, as they will be re-fetched on demand

	return nil
}

func (r *LeaderboardRepository) GetTopPlayers(limit int) ([]models.LeaderboardEntry, error) {
	ctx := context.Background()

	// Try to get from Redis cache
	cached, err := r.redis.Get(ctx, "top_players").Bytes()
	if err == nil {
		var entries []models.LeaderboardEntry
		if err := json.Unmarshal(cached, &entries); err == nil {
			return entries, nil
		}
	}

	// If not in cache, fetch from DB
	rows, err := r.db.Query(`
		SELECT u.id, u.username, l.total_score, l.rank
		FROM leaderboard l
		JOIN users u ON l.user_id = u.id
		ORDER BY l.rank ASC
		LIMIT $1
	`, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.LeaderboardEntry
	for rows.Next() {
		var entry models.LeaderboardEntry
		if err := rows.Scan(&entry.UserID, &entry.Username, &entry.TotalScore, &entry.Rank); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	// Store in Redis cache
	serialized, err := json.Marshal(entries)
	if err == nil {
		r.redis.Set(ctx, "top_players", serialized, 5*time.Minute) // Cache for 5 minutes
	}

	return entries, nil
}

func (r *LeaderboardRepository) GetPlayerRank(userID int) (*models.LeaderboardEntry, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("player_rank:%d", userID)

	// Try to get from Redis cache
	cached, err := r.redis.Get(ctx, cacheKey).Bytes()
	if err == nil {
		var entry models.LeaderboardEntry
		if err := json.Unmarshal(cached, &entry); err == nil {
			return &entry, nil
		}
	}

	// If not in cache, fetch from DB
	var entry models.LeaderboardEntry
	err = r.db.QueryRow(`
		SELECT u.id, u.username, l.total_score, l.rank
		FROM leaderboard l
		JOIN users u ON l.user_id = u.id
		WHERE l.user_id = $1
	`, userID).Scan(&entry.UserID, &entry.Username, &entry.TotalScore, &entry.Rank)

	if err != nil {
		return nil, err
	}

	// Store in Redis cache
	serialized, err := json.Marshal(entry)
	if err == nil {
		r.redis.Set(ctx, cacheKey, serialized, 1*time.Minute) // Cache for 1 minute
	}

	return &entry, nil
}