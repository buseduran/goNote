package domain

import (
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionTodo = "todos"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

func (todo *Todo) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["_id"] = todo.ID
	response["completed"] = todo.Completed
	response["body"] = todo.Body
	return response
}

type TodoRepository interface {
	GetAll() (*[]Todo, error)
	CreateTodo(todo *Todo) (*mongo.InsertOneResult, error)
	UpdateTodo(id string, todo *models.Todo) error
	DeleteTodo(id string, c *fiber.Ctx) error
}

type TodoUseCase interface {
	GetAll() (*[]Todo, error)
	CreateTodo(todo *Todo) (*mongo.InsertOneResult, error)
	UpdateTodo(id string, todo *models.Todo) error
	DeleteTodo(id string, c *fiber.Ctx) error
}
