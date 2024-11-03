package domain

import (
	"time"

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
	UserID    primitive.ObjectID `json:"user_id"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// get todolist response
type UserTodo struct {
	Completed bool      `json:"completed"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (todo *Todo) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["_id"] = todo.ID
	response["completed"] = todo.Completed
	response["body"] = todo.Body
	response["user_iD"] = todo.UserID
	response["created_at"] = todo.CreatedAt
	response["updated_at"] = todo.UpdatedAt
	return response
}

type TodoRepository interface {
	GetAll() (*[]Todo, error)
	CreateTodo(todo *Todo) (*mongo.InsertOneResult, error)
	UpdateTodo(id string, todo *models.Todo) error
	DeleteTodo(id string, c *fiber.Ctx) error
	GetTodoByUserID(userID primitive.ObjectID) (*[]UserTodo, error)
}

type TodoUseCase interface {
	GetAll() (*[]Todo, error)
	CreateTodo(todo *Todo) (*mongo.InsertOneResult, error)
	UpdateTodo(id string, todo *models.Todo) error
	DeleteTodo(id string, c *fiber.Ctx) error
	GetTodoByUserID(userID primitive.ObjectID) (*[]UserTodo, error)
}
