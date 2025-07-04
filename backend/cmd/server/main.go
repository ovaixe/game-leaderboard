package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/ovaixe/game-leaderboard/internal/config"
	"github.com/ovaixe/game-leaderboard/internal/controllers"
	"github.com/ovaixe/game-leaderboard/pkg/database"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize New Relic
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.NewRelicAppName),
		newrelic.ConfigLicense(cfg.NewRelicLicenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatalf("Failed to initialize New Relic: %v", err)
	}

	// Initialize database
	rawDB, err := database.InitDB(cfg.DBConfig)

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer rawDB.Close()

	// Wrap the DB with our interface
	db := database.NewDBWrapper(rawDB)

	// Initialize controllers
	leaderboardController := controllers.NewLeaderboardController(db)

	// Set up Gin
	router := gin.Default()

	// Use New Relic middleware
	router.Use(nrgin.Middleware(app))

	// Start periodic rank updates
	go func() {
		ticker := time.NewTicker(1 * time.Minute) // Update every 1 minute
		defer ticker.Stop()
		for range ticker.C {
			log.Println("Updating leaderboard ranks...")
			if err := leaderboardController.Service().UpdateRanks(); err != nil {
				log.Printf("Error updating ranks: %v", err)
			} else {
				log.Println("Leaderboard ranks updated successfully.")
			}
		}
	}()

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
