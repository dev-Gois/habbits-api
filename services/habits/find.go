package habits

import (
	"errors"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
)

func FindByID(habitID, userID uint) (*models.Habit, error) {
	var habit models.Habit
	if err := config.DB.Where("id = ? AND user_id = ? AND deleted_at IS NULL", habitID, userID).Preload("User").First(&habit).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("hábito não encontrado")
		}
		return nil, err
	}
	return &habit, nil
}
