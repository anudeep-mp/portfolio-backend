package database

import (
	"context"
	"fmt"
	"os"

	"github.com/anudeep-mp/portfolio-backend/utilities"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init() {

	//load env variables
	err := godotenv.Load()
	utilities.CheckError(err)

	//mongo db connection
	clientOption := options.Client().ApplyURI(os.Getenv("MONGO_DB_CONNECTION_STRING"))
	client, err := mongo.Connect(context.TODO(), clientOption)

	utilities.CheckError(err)

	fmt.Println("MongoDB Connection Succesful")

	Collection = (*mongo.Collection)(client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME")))
}
