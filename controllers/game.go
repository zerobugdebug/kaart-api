package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerobugdebug/kaart-api/models"
	"gorm.io/gorm"
)

type GameController struct {
	Database *gorm.DB
}

// GET /games
// Read all games
func (gameController GameController) ReadAll(context *gin.Context) {
	var games []models.Game
	gameController.Database.Preload("Players").Preload("Turns").Find(&games)
	context.JSON(http.StatusOK, games)
}

// POST /games
// Create new game
func (gameController GameController) Create(context *gin.Context) {
	// Validate input
	var game models.Game
	if err := context.ShouldBindJSON(&game); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create game
	game.GameID = 0
	if err := gameController.Database.Create(&game).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, game)
}

// GET /games/:id
// Read single game
func (gameController GameController) Read(context *gin.Context) {
	var game models.Game
	if err := gameController.Database.Preload("Players").Preload("Turns").First(&game, context.Param("game_id")).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, game)
}
