package models

import (
	"time"
)

type Gamecard struct {
	GamecardID uint      `json:"gamecard_id" gorm:"primary_key"`
	PlayerID   uint      `json:"player_id"`
	Damage     uint      `json:"damage"`
	Power      uint      `json:"power"`
	Playable   bool      `json:"playable"`
	Name       string    `json:"name"`
	Rarity     int       `json:"rarity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
