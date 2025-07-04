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
	const maxRetries = 30 // ~2.5 minutes with 5 second intervals
	retryCount := 0

	for {
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			fmt.Printf("Failed to open database: %v. Retrying in 5 seconds...\n", err)
			retryCount++
			if retryCount >= maxRetries {
				return nil, fmt.Errorf("max retries reached, failed to open database: %w", err)
			}
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			db.Close() // Close the connection before retrying
			fmt.Printf("Failed to ping database: %v. Retrying in 5 seconds...\n", err)
			retryCount++
			if retryCount >= maxRetries {
				return nil, fmt.Errorf("max retries reached, failed to ping database: %w", err)
			}
			time.Sleep(5 * time.Second)
			continue
		}

		// Connection successful, set connection pool settings
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)

		return db, nil
	}
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
