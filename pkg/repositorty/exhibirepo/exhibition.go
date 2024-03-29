package exhibirepo

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IExhibitionRepository interface {
	GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error)
	GetExhibitionsIsPublic(ctx context.Context) ([]model.ResponseExhibition, error)
	CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error)
	DeleteExhibition(ctx context.Context, exhibitionID string) error
	UpdateExhibition(ctx context.Context, exhibitionID string, update *model.RequestUpdateExhibition) (*primitive.ObjectID, error)
	UpdateVisitedNumber(ctx context.Context, exhibitionID string, visitedNumber int) error
	LikeExhibition(ctx context.Context, exhibitionID string) error
	UnlikeExhibition(ctx context.Context, exhibitionID string) error
}

// ExhibitionRepository is the MongoDB implementation of the Repository interface.
type ExhibitionRepository struct {
	Collection         *mongo.Collection
	SectionsCollection *mongo.Collection
}

// NewExhibitionRepository creates a new instance of ExhibitionRepository.
func NewExhibitionRepository(ctx context.Context, client *mongo.Client, databaseName string) (*ExhibitionRepository, error) {
	// Specify the collection names
	collection := client.Database(databaseName).Collection("exhibitions")
	sectionCollection := client.Database(databaseName).Collection("exhibitionSections")

	return &ExhibitionRepository{
		Collection:         collection,
		SectionsCollection: sectionCollection,
	}, nil
}

func (r *ExhibitionRepository) GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	// Define the aggregation pipeline
	pipeline := primitive.A{
		bson.M{"$match": bson.M{}},
		bson.M{"$sort": bson.M{"startDate": 1}},
	}

	// Execute the aggregation
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode the results into a slice of documents
	var exhibitions []model.ResponseExhibition
	if err := cursor.All(ctx, &exhibitions); err != nil {
		return nil, err
	}

	return exhibitions, nil
}

// GetExhibitionByID retrieves an exhibition by its ID along with its sections.
func (r *ExhibitionRepository) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error) {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return nil, fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	// Define the aggregation pipeline stages
	pipeline := bson.A{
		bson.D{{"$lookup", bson.D{
			{"from", "exhibitionSections"},
			{"localField", "_id"},
			{"foreignField", "exhibitionID"},
			{"as", "exhibitionSections"},
		}}},
		bson.D{{"$project", bson.D{
			{"_id", 1},
			{"exhibitionName", 1},
			{"exhibitionDescription", 1},
			{"thumbnailImg", 1},
			{"startDate", 1},
			{"endDate", 1},
			{"isPublic", 1},
			{"exhibitionCategories", 1},
			{"exhibitionTags", 1},
			{"userId", 1},
			{"layoutUsed", 1},
			{"exhibitionSections", 1},
			{"visitedNumber", 1},
			{"likeCount", 1},
			{"rooms", 1},
			{"status", 1},
		}}},
		bson.D{{"$match", bson.D{
			{"_id", objectID}, // Match using the converted ObjectID
		}}},
	}

	// Execute the aggregation for exhibition collection
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)

	// Check if any result is found
	if !cursor.Next(ctx) {
		return nil, fmt.Errorf("exhibition not found for ID %s", exhibitionID)
	}

	// Decode the main document from exhibition collection
	var exhibition model.ResponseExhibition
	if err := cursor.Decode(&exhibition); err != nil {
		return nil, fmt.Errorf("decoding error: %v", err)
	}

	// Ensure correct ExhibitionID in ExhibitionSections
	for i := range exhibition.ExhibitionSections {
		// Set the ExhibitionID of each section to the exhibition's ID
		exhibition.ExhibitionSections[i].ExhibitionID = exhibition.ID
	}

	return &exhibition, nil
}

func (r *ExhibitionRepository) GetExhibitionsIsPublic(ctx context.Context) ([]model.ResponseExhibition, error) {
	// Define the match stage for the aggregation pipeline
	matchStage := bson.M{"$match": bson.M{"isPublic": true}}

	// Define the sort stage for the aggregation pipeline
	sortStage := bson.M{"$sort": bson.M{"startDate": 1}}

	// Aggregate pipeline
	pipeline := []bson.M{matchStage, sortStage}

	// Execute the aggregation
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)

	// Decode the results into a slice of documents
	var exhibitions []model.ResponseExhibition
	if err := cursor.All(ctx, &exhibitions); err != nil {
		return nil, err
	}

	return exhibitions, nil
}

func (r *ExhibitionRepository) CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error) {
	result, err := r.Collection.InsertOne(ctx, exhibition)
	if err != nil {
		return nil, err
	}

	// Extract the generated ObjectID from the result
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	return &objectID, nil
}

func (r *ExhibitionRepository) DeleteExhibition(ctx context.Context, exhibitionID string) error {
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

	// Specify the collection names
	sectionCollection := client.Database("atommuse").Collection("exhibitionSections")

	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	// Define the match stage for the exhibition document in the aggregation pipeline
	matchStage := bson.M{"$match": bson.M{"_id": objectID}}

	// Aggregate pipeline for finding the exhibition document
	pipeline := []bson.M{matchStage}

	// Execute the aggregation to find the exhibition document
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Check if any result is found
	if !cursor.Next(ctx) {
		return fmt.Errorf("exhibition not found for ID %s", exhibitionID)
	}

	// Decode the main exhibition document
	var exhibition model.ResponseExhibitionForDelete
	if err := cursor.Decode(&exhibition); err != nil {
		return err
	}

	// Retrieve the exhibitionSectionsIDs from the exhibition document
	exhibitionSectionsIDs := exhibition.ExhibitionSectionsID

	// Perform the deletion of the exhibition document
	deleteResult, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("exhibition not deleted for ID %s", exhibitionID)
	}

	// Now, delete associated exhibitionSections
	for _, sectionID := range exhibitionSectionsIDs {
		sectionObjectID, err := primitive.ObjectIDFromHex(sectionID)
		fmt.Println(sectionObjectID)
		if err != nil {
			// Handle error
			fmt.Println("err1", err)
		}

		_, err = sectionCollection.DeleteOne(ctx, bson.M{"_id": sectionObjectID})
		if err != nil {
			// Handle error
			fmt.Println("err2", err)
		}
	}

	return nil
}

func (r *ExhibitionRepository) UpdateExhibition(ctx context.Context, exhibitionID string, update *model.RequestUpdateExhibition) (*primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	updateDoc := bson.M{}

	// Iterate over fields in the update struct and set them in the update document
	updateDoc["$set"] = bson.M{
		"exhibitionName":        update.ExhibitionName,
		"exhibitionDescription": update.ExhibitionDescription,
		"thumbnailImg":          update.ThumbnailImg,
		"startDate":             update.StartDate,
		"endDate":               update.EndDate,
		"isPublic":              update.IsPublic,
		"exhibitionCategories":  update.ExhibitionCategories,
		"exhibitionTags":        update.ExhibitionTags,
		"userId":                update.UserID,
		"layoutUsed":            update.LayoutUsed,
		"exhibitionSectionsID":  update.ExhibitionSectionsID,
		"visitedNumber":         update.VisitedNumber,
		"rooms":                 update.Room,
	}

	// Perform the update operation
	result, err := r.Collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("no exhibition updated")
	}

	return &objectID, nil
}

func (r *ExhibitionRepository) UpdateVisitedNumber(ctx context.Context, exhibitionID string, visitedNumber int) error {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	// Define the filter to find the exhibition by its ID
	filter := bson.M{"_id": objectID}

	// Define the update to increment the visited number
	update := bson.M{"$set": bson.M{"visitedNumber": visitedNumber}}

	// Perform the update operation
	_, err = r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating visited number: %v", err)
	}

	return nil
}

func (r *ExhibitionRepository) LikeExhibition(ctx context.Context, exhibitionID string) error {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return fmt.Errorf("invalid exhibition ID format: %v", err)
	}
	result, err := r.Collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$inc": bson.M{"likeCount": 1}})
	if err != nil {
		return fmt.Errorf("failed to update like count: %v", err)
	}
	if result.ModifiedCount == 0 {
		return errors.New("no documents updated")
	}
	return nil
}

func (r *ExhibitionRepository) UnlikeExhibition(ctx context.Context, exhibitionID string) error {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return fmt.Errorf("invalid exhibition ID format: %v", err)
	}
	result, err := r.Collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$inc": bson.M{"likeCount": -1}})
	if err != nil {
		return fmt.Errorf("failed to update like count: %v", err)
	}
	if result.ModifiedCount == 0 {
		return errors.New("no documents updated")
	}
	return nil
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
