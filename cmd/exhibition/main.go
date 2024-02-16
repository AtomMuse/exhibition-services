package main

import (
	"atommuse/backend/exhibition-service/handler/exhibihandler"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"atommuse/backend/exhibition-service/pkg/service"
	"context"
	"log"
	"os"
	"time"

	_ "atommuse/backend/exhibition-service/cmd/exhibition/doc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Exhibition Service API
// @version v0
// @description Exhibition Service สำหรับขอจัดการเกี่ยวกับ Exhibition ทั้งการสร้าง แก้ไข ลบ exhibition
// @schemes https http
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set.")
	}
	log.Println("MongoURI:", mongoURI)

	client, err := connectToMongoDB(mongoURI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	dbCollection := client.Database("atommuse").Collection("exhibition")
	repo := &exhibirepo.ExhibitionRepository{Collection: dbCollection}
	service := &service.ExhibitionServices{Repository: repo}
	handler := &exhibihandler.Handler{Service: service}

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	// Swagger documentation route
	router := gin.Default()
	router.Use(cors.New(config))

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api := router.Group("/api")
	{
		api.GET("/exhibitions", func(c *gin.Context) {
			handler.GetAllExhibitions(c)
		})
		api.GET("/exhibitions/:id", func(c *gin.Context) {
			handler.GetExhibitionByID(c)
		})
		api.POST("/exhibitions", func(c *gin.Context) {
			handler.CreateExhibition(c)
		})
		api.DELETE("/exhibitions/:id", func(c *gin.Context) {
			handler.DeleteExhibition(c)
		})
	}

	log.Println("Server started on :8080")
	log.Fatal(router.Run(":8080"))
}

func connectToMongoDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Printf("Connected to MongoDB at %s\n", uri)
	return client, nil
}
