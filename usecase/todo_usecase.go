package usecase

import (
	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUseCase(todoRepo domain.TodoRepository) (domain.TodoUseCase, error) {
	return &todoUseCase{todoRepo: todoRepo}, nil
}

func (t *todoUseCase) GetAll() (*[]domain.Todo, error) {
	return t.todoRepo.GetAll()
}

func (t *todoUseCase) CreateTodo(todo *domain.Todo, ctx *fiber.Ctx) (*mongo.InsertOneResult, error) {
	return t.todoRepo.CreateTodo(todo, ctx)
}
func (t *todoUseCase) UpdateTodo(id string, todo *models.Todo, c *fiber.Ctx) error {
	return t.todoRepo.UpdateTodo(id, todo, c)
}
func (t *todoUseCase) DeleteTodo(id string, c *fiber.Ctx) error {
	return t.todoRepo.DeleteTodo(id, c)
}
