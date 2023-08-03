package database

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "test"
	collectionName = "tasks"
)

var collection *mongo.Collection

func Collection() *mongo.Collection {
	return collection
}

func Init() {
	err := godotenv.Load()
	checkError(err)
	// connectionString := os.Getenv("DB_STRING")
	connectionString := "mongodb+srv://chinmay:chinmay@tasks.rqbaptf.mongodb.net/?retryWrites=true&w=majority"
	println("hello world")
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkError(err)

	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance is ready")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
