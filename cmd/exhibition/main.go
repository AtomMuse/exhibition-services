package main

import (
	"atommuse/backend/exhibition-service/handler/exhibihandler"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"atommuse/backend/exhibition-service/pkg/service"
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Load MongoDB connection string from environment variable
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set.")
	}
	log.Println(mongoURI)
	// Connect to MongoDB
	client, err := connectToMongoDB(mongoURI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// Check if the connection to MongoDB is successful
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	dbCollection := client.Database("atommuse").Collection("exhibition")

	repo := &exhibirepo.MongoDBRepository{Collection: dbCollection}
	useCase := &service.ExhibitionUseCase{Repository: repo}
	handler := &exhibihandler.Handler{UseCase: useCase}

	// Create a new router with CORS middleware
	r := mux.NewRouter()

	// CORS middleware configuration
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Use CORS middleware with your router
	corsHandler := handlers.CORS(headersOk, originsOk, methodsOk)(r)

	r.HandleFunc("/exhibitions", handler.GetAllExhibitions).Methods("GET")
	r.HandleFunc("/exhibition/{id}", handler.GetExhibitionHandler).Methods("GET")

	// Update to bind to all available interfaces
	server := &http.Server{
		Addr:    ":8080",
		Handler: corsHandler, // Use the router with CORS middleware
	}

	log.Println("Server started on :8080")
	log.Fatal(server.ListenAndServe())
}

func connectToMongoDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetTLSConfig(&tls.Config{}) // Add an empty TLS config

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		// Log the error and return it without disconnecting
		log.Printf("Error pinging MongoDB: %v\n", err)
		return nil, err
	}

	log.Printf("Connected to MongoDB at %s\n", uri)
	return client, nil
}
