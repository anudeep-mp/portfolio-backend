package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init() {

	//load env variables
	err := godotenv.Load()
	CheckError(err)

	//mongo db connection
	clientOption := options.Client().ApplyURI(os.Getenv("MONGO_DB_CONNECTION_STRING"))
	client, err := mongo.Connect(context.TODO(), clientOption)

	CheckError(err)

	fmt.Println("MongoDB Connection Succesful")

	Collection = (*mongo.Collection)(client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME")))
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
