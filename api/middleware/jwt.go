package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func JWTAuthMiddleware() fiber.Handler {
	// Retrieve the JWT_SECRET from environment variables.
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is required but not set.")
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secret),
		ContextKey:   "jwt", // Custom context key for JWT data
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	// Log the error for internal tracking.
	log.Println("JWT Error:", err)

	// Return status 401 and a generic error message.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   "Missing or malformed JWT",
	})
}
