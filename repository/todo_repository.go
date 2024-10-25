package repository

import (
	"context"
	"log"

	"github.com/buwud/goNote/db"
	"github.com/buwud/goNote/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type todoRepository struct {
	db db.Database
}

func NewTodoRepository(database db.Database) *todoRepository {
	return &todoRepository{db: database}
}

func (t *todoRepository) GetAll() (*[]domain.Todo, error) {
	var todos []domain.Todo
	cursor, err := t.db.DB.Collection("todos").Find(context.Background(), bson.M{})
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
