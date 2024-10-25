package controller

import (
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoController struct {
	TodoUseCase domain.TodoUseCase
}

func (todoController *TodoController) GetAll(c *fiber.Ctx) error {
	todos, err := todoController.TodoUseCase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func (todoController *TodoController) CreateTodo(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "todo body cannot be empty"})
	}

	result, err := todoController.TodoUseCase.CreateTodo(todo, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusOK).JSON(todo)
}
