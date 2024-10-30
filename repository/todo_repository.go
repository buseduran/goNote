package repository

import (
	"context"
	"log"

	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(collection *mongo.Collection) *todoRepository {
	return &todoRepository{collection: collection}
}

func (t *todoRepository) GetAll() (*[]domain.Todo, error) {
	var todos []domain.Todo
	cursor, err := t.collection.Find(context.Background(), bson.M{})
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

func (t *todoRepository) CreateTodo(todo *domain.Todo) (*mongo.InsertOneResult, error) {
	return t.collection.InsertOne(context.Background(), todo)
}

func (t *todoRepository) UpdateTodo(id string, todo *models.Todo) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"completed": todo.Completed,
		"body":      todo.Body,
	}}
	_, err = t.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (t *todoRepository) DeleteTodo(id string, c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid todo ID"})
	}
	filter := bson.M{"_id": objectID}
	_, err = t.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}
	return nil
}
