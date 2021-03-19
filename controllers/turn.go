package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerobugdebug/kaart-api/models"
	"gorm.io/gorm"
)

type TurnController struct {
	Database *gorm.DB
}

// GET /turns
// Read all turns
func (turnController TurnController) ReadAll(context *gin.Context) {
	var turns []models.Turn
	turnController.Database.Preload("Actions").Where("game_id = ?", context.Param("game_id")).Find(&turns)
	context.JSON(http.StatusOK, turns)
}

// POST /turns
// Create new turn
func (turnController TurnController) Create(context *gin.Context) {
	// Validate input
	var turn models.Turn
	if err := context.ShouldBindJSON(&turn); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create turn
	turn.TurnID = 0
	if err := turnController.Database.Create(&turn).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, turn)
}

// GET /turns/:id
// Read single turn
func (turnController TurnController) Read(context *gin.Context) {
	var turn models.Turn
	statement := turnController.Database.Preload("Actions").Where("game_id = ? AND turn_id = ?", context.Param("game_id"), context.Param("turn_id"))
	if err := statement.First(&turn).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, turn)
}
