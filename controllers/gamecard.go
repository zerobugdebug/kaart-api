package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerobugdebug/kaart-api/models"
	"gorm.io/gorm"
)

type GamecardController struct {
	Database *gorm.DB
}

// GET /gamecards
// Read all gamecards
func (gamecardController GamecardController) ReadAll(context *gin.Context) {
	var gamecards []models.Gamecard
	statement := gamecardController.Database.Joins("JOIN players ON players.player_id = game_cards.player_id")
	statement.Where("game_id = ? AND game_cards.player_id = ?", context.Param("game_id"), context.Param("player_id")).Find(&gamecards)
	context.JSON(http.StatusOK, gamecards)
}

// POST /gamecards
// Create new gamecard
func (gamecardController GamecardController) Create(context *gin.Context) {
	// Validate input
	var gamecard models.Gamecard
	if err := context.ShouldBindJSON(&gamecard); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create gamecard
	gamecard.GamecardID = 0
	if err := gamecardController.Database.Create(&gamecard).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gamecard)
}

// GET /gamecards/:id
// Read single gamecard
func (gamecardController GamecardController) Read(context *gin.Context) {
	var gamecard models.Gamecard
	statement := gamecardController.Database.Joins("JOIN players ON players.player_id = game_cards.player_id")
	statement = statement.Where("game_id = ? AND game_cards.player_id = ?", context.Param("game_id"), context.Param("player_id"))
	if err := statement.First(&gamecard).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, gamecard)
}
