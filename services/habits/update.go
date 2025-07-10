package habits

import (
	"time"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/habit_checks"
)

func Update(habitID, userID uint, updated *models.Habit) error {
	habit, err := FindByID(habitID, userID)
	if err != nil {
		return err
	}

	today := time.Now().Weekday()
	if err := habit_checks.HandleToday(today, habit, updated, GetHabitDayValue); err != nil {
		return err
	}

	updateHabitFields(habit, updated)

	return config.DB.Save(habit).Error
}

func updateHabitFields(habit, updated *models.Habit) {
	habit.Title = updated.Title
	habit.Icon = updated.Icon
	habit.Sunday = updated.Sunday
	habit.Monday = updated.Monday
	habit.Tuesday = updated.Tuesday
	habit.Wednesday = updated.Wednesday
	habit.Thursday = updated.Thursday
	habit.Friday = updated.Friday
	habit.Saturday = updated.Saturday
}
