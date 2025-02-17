package usecase

import (
	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (t *todoUseCase) GetTodoByUserID(userID primitive.ObjectID) (*[]domain.UserTodo, error) {
	return t.todoRepo.GetTodoByUserID(userID)
}

func (t *todoUseCase) CreateTodo(todo *domain.Todo) (*mongo.InsertOneResult, error) {
	return t.todoRepo.CreateTodo(todo)
}
func (t *todoUseCase) UpdateTodo(id string, todo *models.Todo) error {
	return t.todoRepo.UpdateTodo(id, todo)
}
func (t *todoUseCase) DeleteTodo(id string, c *fiber.Ctx) error {
	return t.todoRepo.DeleteTodo(id, c)
}
