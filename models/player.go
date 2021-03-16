package models

import (
	"time"
)

type Player struct {
	PlayerID  uint `json:"player_id" gorm:"primary_key"`
	GameID    uint
	Health    uint      `json:"health"`
	Power     uint      `json:"power"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
