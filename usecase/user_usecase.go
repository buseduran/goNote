package usecase

import (
	"github.com/buwud/goNote/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) (domain.UserUseCase, error) {
	return &userUseCase{userRepo: userRepo}, nil
}

func (u *userUseCase) SignUp(user *domain.UserSignup) (*mongo.InsertOneResult, error) {
	return u.userRepo.SignUp(user)
}

func (u *userUseCase) SignIn(user *domain.UserSignin) (string, error) {
	return u.userRepo.SignIn(user)
}
