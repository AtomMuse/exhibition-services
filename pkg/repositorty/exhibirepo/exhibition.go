package exhibirepo

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// MongoDBRepository is the MongoDB implementation of the Repository interface.
type MongoDBRepository struct {
	Collection *mongo.Collection
}

func (r *MongoDBRepository) GetAllExhibitions(ctx context.Context) ([]model.Exhibition, error) {
	// Define the aggregation pipeline
	pipeline := primitive.A{
		bson.M{"$match": bson.M{}}, // You can add match conditions here if needed
	}

	// Execute the aggregation
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode the results into a slice of documents
	var exhibitions []model.Exhibition
	if err := cursor.All(ctx, &exhibitions); err != nil {
		return nil, err
	}

	return exhibitions, nil
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
