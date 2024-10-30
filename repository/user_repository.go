package repository

import (
	"context"
	"time"

	"github.com/buwud/goNote/api/utils"
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
	newUser.Password = utils.GeneratePassword(user.Password)
	newUser.IsActive = true
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	result, err := u.collection.InsertOne(context.Background(), &newUser)

	if err != nil {
		return nil, err
	}
	return result, nil
}
