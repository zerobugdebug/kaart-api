package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerobugdebug/kaart-api/controllers"
	"github.com/zerobugdebug/kaart-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("kaart.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Game{})
	database.AutoMigrate(&models.Player{})

	return database
}

func main() {
	var database *gorm.DB
	ginEngine := gin.Default()

	database = ConnectDatabase()

	ginEngine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{"data": "hello world"})
	})

	var userController controllers.UserController
	userController.Database = database
	ginEngine.GET("/users", userController.ReadAll)
	ginEngine.POST("/users", userController.Create)
	ginEngine.GET("/users/:user_id", userController.Read)

	var gameController controllers.GameController
	gameController.Database = database
	ginEngine.GET("/games", gameController.ReadAll)
	ginEngine.POST("/games", gameController.Create)
	ginEngine.GET("/games/:game_id", gameController.Read)

	ginEngine.Run()
}
