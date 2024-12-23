package route

import (
	"github.com/buwud/goNote/api/controller"
	"github.com/buwud/goNote/db"
	"github.com/buwud/goNote/repository"
	"github.com/buwud/goNote/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewUserRouter(publicRouter fiber.Router) {
	userRepo := repository.NewUserRepository(db.GetCollections().UserCollection)
	userUseCase, err := usecase.NewUserUseCase(userRepo)
	if err != nil {
		publicRouter.Use(func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		})
		return
	}
	userController := &controller.UserController{
		UserUseCase: userUseCase,
	}

	publicRouter.Post("/login", userController.SignIn)
	publicRouter.Post("/register", userController.SignUp)
	publicRouter.Post("/logout", userController.SignOut)
}
