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

	database.AutoMigrate(&models.User{}, &models.Game{}, &models.Player{}, &models.Gamecard{}, &models.Turn{}, &models.Action{})
	database.Exec("PRAGMA foreign_keys = ON;")

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

	var playerController controllers.PlayerController
	playerController.Database = database
	ginEngine.GET("/games/:game_id/players", playerController.ReadAll)
	ginEngine.POST("/games/:game_id/players", playerController.Create)
	ginEngine.GET("/games/:game_id/players/:player_id", playerController.Read)

	var gamecardController controllers.GamecardController
	gamecardController.Database = database
	ginEngine.GET("/games/:game_id/players/:player_id/gamecards", gamecardController.ReadAll)
	ginEngine.POST("/games/:game_id/players/:player_id/gamecards", gamecardController.Create)
	ginEngine.GET("/games/:game_id/players/:player_id/gamecards/:gamecard_id", gamecardController.Read)

	var turnController controllers.TurnController
	turnController.Database = database
	ginEngine.GET("/games/:game_id/turns", turnController.ReadAll)
	ginEngine.POST("/games/:game_id/turns", turnController.Create)
	ginEngine.GET("/games/:game_id/turns/:turn_id", turnController.Read)

	var actionController controllers.ActionController
	actionController.Database = database
	ginEngine.GET("/games/:game_id/turns/:turn_id/actions", actionController.ReadAll)
	ginEngine.POST("/games/:game_id/turns/:turn_id/actions", actionController.Create)
	ginEngine.GET("/games/:game_id/turns/:turn_id/actions/:action_id", actionController.Read)

	ginEngine.Run()
}
