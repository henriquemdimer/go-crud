package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
