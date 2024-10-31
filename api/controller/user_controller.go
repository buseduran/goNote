package controller

import (
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (userController *UserController) SignUp(c *fiber.Ctx) error {
	user := new(domain.UserSignup)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	result, err := userController.UserUseCase.SignUp(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func (userController *UserController) SignIn(c *fiber.Ctx) error {
	user := new(domain.UserSignin)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	err := userController.UserUseCase.SignIn(user, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON("logged in")
}

func (userController *UserController) SignOut(c *fiber.Ctx) error {
	userController.UserUseCase.SignOut(c)
	return c.Status(fiber.StatusOK).JSON("signed out")
}
