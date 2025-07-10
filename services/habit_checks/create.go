package habit_checks

import (
	"time"

	"github.com/dev-Gois/habbits-api/models"
)

func Create(habitID uint) error {
	habitCheck := &models.HabitCheck{
		HabitID: habitID,
		Done:    false,
		Date:    time.Now(),
	}

	return habitCheck.Create()
}
