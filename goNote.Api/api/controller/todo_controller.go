package controller

import (
	responser "github.com/buwud/goNote/api/errors"
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
		return responser.FetchFailed(c)
	}

	return responser.FetchSuccess(c, todos)
}

func (todoController *TodoController) GetTodoByUserID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(primitive.ObjectID)
	todos, err := todoController.TodoUseCase.GetTodoByUserID(userID)
	if err != nil {
		return responser.FetchFailed(c)
	}
	return responser.FetchSuccess(c, todos)
}

func (todoController *TodoController) CreateTodo(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	if err := c.BodyParser(todo); err != nil {
		return responser.InvalidBody(c)
	}
	if todo.Body == "" {
		return responser.NotEmpty(c)
	}

	userID := c.Locals("userID").(primitive.ObjectID)
	todo.UserID = userID

	result, err := todoController.TodoUseCase.CreateTodo(todo)
	if err != nil {
		return responser.CreateFailed(c)
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return responser.CreateSuccess(c)
}

func (todoController *TodoController) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return responser.InvalidBody(c)
	}
	if todo.Body == "" {
		return responser.NotEmpty(c)
	}
	err := todoController.TodoUseCase.UpdateTodo(id, todo)
	if err != nil {
		return responser.UpdateFailed(c)
	}
	return responser.UpdateSuccess(c)
}
func (todoController *TodoController) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := todoController.TodoUseCase.DeleteTodo(id, c)
	if err != nil {
		return responser.DeleteFailed(c)
	}
	return responser.DeleteSuccess(c)
}
