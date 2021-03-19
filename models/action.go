package models

import (
	"time"
)

type Action struct {
	ActionID  uint      `json:"action_id" gorm:"primary_key"`
	TurnID    uint      `json:"turn_id"`
	PlayerID  uint      `json:"player_id"`
	Player    Player    `json:"player" gorm:"foreignKey:PlayerID"`
	Type      uint      `json:"type"`
	Boost     uint      `json:"boost"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
