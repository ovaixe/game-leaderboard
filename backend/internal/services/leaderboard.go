package services

import (
	"log"

	"github.com/ovaixe/game-leaderboard/internal/models"
	"github.com/ovaixe/game-leaderboard/internal/repositories"
	"github.com/ovaixe/game-leaderboard/pkg/database"
	"github.com/ovaixe/game-leaderboard/pkg/redis"
)

type LeaderboardService struct {
	gameSessionRepo *repositories.GameSessionRepository
	leaderboardRepo *repositories.LeaderboardRepository
}

func NewLeaderboardService(db database.Database) *LeaderboardService {
	// Get the concrete *sql.DB instance
	sqlDB := db.DB()

	// Initialize Redis client
	redisClient, err := redis.NewRedisClient()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &LeaderboardService{
		gameSessionRepo: repositories.NewGameSessionRepository(sqlDB),
		leaderboardRepo: repositories.NewLeaderboardRepository(sqlDB, redisClient.Client),
	}
}

func (s *LeaderboardService) SubmitScore(userID, score int) error {
	if err := s.gameSessionRepo.RecordSession(userID, score); err != nil {
		return err
	}

	if err := s.leaderboardRepo.SubmitScore(userID, score); err != nil {
		return err
	}

	go s.leaderboardRepo.UpdateRanks()
	return nil
}

func (s *LeaderboardService) GetTopPlayers(limit int) ([]models.LeaderboardEntry, error) {
	return s.leaderboardRepo.GetTopPlayers(limit)
}

func (s *LeaderboardService) GetPlayerRank(userID int) (*models.LeaderboardEntry, error) {
	return s.leaderboardRepo.GetPlayerRank(userID)
}