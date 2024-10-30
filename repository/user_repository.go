package repository

import (
	"context"
	"errors"
	"time"

	"github.com/buwud/goNote/api/utils"
	"github.com/buwud/goNote/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Collection) *userRepository {
	return &userRepository{collection: database}
}

func (u *userRepository) SignUp(user *domain.UserSignup) (*mongo.InsertOneResult, error) {
	newUser := domain.User{
		UserName:  user.Username,
		Password:  utils.GeneratePassword(user.Password),
		IsActive:  true,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := u.collection.InsertOne(context.Background(), &newUser)

	if err != nil {
		return nil, err
	}
	return result, nil
}
func (u *userRepository) SignIn(user *domain.UserSignin) (string, error) {
	signedUser := domain.User{}
	filter := bson.M{"username": user.Username}
	err := u.collection.FindOne(context.Background(), filter).Decode(&signedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		}
		return "", err
	}
	var token string
	if utils.ComparePassword(signedUser.Password, user.Password) {
		token, err = utils.GenerateToken(signedUser.ID)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return token, errors.New("incorrect password")
}
