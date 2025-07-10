package habits

import (
	"errors"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
)

func Delete(habitID uint, userID uint) error {
	var habit models.Habit

	if err := config.DB.Where("id = ? AND user_id = ?", habitID, userID).First(&habit).Error; err != nil {
		if err.Error() == "record not found" {
			return errors.New("hábito não encontrado")
		}
		return err
	}

	return config.DB.Delete(&habit).Error
}
