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
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	UserName  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Todos     []Todo             `json:"notes,omitempty" bson:"notes,omitempty"`
	IsActive  bool               `json:"is_active" bson:"is_active"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UserSignup struct {
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
}

type UserSignin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["_id"] = user.ID
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
	SignIn(user *UserSignin) (string, error)
}

type UserUseCase interface {
	SignUp(user *UserSignup) (*mongo.InsertOneResult, error)
	SignIn(user *UserSignin) (string, error)
}
