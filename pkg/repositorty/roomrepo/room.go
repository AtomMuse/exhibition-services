package roomrepo

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRoomRepository interface {
	CreateExhibitionRoom(ctx context.Context, Room *model.RequestCreateExhibitionRoom) (*primitive.ObjectID, error)
	DeleteExhibitionRoomByID(ctx context.Context, RoomID string) error
	GetExhibitionRoomByID(ctx context.Context, RoomID string) (*model.ResponseExhibitionRoom, error)
	GetAllExhibitionRooms(ctx context.Context) ([]model.ResponseExhibitionRoom, error)
	GetRoomsByExhibitionID(ctx context.Context, exhibitionID string) ([]model.Room, error)
	UpdateExhibitionRoom(ctx context.Context, RoomID string, updatedRoom *model.RequestUpdateExhibitionRoom) (*primitive.ObjectID, error)
}

// RoomRepository is the MongoDB implementation of the Repository interface.
type RoomRepository struct {
	Collection *mongo.Collection
}

func (r *RoomRepository) CreateExhibitionRoom(ctx context.Context, Room *model.RequestCreateExhibitionRoom) (*primitive.ObjectID, error) {
	result, err := r.Collection.InsertOne(ctx, Room)
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

func (r *RoomRepository) DeleteExhibitionRoomByID(ctx context.Context, RoomID string) error {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(RoomID)
	if err != nil {
		return fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	// Define the match stage for the aggregation pipeline
	matchStage := bson.M{"$match": bson.M{"_id": objectID}}

	// Aggregate pipeline
	pipeline := []bson.M{matchStage}

	// Execute the aggregation
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Check if any result is found
	if !cursor.Next(ctx) {
		return fmt.Errorf("exhibitionRoom not found for ID %s", RoomID)
	}

	// Decode the main document (assuming ResponseExhibitionRoom is the type of your MongoDB documents)
	Room := model.Room{}
	if err := cursor.Decode(&Room); err != nil {
		return err
	}

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

	// Specify the collection names
	exhibitionCollection := client.Database("atommuse").Collection("exhibitions")

	// Define the filter to match the main exhibition document
	mainExhibitionFilter := bson.M{"_id": Room.ExhibitionID}

	// Define the update to pull the RoomID from the array
	update := bson.M{"$pull": bson.M{"roomsID": RoomID}}

	// Perform the update operation on the main exhibition document
	updateResult, err := exhibitionCollection.UpdateMany(ctx, mainExhibitionFilter, update)
	if err != nil {
		return err
	}

	// Check if any document was modified
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("exhibitionRoom not found for ID %s", RoomID)
	}

	// Perform the deletion
	deleteResult, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("exhibitionRoom not deleted for ID %s", RoomID)
	}

	// Check if exactly one document was deleted
	if deleteResult.DeletedCount == 1 {
		// Assuming UpdateExhibition takes a context, an ExhibitionID, and an update as arguments
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RoomRepository) GetExhibitionRoomByID(ctx context.Context, RoomID string) (*model.ResponseExhibitionRoom, error) {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(RoomID)
	if err != nil {
		return nil, fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	// Define the match stage for the aggregation pipeline
	matchStage := bson.M{"$match": bson.M{"_id": objectID}}

	// Aggregate pipeline
	pipeline := []bson.M{matchStage}

	// Execute the aggregation
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)

	// Check if the cursor is nil
	if cursor == nil {
		return nil, fmt.Errorf("cursor is nil")
	}

	// Check if any result is found
	if !cursor.Next(ctx) {
		return nil, fmt.Errorf("exhibition not found for ID %s", RoomID)
	}

	// Decode the main document
	var exhibitionRoom model.ResponseExhibitionRoom
	if err := cursor.Decode(&exhibitionRoom); err != nil {
		return nil, fmt.Errorf("decoding error: %v", err)
	}

	return &exhibitionRoom, nil
}

func (r *RoomRepository) GetAllExhibitionRooms(ctx context.Context) ([]model.ResponseExhibitionRoom, error) {
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
	var exhibitionRooms []model.ResponseExhibitionRoom
	if err := cursor.All(ctx, &exhibitionRooms); err != nil {
		return nil, err
	}

	return exhibitionRooms, nil
}

// GetRoomsByExhibitionID fetches Rooms for a given exhibition ID from MongoDB.
func (r *RoomRepository) GetRoomsByExhibitionID(ctx context.Context, exhibitionID string) ([]model.Room, error) {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return nil, fmt.Errorf("invalid exhibition ID format: %v", err)
	}
	cursor, err := r.Collection.Find(ctx, bson.M{"exhibitionID": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var Rooms []model.Room
	for cursor.Next(ctx) {
		var Room model.Room
		if err := cursor.Decode(&Room); err != nil {
			return nil, err
		}
		Rooms = append(Rooms, Room)
	}

	return Rooms, nil
}

func (r *RoomRepository) UpdateExhibitionRoom(ctx context.Context, roomID string, updatedRoom *model.RequestUpdateExhibitionRoom) (*primitive.ObjectID, error) {
	// Convert roomID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, err
	}

	// Define filter to identify the room to update
	filter := bson.M{"_id": objectID}

	// Define update operation
	updateDoc := bson.M{}

	// Update all fields from the updatedRoom
	updateDoc["$set"] = bson.M{
		"mapThumbnail": updatedRoom.MapThumbnail,
		"left":         updatedRoom.Left,
		"center":       updatedRoom.Center,
		"right":        updatedRoom.Right,
		"exhibitionID": updatedRoom.ExhibitionID,
	}

	// Perform update operation
	result, err := r.Collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("no exhibition room updated")
	}

	return &objectID, nil
}
