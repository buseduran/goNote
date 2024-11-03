package controller

import (
	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
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

func (todoController *TodoController) GetTodoByUserID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(primitive.ObjectID)
	todos, err := todoController.TodoUseCase.GetTodoByUserID(userID)
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

	userID := c.Locals("userID").(primitive.ObjectID)
	todo.UserID = userID

	result, err := todoController.TodoUseCase.CreateTodo(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusOK).JSON(todo)
}

func (todoController *TodoController) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "todo body cannot be empty"})
	}
	err := todoController.TodoUseCase.UpdateTodo(id, todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(todo)
}
func (todoController *TodoController) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := todoController.TodoUseCase.DeleteTodo(id, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(id)
}
