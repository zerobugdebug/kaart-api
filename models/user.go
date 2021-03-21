package models

import (
	"time"
)

type User struct {
	UserID    uint      `json:"user_id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Players   []Player  `json:"players" gorm:"foreignKey:UserID;references:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
