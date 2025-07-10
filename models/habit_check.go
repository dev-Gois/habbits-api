package models

import (
	"time"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type HabitCheck struct {
	gorm.Model
	HabitID uint      `json:"habit_id" gorm:"not null"`
	Habit   Habit     `gorm:"foreignKey:HabitID" validate:"-"`
	Done    bool      `json:"done"`
	Date    time.Time `json:"date" gorm:"type:date;not null" validate:"required"`
}

func (h *HabitCheck) Create() error {
	if err := validator.New().Struct(h); err != nil {
		return err
	}

	return config.DB.Create(h).Error
}
