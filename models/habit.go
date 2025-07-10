package models

import (
	"github.com/dev-Gois/habbits-api/config"
	"gorm.io/gorm"
)

type Habit struct {
	gorm.Model
	Title     string       `json:"title" gorm:"not null" validate:"required"`
	Icon      string       `json:"icon" gorm:"not null" validate:"required"`
	Sunday    bool         `json:"sunday"`
	Monday    bool         `json:"monday"`
	Tuesday   bool         `json:"tuesday"`
	Wednesday bool         `json:"wednesday"`
	Thursday  bool         `json:"thursday"`
	Friday    bool         `json:"friday"`
	Saturday  bool         `json:"saturday"`
	UserID    uint         `json:"user_id" gorm:"not null"`
	User      User         `json:"user" gorm:"foreignKey:UserID"`
	Checks    []HabitCheck `json:"checks,omitempty" gorm:"foreignKey:HabitID"`
}

func (h *Habit) Create() error {
	return config.DB.Create(h).Error
}
