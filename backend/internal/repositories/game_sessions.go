package repositories

import (
	"database/sql"
)

type GameSessionRepository struct {
	db *sql.DB
}

func NewGameSessionRepository(db *sql.DB) *GameSessionRepository {
	return &GameSessionRepository{db: db}
}

func (r *GameSessionRepository) RecordSession(userID, score int) error {
	_, err := r.db.Exec(
		"INSERT INTO game_sessions(user_id, score, game_mode) VALUES($1, $2, $3)",
		userID, score, "default",
	)
	return err
}