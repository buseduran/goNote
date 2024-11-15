package route

import "github.com/gofiber/fiber/v2"

func SetupRoutes(c *fiber.App) {
	publicRouter := c.Group("/api")
	NewTodoRouter(publicRouter)
	NewUserRouter(publicRouter)
	NewAssetRouter(publicRouter)
}
