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
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		"INSERT INTO game_sessions(user_id, score, game_mode) VALUES($1, $2, $3)",
		userID, score, "default",
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}