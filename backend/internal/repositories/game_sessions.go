package repositories

import (
	"database/sql"
	"math/rand"
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

	gameMode := "solo"
	if rand.Float64() > 0.5 {
		gameMode = "team"
	}

	_, err = tx.Exec(
		"INSERT INTO game_sessions(user_id, score, game_mode) VALUES($1, $2, $3)",
		userID, score, gameMode,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
