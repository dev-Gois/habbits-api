package jwt

import (
	"errors"
	"os"
	"strconv"
	"github.com/golang-jwt/jwt/v5"
)

func Decode(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}