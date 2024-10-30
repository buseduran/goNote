package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
	})

	tokeny, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}
	return tokeny, nil
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
