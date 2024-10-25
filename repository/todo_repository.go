package repository

import (
	"context"
	"log"

	"github.com/buwud/goNote/db"
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	db db.Collection
}

func NewTodoRepository(database db.Collection) *todoRepository {
	return &todoRepository{db: database}
}

func (t *todoRepository) GetAll() (*[]domain.Todo, error) {
	var todos []domain.Todo
	cursor, err := t.db.TodoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var todo domain.Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	return &todos, nil
}

func (t *todoRepository) CreateTodo(todo *domain.Todo, c *fiber.Ctx) (*mongo.InsertOneResult, error) {
	return t.db.TodoCollection.InsertOne(context.Background(), todo)
}
