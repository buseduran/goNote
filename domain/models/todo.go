package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
