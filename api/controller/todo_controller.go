package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/buwud/goNote/domain"
)

type TodoController struct {
	TodoUseCase domain.TodoUseCase
}

func (taskController *TodoController) GetAll(c *fiber.Ctx) error {
	todos, err := taskController.TodoUseCase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}
