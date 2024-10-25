package domain

import (
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

func (todo *Todo) ResponseMap() {
	response := make(map[string]interface{})
	response["_id"] = todo.ID
	response["completed"] = todo.Completed
	response["body"] = todo.Body
}

type TodoRepository interface {
	GetAll() (*[]Todo, error)
	CreateTodo(todo *Todo, c *fiber.Ctx) (*mongo.InsertOneResult, error)
}

type TodoUseCase interface {
	GetAll() (*[]Todo, error)
	CreateTodo(todo *Todo, c *fiber.Ctx) (*mongo.InsertOneResult, error)
}
