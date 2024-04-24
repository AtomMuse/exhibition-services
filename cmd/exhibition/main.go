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
	"atommuse/backend/exhibition-service/handler/roomhandler"
	"atommuse/backend/exhibition-service/handler/sectionhandler"
	"atommuse/backend/exhibition-service/pkg/model"
	roomrepo "atommuse/backend/exhibition-service/pkg/repositorty/Roomrepo"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"atommuse/backend/exhibition-service/pkg/repositorty/sectionrepo"
	"atommuse/backend/exhibition-service/pkg/service/exhibisvc"
	"atommuse/backend/exhibition-service/pkg/service/roomsvc"
	"atommuse/backend/exhibition-service/pkg/service/sectionsvc"
	"atommuse/backend/exhibition-service/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
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

	client, err := utils.ConnectToMongoDB(mongoURI)
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

	// Start the HTTP server in a goroutine
	go func() {
		log.Println("Server started on :8080")
		log.Fatal(router.Run(":8080"))
	}()

	// Set the time zone to Asia/Bangkok (UTC+7)
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal(err)
	}

	// Run the scheduler
	for {
		// Get the current time in the specified time zone
		currentTime := time.Now().In(location)

		// Parse the current time into string
		currentTimeString := currentTime.Format(time.RFC3339)
		fmt.Println("time now:", currentTimeString)

		// Define the database and collection
		db := client.Database("atommuse")
		collection := db.Collection("exhibitions")

		// Find exhibitions where endDate has passed
		filterEndDate := bson.M{"endDate": bson.M{"$lt": currentTimeString}}
		updateEndDate := bson.M{"$set": bson.M{"isPublic": false}}
		updateResultEndDate, err := collection.UpdateMany(context.Background(), filterEndDate, updateEndDate)
		if err != nil {
			log.Println("Error updating endDate:", err)
		} else {
			log.Printf("Updated %d documents where endDate has passed", updateResultEndDate.ModifiedCount)
		}

		// Find exhibitions where startDate is equal to the current time
		// filterStartDate := bson.M{"startDate": currentTimeString}
		// updateStartDate := bson.M{"$set": bson.M{"isPublic": true}}
		// updateResultStartDate, err := collection.UpdateMany(context.Background(), filterStartDate, updateStartDate)
		// if err != nil {
		// 	log.Println("Error updating startDate:", err)
		// } else {
		// 	log.Printf("Updated %d documents where startDate is equal to current time", updateResultStartDate.ModifiedCount)
		// }

		// Sleep for a period before running the scheduler again
		time.Sleep(4 * time.Hour) // Adjust the duration as needed
	}
}

// initializeEnvironment initializes environment variables from .env file
func initializeEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

// authMiddleware is middleware to validate the token and check the role
func authMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// Check if a token is provided
		if token != "" {
			// Token provided, perform authentication
			secretKey := os.Getenv("secret_key")

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

			// Set user ID in context
			c.Set("user_id", claims.ID)
			c.Set("user_first_name", claims.FirstName)
			c.Set("user_last_name", claims.LastName)
			c.Set("user_image", claims.ProfileImage)
			c.Set("user_username", claims.UserName)

			fmt.Println("User ID:", claims.ID)

			// Check if the role admin
			if claims.Role == "admin" {
				c.Next()
			} else if claims.Role == "exhibitor" && role != "admin" {
				c.Next()
			} else if claims.Role != role {
				c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
				c.Abort()
				fmt.Println("Insufficient permissions")
				return
			}

			// Continue down the chain to handler etc
			c.Next()
		} else {
			// No token provided, just check the role
			if role == "" {
				// No role specified, allow the request to proceed
				c.Next()
				return
			}

			// Return an error for missing token
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			fmt.Println("Authorization token is required")
			return
		}
	}
}

// setupRouter initializes the Gin router with routes and middleware
func setupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // Replace "*" with allowed origins
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	// Initialize handlers and services
	exhibitionHandler := initExhibitionHandler(client)
	sectionHandler := initSectionHandler(client)
	roomHandler := initRoomHandler(client)

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Group routes
	api := router.Group("/api")
	{
		//Exhibitions
		api.GET("/exhibitions/all", authMiddleware("admin"), exhibitionHandler.GetAllExhibitions)
		api.GET("/exhibitions/:id", authMiddleware(""), exhibitionHandler.GetExhibitionByID)
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
		api.GET("/exhibitions/:id/sections", authMiddleware("exhibitor"), sectionHandler.GetSectionsByExhibitionID)
		api.PUT("/sections/:id", authMiddleware("exhibitor"), sectionHandler.UpdateExhibitionSection)
		//Rooms
		api.POST("/rooms", authMiddleware("exhibitor"), roomHandler.CreateExhibitionRoom)
		api.DELETE("/rooms/:id", authMiddleware("exhibitor"), roomHandler.DeleteExhibitionRoomByID)
		api.GET("/rooms/:id", authMiddleware("exhibitor"), roomHandler.GetExhibitionRoomByID)
		api.GET("/rooms/all", authMiddleware("admin"), roomHandler.GetAllExhibitionRooms)
		api.GET("/exhibitions/:id/rooms", authMiddleware("exhibitor"), roomHandler.GetRoomsByExhibitionID)
		api.PUT("/rooms/:id", authMiddleware("exhibitor"), roomHandler.UpdateExhibitionRoom)
		//like & Unlike
		api.PUT("/exhibitions/:id/like", authMiddleware("exhibitor"), exhibitionHandler.LikeExhibition)
		api.PUT("/exhibitions/:id/unlike", authMiddleware("exhibitor"), exhibitionHandler.UnlikeExhibition)
		//ban
		api.POST("/exhibitions/:id/ban", authMiddleware("admin"), exhibitionHandler.BanExhibition)
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

// initRoomHandler initializes the Room handler with required dependencies
func initRoomHandler(client *mongo.Client) *roomhandler.Handler {
	dbCollection := client.Database("atommuse").Collection("exhibitionRooms")
	repo := &roomrepo.RoomRepository{Collection: dbCollection}
	service := &roomsvc.RoomServices{Repository: repo}
	return &roomhandler.Handler{RoomService: service}
}
