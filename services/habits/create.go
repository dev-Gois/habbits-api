package habits

import (
	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
)

func Create(habit *models.Habit, userID uint) error {
	habit.UserID = userID

	if err := habit.Create(); err != nil {
		return err
	}

	return config.DB.Preload("User").First(habit, habit.ID).Error
}
