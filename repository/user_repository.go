package repository

import (
	"context"

	"github.com/buwud/goNote/db"
	"github.com/buwud/goNote/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db db.Collection
}

func NewUserRepository(database db.Collection) *userRepository {
	return &userRepository{db: database}
}

func (u *userRepository) SignUp(user *domain.UserSignup) (*mongo.InsertOneResult, error) {
	newUser := domain.User{}
	newUser.UserName = user.Username
	newUser.Password = user.Password
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	result, err := u.db.TodoCollection.InsertOne(context.Background(), &newUser)

	if err != nil {
		return nil, err
	}
	return result, nil
}
