package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buwud/goNote/api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// omitempty --> if the value is wrong, just dont add it, mongo creates own id
type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

func main() {
	fmt.Println("Hello worldaaaa!")

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173/",
		AllowCredentials: true,
		AllowHeaders:     "Authorization, Content-Type, Accept",
	}))

	route.SetupRoutes(app)

	port := os.Getenv("PORT")

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
