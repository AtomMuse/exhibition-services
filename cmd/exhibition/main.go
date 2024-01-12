package main

import (
	"atommuse/backend/exhibition-service/handler/exhibihandler"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"atommuse/backend/exhibition-service/pkg/service"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:root123456@cluster0.eshkyjb.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
		return
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
		return
	}

	dbCollection := client.Database("atommuse").Collection("exhibition")

	repo := &exhibirepo.MongoDBRepository{Collection: dbCollection}
	useCase := &service.ExhibitionUseCase{Repository: repo}
	handler := &exhibihandler.Handler{UseCase: useCase}

	r := mux.NewRouter()
	r.HandleFunc("/exhibitions", handler.GetAllExhibitions).Methods("GET")
	r.HandleFunc("/exhibition/{id}", handler.GetExhibitionHandler).Methods("GET")

	http.Handle("/", r)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
