package models

import (
	"github.com/dev-Gois/habbits-api/config"
	"github.com/go-playground/validator/v10"
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
	if err := validator.New().Struct(h); err != nil {
		return err
	}

	return config.DB.Create(h).Error
}
