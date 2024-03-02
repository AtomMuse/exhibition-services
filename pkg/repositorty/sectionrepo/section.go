package sectionrepo

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type ISectionRepository interface {
	CreateExhibitionSection(ctx context.Context, section *model.RequestCreateExhibitionSection) (*primitive.ObjectID, error)
	DeleteExhibitionSectionByID(ctx context.Context, sectionID string) error
	GetExhibitionSectionByID(ctx context.Context, sectionID string) (*model.ResponseExhibitionSection, error)
}

// SectionRepository is the MongoDB implementation of the Repository interface.
type SectionRepository struct {
	Collection *mongo.Collection
}

func (r *SectionRepository) CreateExhibitionSection(ctx context.Context, section *model.RequestCreateExhibitionSection) (*primitive.ObjectID, error) {
	result, err := r.Collection.InsertOne(ctx, section)
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

func (r *SectionRepository) DeleteExhibitionSectionByID(ctx context.Context, sectionID string) error {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	// Create an instance of the concrete type that implements IExhibitionRepository
	exhibitionRepo := &exhibirepo.ExhibitionRepository{Collection: r.Collection}

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
		return fmt.Errorf("exhibitionSection not found for ID %s", sectionID)
	}

	// Decode the main document (assuming ResponseExhibitionSection is the type of your MongoDB documents)
	section := model.ExhibitionSection{}
	if err := cursor.Decode(&section); err != nil {
		return err
	}
	fmt.Println(section)

	// Perform the deletion
	deleteResult, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("exhibitionSection not deleted for ID %s", sectionID)
	}

	// Check if exactly one document was deleted
	if deleteResult.DeletedCount == 1 {
		// Assuming UpdateExhibition takes a context, an ExhibitionID, and an update as arguments
		err := exhibitionRepo.DeleteExhibitionSectionID(ctx, section.ExhibitionID.Hex(), sectionID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SectionRepository) GetExhibitionSectionByID(ctx context.Context, sectionID string) (*model.ResponseExhibitionSection, error) {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(sectionID)
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
		return nil, fmt.Errorf("exhibition not found for ID %s", sectionID)
	}

	// Decode the main document
	var exhibitionSection model.ResponseExhibitionSection
	if err := cursor.Decode(&exhibitionSection); err != nil {
		return nil, fmt.Errorf("decoding error: %v", err)
	}

	return &exhibitionSection, nil
}
