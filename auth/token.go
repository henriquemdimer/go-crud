package auth

import (
	"fmt"
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

func VerifyToken(tk string) (int64, error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("error with signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return -1, fmt.Errorf("invalid token")
	}

	id := claims["id"].(float64)

	return int64(id), nil
}
