package models

import (
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type GameSession struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Score     int       `json:"score"`
	GameMode  string    `json:"game_mode"`
	Timestamp time.Time `json:"timestamp"`
}

type LeaderboardEntry struct {
	UserID     int    `json:"user_id"`
	Username   string `json:"username"`
	TotalScore int    `json:"total_score"`
	Rank       int    `json:"rank"`
}

type ScoreRequest struct {
	UserID int `json:"user_id" binding:"required"`
	Score  int `json:"score" binding:"required"`
}
