package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Todos     []Todo             `json:"todos,omitempty" bson:"_id,omitempty"`
	IsActive  bool               `json:"is_active"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
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
	//signup, signin, updateprofile
}

type UserUseCase interface {
	//signup, signin, updateprofile
}
