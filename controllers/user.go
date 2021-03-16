package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerobugdebug/kaart-api/models"
	"gorm.io/gorm"
)

type UserController struct {
	Database *gorm.DB
}

// GET /users
// Read all users
func (userController UserController) ReadAll(context *gin.Context) {
	var users []models.User
	userController.Database.Find(&users)

	context.JSON(http.StatusOK, users)
}

// POST /books
// Create new user
func (userController UserController) Create(context *gin.Context) {
	// Validate input
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user.UserID = 0
	userController.Database.Create(&user)

	context.JSON(http.StatusCreated, user)
}

// GET /users/:id
// Read single user
func (userController UserController) Read(context *gin.Context) {
	var user models.User

	if err := userController.Database.Where("user_id = ?", context.Param("user_id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, user)
}
