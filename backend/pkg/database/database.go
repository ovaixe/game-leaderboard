package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/ovaixe/game-leaderboard/internal/config"
)

func InitDB(cfg config.DBConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	var db *sql.DB
	var err error
	for {
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			fmt.Printf("Failed to open database: %v. Retrying in 5 seconds...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			db.Close() // Close the connection before retrying
			fmt.Printf("Failed to ping database: %v. Retrying in 5 seconds...", err)
			time.Sleep(5 * time.Second)
			continue
		}
		return db, nil // Successfully connected and pinged
	}
	return nil, fmt.Errorf("failed to connect to database after multiple retries: %w", err)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

// Database interface wraps the common database operations
type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Begin() (*sql.Tx, error)
	DB() *sql.DB
}

// DBWrapper implements Database interface and holds *sql.DB
type DBWrapper struct {
	db *sql.DB
}

func NewDBWrapper(db *sql.DB) *DBWrapper {
	return &DBWrapper{db: db}
}

func (w *DBWrapper) CreateTables() {
	w.Exec("CREATE TABLE users (id SERIAL PRIMARY KEY, username VARCHAR(255) UNIQUE NOT NULL, join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	w.Exec("CREATE TABLE game_sessions (id SERIAL PRIMARY KEY, user_id INT REFERENCES users(id) ON DELETE CASCADE, score INT NOT NULL, game_mode VARCHAR(50) NOT NULL, timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	w.Exec("CREATE TABLE leaderboard (id SERIAL PRIMARY KEY, user_id INT REFERENCES users(id) ON DELETE CASCADE, total_score INT NOT NULL, rank INT)")
}

func (w *DBWrapper) PopulateDB() {
	w.Exec("INSERT INTO users (username) SELECT 'user' || generate_series(1, 1000000)")
	w.Exec("INSERT INTO game_sessions (user_id, score, game_mode, timestamp) SELECT floor(random() * 1000000 + 1)::int, floor(random() * 10000 + 1)::int, CASE WHEN random() > 0.5 THEN 'solo' ELSE 'team' END, NOW() - INTERVAL '1 day' * floor(random() * 365) FROM generate_series(1, 5000000)")
	w.Exec("INSERT INTO leaderboard (user_id, total_score, rank) SELECT user_id, AVG(score) as total_score, RANK() OVER (ORDER BY SUM(score) DESC) FROM game_sessions GROUP BY user_id")
}

func (w *DBWrapper) Exec(query string, args ...interface{}) (sql.Result, error) {
	return w.db.Exec(query, args...)
}

func (w *DBWrapper) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return w.db.Query(query, args...)
}

func (w *DBWrapper) QueryRow(query string, args ...interface{}) *sql.Row {
	return w.db.QueryRow(query, args...)
}

func (w *DBWrapper) Begin() (*sql.Tx, error) {
	return w.db.Begin()
}

func (w *DBWrapper) DB() *sql.DB {
	return w.db
}
