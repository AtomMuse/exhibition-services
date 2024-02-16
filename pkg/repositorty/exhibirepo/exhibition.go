package exhibirepo

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IExhibitionRepository interface {
	GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error)
	CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error)
	DeleteExhibition(ctx context.Context, exhibitionID string) error
}

// ExhibitionRepository is the MongoDB implementation of the Repository interface.
type ExhibitionRepository struct {
	Collection *mongo.Collection
}

func (r *ExhibitionRepository) GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
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
	var exhibitions []model.ResponseExhibition
	if err := cursor.All(ctx, &exhibitions); err != nil {
		return nil, err
	}

	return exhibitions, nil
}

func (r *ExhibitionRepository) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error) {
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
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)

	// Check if any result is found
	if !cursor.Next(ctx) {
		return nil, fmt.Errorf("exhibition not found for ID %s", exhibitionID)
	}

	// Decode the main document
	var exhibition model.ResponseExhibition
	if err := cursor.Decode(&exhibition); err != nil {
		return nil, fmt.Errorf("decoding error: %v", err)
	}

	return &exhibition, nil
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
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
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
		return fmt.Errorf("exhibition not found for ID %s", exhibitionID)
	}

	// Decode the main document (assuming ResponseExhibition is the type of your MongoDB documents)
	var exhibition model.ResponseExhibition
	if err := cursor.Decode(&exhibition); err != nil {
		return err
	}

	// Perform the deletion
	deleteResult, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("exhibition not deleted for ID %s", exhibitionID)
	}

	return nil
}
