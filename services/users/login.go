package users

import (
	"fmt"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/jwt"
)

func Login(user models.User) (string, error) {
	var existing models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existing).Error; err != nil {
		return "", fmt.Errorf("Usuário não encontrado!")
	}

	if !existing.CheckPassword(user.PlainPassword) {
		return "", fmt.Errorf("Senha incorreta!")
	}

	token, err := jwt.Encode(existing.ID)
	if err != nil {
		return "", fmt.Errorf("Erro ao gerar token!")
	}
	return token, nil
}