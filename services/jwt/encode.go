package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Encode(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": float64(userID),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
