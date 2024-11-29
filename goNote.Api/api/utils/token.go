package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateToken(id primitive.ObjectID, c *fiber.Ctx) error {

	//generate jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id.Hex(),
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := claims.SignedString(secret)
	if err != nil {
		return err
	}

	//set jwt token in cookie
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 1),
		//HTTPOnly: true,
		Secure: true,
	}
	c.Cookie(&cookie)

	// Authentication successful, return token
	return nil
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
