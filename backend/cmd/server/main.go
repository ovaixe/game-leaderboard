package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/ovaixe/game-leaderboard/internal/config"
	"github.com/ovaixe/game-leaderboard/internal/controllers"
	"github.com/ovaixe/game-leaderboard/pkg/database"
)

func main() {
	time.Sleep(time.Second * 10)
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize database
	rawDB, err := database.InitDB(cfg.DBConfig)

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer rawDB.Close()

	// Wrap the DB with our interface
	db := database.NewDBWrapper(rawDB)

	// Create DB Tables and Populate data
	//db.CreateTables()
	//db.PopulateDB()

	// Initialize controllers
	leaderboardController := controllers.NewLeaderboardController(db)

	// Set up Gin
	router := gin.Default()

	// Set up routes
	api := router.Group("/api")
	{
		api.POST("/leaderboard/submit", leaderboardController.SubmitScore)
		api.GET("/leaderboard/top", leaderboardController.GetTopPlayers)
		api.GET("/leaderboard/rank/:user_id", leaderboardController.GetPlayerRank)
	}

	// Start server
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
