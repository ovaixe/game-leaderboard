package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ovaixe/game-leaderboard/pkg/database"
	"github.com/ovaixe/game-leaderboard/internal/models"
	"github.com/ovaixe/game-leaderboard/internal/services"
	"github.com/ovaixe/game-leaderboard/internal/utils"
)

type LeaderboardController struct {
	service *services.LeaderboardService
}

func NewLeaderboardController(db database.Database) *LeaderboardController {
	return &LeaderboardController{
		service: services.NewLeaderboardService(db),
	}
}

func (lc *LeaderboardController) SubmitScore(c *gin.Context) {
	var req models.ScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := lc.service.SubmitScore(req.UserID, req.Score); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, gin.H{"result": "success"})
}

func (lc *LeaderboardController) GetTopPlayers(c *gin.Context) {
	players, err := lc.service.GetTopPlayers(10)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, players)
}

func (lc *LeaderboardController) GetPlayerRank(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	entry, err := lc.service.GetPlayerRank(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Player not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, entry)
}