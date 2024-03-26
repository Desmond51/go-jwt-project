package database

import (
	"context"
	"fmt"
	"go/mongo.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func DBinstance() *mongo.Client{
	err:= godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err =client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connect to MongoDB")
	return client
}

var Client *mongo.client = DBinstance()
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collecton
}