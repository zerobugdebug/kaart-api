package models

import (
	"time"
)

type Player struct {
	PlayerID  uint       `json:"player_id" gorm:"primary_key"`
	GameID    uint       `json:"game_id"`
	UserID    uint       `json:"user_id" `
	User      User       `json:"user" gorm:"foreignKey:UserID"`
	Health    uint       `json:"health"`
	Power     uint       `json:"power"`
	Gamecards []Gamecard `json:"gamecards" gorm:"foreignKey:PlayerID;references:PlayerID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
