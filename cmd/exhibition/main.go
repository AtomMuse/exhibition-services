package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "atommuse/backend/exhibition-service/cmd/exhibition/doc"
	"atommuse/backend/exhibition-service/handler/exhibihandler"
	"atommuse/backend/exhibition-service/handler/sectionhandler"
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"atommuse/backend/exhibition-service/pkg/repositorty/sectionrepo"
	"atommuse/backend/exhibition-service/pkg/service/exhibisvc"
	"atommuse/backend/exhibition-service/pkg/service/sectionsvc"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Title						Exhibition Service API
// @Version					v0
// @Description				Exhibition Service สำหรับขอจัดการเกี่ยวกับ Exhibition ทั้งการสร้าง แก้ไข ลบ exhibition
// @Schemes					http
// @SecurityDefinitions.apikey	BearerAuth
// @In							header
// @Name						Authorization
func main() {
	initializeEnvironment()

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

	router := setupRouter(client)

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(router.Run(":8080"))
}

// initializeEnvironment initializes environment variables from .env file
func initializeEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

// connectToMongoDB connects to MongoDB
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

// authMiddleware is middleware to validate the token and check the role
func authMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		secretKey := os.Getenv("secret_key")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		// Parse the token
		claims := &model.JwtCustomClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			// Check the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret key for validation

			fmt.Println(secretKey)
			fmt.Println(claims)

			return []byte(secretKey), nil
		})
		// Handle token parsing errors
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			fmt.Println("Token parsing error:", err)
			return
		}

		// Check if the token is valid
		if !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			fmt.Println("Invalid token")
			return
		}

		// Check if the role admin
		if claims.Role == "admin" {
			c.Next()
		} else if claims.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			fmt.Println("Insufficient permissions")
			return
		}

		// Continue down the chain to handler etc
		c.Next()
	}
}

// setupRouter initializes the Gin router with routes and middleware
func setupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	// Initialize handlers and services
	exhibitionHandler := initExhibitionHandler(client)
	sectionHandler := initSectionHandler(client)

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Group routes
	api := router.Group("/api")
	{
		//Exhibitions
		api.GET("/exhibitions/all", authMiddleware("admin"), exhibitionHandler.GetAllExhibitions)
		api.GET("/exhibitions/:id", exhibitionHandler.GetExhibitionByID)
		api.GET("/exhibitions", exhibitionHandler.GetExhibitionsIsPublic)
		api.GET("/:userId/exhibitions", authMiddleware("exhibitor"), exhibitionHandler.GetExhibitionByUserID)
		api.POST("/exhibitions", authMiddleware("exhibitor"), exhibitionHandler.CreateExhibition)
		api.DELETE("/exhibitions/:id", authMiddleware("exhibitor"), exhibitionHandler.DeleteExhibition)
		api.PUT("/exhibitions/:id", authMiddleware("exhibitor"), exhibitionHandler.UpdateExhibition)
		//ExhibitionSections
		api.POST("/sections", authMiddleware("exhibitor"), sectionHandler.CreateExhibitionSection)
		api.DELETE("/sections/:id", authMiddleware("exhibitor"), sectionHandler.DeleteExhibitionSectionByID)
		api.GET("/sections/:id", authMiddleware("exhibitor"), sectionHandler.GetExhibitionSectionByID)
		api.GET("/sections/all", authMiddleware("admin"), sectionHandler.GetAllExhibitionSections)
		api.GET("/exhibitions/:id/sections", authMiddleware("admin"), sectionHandler.GetSectionsByExhibitionID)
		api.PUT("/sections/:id", authMiddleware("exhibitor"), sectionHandler.UpdateExhibitionSection)
		//like & Unlike
		api.PUT("/exhibitions/:id/like", authMiddleware("exhibitor"), exhibitionHandler.LikeExhibition)
		api.PUT("/exhibitions/:id/unlike", authMiddleware("exhibitor"), exhibitionHandler.UnlikeExhibition)

	}

	return router
}

// initExhibitionHandler initializes the exhibition handler with required dependencies
func initExhibitionHandler(client *mongo.Client) *exhibihandler.Handler {
	dbCollection := client.Database("atommuse").Collection("exhibitions")
	repo := &exhibirepo.ExhibitionRepository{Collection: dbCollection}
	service := &exhibisvc.ExhibitionServices{Repository: repo}
	return &exhibihandler.Handler{ExhibitionService: service}
}

// initSectionHandler initializes the section handler with required dependencies
func initSectionHandler(client *mongo.Client) *sectionhandler.Handler {
	dbCollection := client.Database("atommuse").Collection("exhibitionSections")
	repo := &sectionrepo.SectionRepository{Collection: dbCollection}
	service := &sectionsvc.SectionServices{Repository: repo}
	return &sectionhandler.Handler{SectionService: service}
}
