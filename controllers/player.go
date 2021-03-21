package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zerobugdebug/kaart-api/models"
	"gorm.io/gorm"
)

type PlayerController struct {
	Database *gorm.DB
}

// GET /players
// Read all players
func (playerController PlayerController) ReadAll(context *gin.Context) {
	var players []models.Player
	playerController.Database.Preload("Gamecards").Preload("Actions").Where("game_id = ?", context.Param("game_id")).Find(&players)
	context.JSON(http.StatusOK, players)
}

// POST /players
// Create new player
func (playerController PlayerController) Create(context *gin.Context) {
	// Validate input
	var player models.Player
	if err := context.ShouldBindJSON(&player); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create player
	player.PlayerID = 0
	tmpGameID, err := strconv.Atoi(context.Param("game_id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	player.GameID = uint(tmpGameID)
	if err := playerController.Database.Create(&player).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, player)
}

// GET /games/:game_id/players/:player_id
// Read single player
func (playerController PlayerController) Read(context *gin.Context) {
	var player models.Player
	statement := playerController.Database.Preload("Gamecards").Preload("Actions").Where("game_id = ? AND player_id = ?", context.Param("game_id"), context.Param("player_id"))
	if err := statement.First(&player).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, player)
}
