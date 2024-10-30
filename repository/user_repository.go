package repository

import (
	"context"

	"github.com/buwud/goNote/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Collection) *userRepository {
	return &userRepository{collection: database}
}

func (u *userRepository) SignUp(user *domain.UserSignup) (*mongo.InsertOneResult, error) {
	newUser := domain.User{}
	newUser.UserName = user.Username
	newUser.Password = user.Password
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	result, err := u.collection.InsertOne(context.Background(), &newUser)

	if err != nil {
		return nil, err
	}
	return result, nil
}
