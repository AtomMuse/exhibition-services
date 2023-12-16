package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Replace the following connection string with your actual MongoDB Atlas connection string
	connectionString := "mongodb+srv://admin:root123456@cluster0.eshkyjb.mongodb.net/"

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Create a MongoDB client
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Atlas!")

	// Close the MongoDB client
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
