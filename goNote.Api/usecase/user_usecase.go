package usecase

import (
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) (domain.UserUseCase, error) {
	return &userUseCase{userRepo: userRepo}, nil
}

func (u *userUseCase) SignUp(user *domain.UserSignup) error {
	return u.userRepo.SignUp(user)
}

func (u *userUseCase) SignIn(user *domain.UserSignin, c *fiber.Ctx) error {
	return u.userRepo.SignIn(user, c)
}

func (u *userUseCase) SignOut(c *fiber.Ctx) {
	u.userRepo.SignOut(c)
}
