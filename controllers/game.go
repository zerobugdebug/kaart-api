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
	//	gameController.Database.Find(&games)
	gameController.Database.Preload("Player").Find(&games)
	context.JSON(http.StatusOK, games)
}

// POST /books
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
	gameController.Database.Create(&game)

	context.JSON(http.StatusCreated, game)
}

// GET /games/:id
// Read single game
func (gameController GameController) Read(context *gin.Context) {
	var game models.Game

	if err := gameController.Database.Where("game_id = ?", context.Param("game_id")).First(&game).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, game)
}
