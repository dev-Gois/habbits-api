package models

import (
	"time"

	"gorm.io/gorm"
)

type HabitCheck struct {
	gorm.Model
	HabitID uint      `json:"habit_id" gorm:"not null"`
	Habit   Habit    `gorm:"foreignKey:HabitID"`
	Done    bool      `json:"done"`
	Date    time.Time `json:"date" gorm:"type:date;not null" validate:"required"`
}