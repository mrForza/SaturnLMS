package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectToMongo(uri string) (*mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("Try to connect to MongoDB")
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("Connected to MongoDB")
	return client, cancel
}

func createCollections(client *mongo.Client) {
	db := client.Database("study_db")

	collections := []string{"course", "lesson", "homework"}

	for _, name := range collections {
		if err := db.CreateCollection(context.TODO(), name); err != nil {
			fmt.Printf("Collection %s may already exist: %v\n", name, err)
		} else {
			fmt.Printf("Created collection: %s\n", name)
		}
	}
}

func main() {
	mongoURI := "mongodb://admin:secret@studyService.mongo:27017"
	client, cancel := connectToMongo(mongoURI)
	defer cancel()
	defer client.Disconnect(context.TODO())
	createCollections(client)
}
