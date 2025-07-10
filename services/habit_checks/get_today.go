package habit_checks

import (
	"time"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
)

func GetAllToday(userID uint) ([]models.HabitCheck, error) {
	var habitChecks []models.HabitCheck
	today := time.Now().Format("2006-01-02")
	err := config.DB.
		Joins("JOIN habits ON habits.id = habit_checks.habit_id").
		Preload("Habit").
		Preload("Habit.User").
		Where("habits.user_id = ? AND habit_checks.date = ? AND habits.deleted_at IS NULL", userID, today).
		Find(&habitChecks).Error
	if err != nil {
		return nil, err
	}
	return habitChecks, nil
}
