package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerobugdebug/kaart-api/models"
	"gorm.io/gorm"
)

type ActionController struct {
	Database *gorm.DB
}

// GET /actions
// Read all actions
func (actionController ActionController) ReadAll(context *gin.Context) {
	var actions []models.Action
	actionController.Database.Joins("JOIN turns ON turn.turn_id = actions.turn_id").Where("game_id = ? AND turns.turn_id = ?", context.Param("game_id"), context.Param("turn_id")).Find(&actions)
	context.JSON(http.StatusOK, actions)
}

// POST /actions
// Create new action
func (actionController ActionController) Create(context *gin.Context) {
	// Validate input
	var action models.Action
	if err := context.ShouldBindJSON(&action); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create action
	action.ActionID = 0
	if err := actionController.Database.Create(&action).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, action)
}

// GET /actions/:id
// Read single action
func (actionController ActionController) Read(context *gin.Context) {
	var action models.Action
	statement := actionController.Database.Joins("JOIN turns ON turn.turn_id = actions.turn_id")
	statement.Where("game_id = ? AND turns.turn_id = ? AND action_id = ?", context.Param("game_id"), context.Param("turn_id"), context.Param("game_card_id"))
	if err := statement.First(&action).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, action)
}
