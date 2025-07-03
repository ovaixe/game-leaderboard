package repositories

import (
	"database/sql"
	"sync"

	"github.com/ovaixe/game-leaderboard/internal/models"
)

type LeaderboardRepository struct {
	db        *sql.DB
	cacheLock sync.RWMutex
}

func NewLeaderboardRepository(db *sql.DB) *LeaderboardRepository {
	return &LeaderboardRepository{db: db}
}

func (r *LeaderboardRepository) SubmitScore(userID, score int) error {
	r.cacheLock.Lock()
	defer r.cacheLock.Unlock()

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

	return tx.Commit()
}

func (r *LeaderboardRepository) UpdateRanks() error {
	r.cacheLock.Lock()
	defer r.cacheLock.Unlock()

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

	return tx.Commit()
}

func (r *LeaderboardRepository) GetTopPlayers(limit int) ([]models.LeaderboardEntry, error) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()

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

	return entries, nil
}

func (r *LeaderboardRepository) GetPlayerRank(userID int) (*models.LeaderboardEntry, error) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()

	var entry models.LeaderboardEntry
	err := r.db.QueryRow(`
		SELECT u.id, u.username, l.total_score, l.rank
		FROM leaderboard l
		JOIN users u ON l.user_id = u.id
		WHERE l.user_id = $1
	`, userID).Scan(&entry.UserID, &entry.Username, &entry.TotalScore, &entry.Rank)

	if err != nil {
		return nil, err
	}

	return &entry, nil
}