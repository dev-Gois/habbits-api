package habit_checks

import (
	"errors"

	"github.com/dev-Gois/habbits-api/config"
)

func Check(habitCheckID uint, userID uint) error {
	habitCheck, err := Find(habitCheckID)
	if err != nil {
		return err
	}

	if habitCheck.Habit.UserID != userID {
		return errors.New("check-in não encontrado")
	}

	habitCheck.Done = !habitCheck.Done
	return config.DB.Save(habitCheck).Error
}
