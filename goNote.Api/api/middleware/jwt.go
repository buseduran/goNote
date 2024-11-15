package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func JWTProtected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token provided"})
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse JWT token with claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// Handle token parsing errors
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"redirect": true,
			"message":  "Unauthorized, please log in",
		})
	}

	// Extract claims from token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"redirect": true,
			"message":  "Unauthorized, please log in",
		})
	}

	// Extract user ID from claims
	id, ok := (*claims)["user_id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"redirect": true,
			"message":  "Unauthorized, please log in",
		})
	}

	// Convert the user ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"redirect": true,
			"message":  "Unauthorized, please log in",
		})
	}

	// Set the user in the context for access in later handlers
	c.Locals("userID", objectID)

	// Proceed to the next handler
	return c.Next()
}
