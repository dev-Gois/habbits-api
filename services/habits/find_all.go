package habits

import (
	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
)

func FindAllByUserID(userID uint) ([]models.Habit, error) {
	var habits []models.Habit

	err := config.DB.Where("user_id = ?", userID).Preload("User").Find(&habits).Error
	if err != nil {
		return nil, err
	}

	return habits, nil
}
