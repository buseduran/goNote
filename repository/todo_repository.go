package repository

import (
	"context"
	"log"

	"github.com/buwud/goNote/db"
	"github.com/buwud/goNote/domain/models"
	"go.mongodb.org/mongo-driver/bson"
)

type taskRepository struct {
	db db.Database
}

func NewTodoRepository(database db.Database) *taskRepository {
	return &taskRepository{db: database}
}

func (t *taskRepository) GetAll() *[]models.Todo {
	var todos []models.Todo
	cursor, err := t.db.DB.Collection("todos").Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var todo models.Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	return &todos
}
