package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	ID        uint
	Name      string
	Position  string
	Shoot     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Player) TableName() string {
	return "players"
}
