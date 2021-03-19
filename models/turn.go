package models

import (
	"time"
)

type Turn struct {
	TurnID      uint      `json:"turn_id" gorm:"primary_key"`
	GameID      uint      `json:"game_id"`
	Description string    `json:"description"`
	Victory     uint      `json:"victory"`
	Actions     []Action  `json:"actions" gorm:"foreignKey:TurnID;references:TurnID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
