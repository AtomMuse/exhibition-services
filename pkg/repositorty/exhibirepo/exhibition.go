package exhibirepo

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Repository interface defines the methods for data access.
type Repository interface {
	GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.Exhibition, error)
}

// MongoDBRepository is the MongoDB implementation of the Repository interface.
type MongoDBRepository struct {
	Collection *mongo.Collection
}

func (r *MongoDBRepository) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.Exhibition, error) {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
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
		return nil, err
	}
	defer cursor.Close(ctx)

	// Check if any result is found
	if !cursor.Next(ctx) {
		return nil, fmt.Errorf("exhibition not found for ID %s", exhibitionID)
	}

	// Decode the main document
	var exhibition model.Exhibition
	if err := cursor.Decode(&exhibition); err != nil {
		return nil, err
	}

	return &exhibition, nil
}
