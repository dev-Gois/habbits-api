package users

import (
	"fmt"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/jwt"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Create(user models.User) (models.User, string, error) {
	var existing models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return models.User{}, "", fmt.Errorf("email already in use")
	}

	if err := validate.Struct(user); err != nil {
		return models.User{}, "", fmt.Errorf("validation error: %w", err)
	}

	if err := user.Create(); err != nil {
		return models.User{}, "", fmt.Errorf("failed to create user: %w", err)
	}

	token, err := jwt.Encode(user.ID)
	if err != nil {
		return models.User{}, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return user, token, nil
}