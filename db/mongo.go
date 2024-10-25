package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DB *mongo.Database
}
type Collection struct {
	TodoCollection *mongo.Collection
}

func GetConnection() Database {
	MONGODB_CONNECTION := os.Getenv(("MONGODB_CONNECTION"))
	clientOptions := options.Client().ApplyURI(MONGODB_CONNECTION)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(("Connected to MongoDB ATLAS!"))
	return Database{
		DB: client.Database("golang_db"),
	}
}

func GetTodoCollection() Collection {
	collection := GetConnection().DB.Collection("todos")
	return Collection{
		TodoCollection: collection,
	}
}
