package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	UserName  string             `json:"username"`
	Password  string             `json:"password"`
	Todos     []Todo             `json:"todos,omitempty" bson:"todos,omitempty"`
	IsActive  bool               `json:"is_active"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}

type UserSignup struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `form:"username" binding:"required"`
	Password  string             `form:"password" binding:"required"`
	FirstName string             `form:"firstname"`
	LastName  string             `form:"lastname"`
}

func (user *User) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["id"] = user.ID
	response["first_name"] = user.FirstName
	response["last_name"] = user.LastName
	response["username"] = user.UserName
	response["todos"] = user.Todos
	response["is_active"] = user.IsActive
	response["created_at"] = user.CreatedAt
	response["updated_at"] = user.UpdatedAt
	return response
}

type UserRepository interface {
	SignUp(user *UserSignup) (*mongo.InsertOneResult, error)
}

type UserUseCase interface {
	SignUp(user *UserSignup) (*mongo.InsertOneResult, error)
}
