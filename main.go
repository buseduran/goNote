package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/buwud/goNote/api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// omitempty --> if the value is wrong, just dont add it, mongo creates own id
type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello worldaaaa!")

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	MONGODB_CONNECTION := os.Getenv(("MONGODB_CONNECTION"))
	clientOptions := options.Client().ApplyURI(MONGODB_CONNECTION)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB ATLAS!")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	route.SetupRoutes(app)

	//app.Get("/api/todos", getTodos)
	//app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	port := os.Getenv("PORT")

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

//	func getTodos(c *fiber.Ctx) error {
//		var todos []Todo
//		//cursor is like a pointer to the result set
//		cursor, err := collection.Find(context.Background(), bson.M{})
//		if err != nil {
//			return err
//		}
//		//postpone the execution of a function until the surrounding one ends
//		defer cursor.Close(context.Background())
//		for cursor.Next(context.Background()) {
//			var todo Todo
//			if err := cursor.Decode(&todo); err != nil {
//				fmt.Printf("Decode error: %v\n", err)
//				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//					"error": "Failed to decode todo",
//				})
//			}
//			todos = append(todos, todo)
//		}
//		return c.JSON(todos)
//	}
// func createTodo(c *fiber.Ctx) error {
// 	todo := new(Todo)

// 	if err := c.BodyParser(todo); err != nil {
// 		return err
// 	}

// 	if todo.Body == "" {
// 		return c.Status(400).JSON(fiber.Map{"error": "todo body cannot be empty"})
// 	}
// 	result, err := collection.InsertOne(context.Background(), todo)

//		if err != nil {
//			return err
//		}
//		todo.ID = result.InsertedID.(primitive.ObjectID)
//		return c.Status(201).JSON(todo)
//	}
func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var completed struct {
		Completed bool `json:"completed"`
	}

	if err := c.BodyParser(&completed); err != nil {
		return err
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": completed.Completed}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid todo ID"})
	}
	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
