package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/buwud/goNote/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DB *mongo.Database
}
type Collection struct {
	TodoCollection *mongo.Collection
	UserCollection *mongo.Collection
}

var (
	dbInstance *Database
	dbOnce     sync.Once
)

func GetConnection() Database {
	dbOnce.Do(func() {
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
		dbInstance = &Database{
			DB: client.Database("golang_db"),
		}
	})
	return *dbInstance
}

func GetCollections() Collection {
	db := GetConnection().DB
	return Collection{
		TodoCollection: db.Collection(domain.CollectionTodo),
		UserCollection: db.Collection(domain.CollectionUser),
	}
}
