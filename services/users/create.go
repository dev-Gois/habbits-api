package users

import (
	"fmt"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/jwt"
)

func Create(user models.User) (models.User, string, error) {
	var existing models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return models.User{}, "", fmt.Errorf("Email já está sendo utilizado!")
	}

	if err := user.Create(); err != nil {
		return models.User{}, "", fmt.Errorf("Erro ao criar usuário: %w", err)
	}

	token, err := jwt.Encode(user.ID)
	if err != nil {
		return models.User{}, "", fmt.Errorf("Erro ao gerar token: %w", err)
	}

	return user, token, nil
}
