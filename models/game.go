package models

import (
	"time"
)

type Game struct {
	GameID      uint      `json:"game_id" gorm:"primary_key"`
	DisplayID   string    `json:"display_id"`
	CurrentTurn uint      `json:"current_turn"`
	Location    string    `json:"location"`
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	Players     []Player  `json:"players" gorm:"foreignKey:GameID;references:GameID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
