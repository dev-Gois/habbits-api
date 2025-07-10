package habit_checks

import (
	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
)

func Find(habitCheckID uint) (*models.HabitCheck, error) {
	var habitCheck models.HabitCheck
	if err := config.DB.Preload("Habit").Where("id = ?", habitCheckID).First(&habitCheck).Error; err != nil {
		return nil, err
	}
	return &habitCheck, nil
}
