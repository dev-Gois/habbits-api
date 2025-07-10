package habit_checks

import (
	"time"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
)

func DeleteIncomplete(habitID uint, date time.Time) error {
	return config.DB.Where("habit_id = ? AND done = ? AND date = ?", habitID, false, date).Delete(&models.HabitCheck{}).Error
}
