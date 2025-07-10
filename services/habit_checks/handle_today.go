package habit_checks

import (
	"time"

	"github.com/dev-Gois/habbits-api/models"
)

func HandleToday(day time.Weekday, oldHabit, newHabit *models.Habit, getHabitDayValue func(habit *models.Habit, date time.Time) bool) error {
	wasEnabled := getHabitDayValue(oldHabit, time.Now())
	isEnabled := getHabitDayValue(newHabit, time.Now())

	todayDate := time.Now().Truncate(24 * time.Hour)

	switch {
	case !wasEnabled && isEnabled:
		return Create(oldHabit.ID)
	case wasEnabled && !isEnabled:
		return DeleteIncomplete(oldHabit.ID, todayDate)
	default:
		return nil
	}
}
