package repository

import (
	"context"
	"errors"
	"time"

	"github.com/buwud/goNote/api/utils"
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Collection) *userRepository {
	return &userRepository{collection: database}
}

func (u *userRepository) SignUp(user *domain.UserSignup) error {
	newUser := domain.User{
		UserName:  user.Username,
		Password:  utils.GeneratePassword(user.Password),
		IsActive:  true,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//check username exist
	existingUser := domain.User{}
	filter := bson.M{"username": user.Username}
	err := u.collection.FindOne(context.Background(), filter).Decode(&existingUser)

	if err == nil {
		return errors.New("username is taken")
	}

	//insert into db
	_, err = u.collection.InsertOne(context.Background(), &newUser)
	if err != nil {
		return err
	}

	return nil
}
func (u *userRepository) SignIn(user *domain.UserSignin, c *fiber.Ctx) error {
	signedUser := domain.User{}

	//check if user exist
	filter := bson.M{"username": user.Username}
	err := u.collection.FindOne(context.Background(), filter).Decode(&signedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		return err
	}

	//compare passwords and generate jwt token
	if utils.ComparePassword(signedUser.Password, user.Password) {
		err = utils.GenerateToken(signedUser.ID, c)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

func (u *userRepository) SignOut(c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
		//HTTPOnly: true,
		Secure: true,
	}
	//delete from local storage
	//clear the header

	c.Locals("userID", nil)
	c.Cookie(&cookie)
}
