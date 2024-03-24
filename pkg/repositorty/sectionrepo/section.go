package sectionrepo

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ISectionRepository interface {
	CreateExhibitionSection(ctx context.Context, section *model.RequestCreateExhibitionSection) (*primitive.ObjectID, error)
	DeleteExhibitionSectionByID(ctx context.Context, sectionID string) error
	GetExhibitionSectionByID(ctx context.Context, sectionID string) (*model.ResponseExhibitionSection, error)
	GetAllExhibitionSections(ctx context.Context) ([]model.ResponseExhibitionSection, error)
	GetSectionsByExhibitionID(ctx context.Context, exhibitionID string) ([]model.ExhibitionSection, error)
	UpdateExhibitionSection(ctx context.Context, sectionID string, updatedSection *model.RequestUpdateExhibitionSection) (*primitive.ObjectID, error)
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

func (r *SectionRepository) GetAllExhibitionSections(ctx context.Context) ([]model.ResponseExhibitionSection, error) {
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
	var exhibitionSections []model.ResponseExhibitionSection
	if err := cursor.All(ctx, &exhibitionSections); err != nil {
		return nil, err
	}

	return exhibitionSections, nil
}

// GetSectionsByExhibitionID fetches sections for a given exhibition ID from MongoDB.
func (r *SectionRepository) GetSectionsByExhibitionID(ctx context.Context, exhibitionID string) ([]model.ExhibitionSection, error) {
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

	var sections []model.ExhibitionSection
	for cursor.Next(ctx) {
		var section model.ExhibitionSection
		if err := cursor.Decode(&section); err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}

	return sections, nil
}

func (r *SectionRepository) UpdateExhibitionSection(ctx context.Context, sectionID string, updatedSection *model.RequestUpdateExhibitionSection) (*primitive.ObjectID, error) {
	// Convert sectionID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return nil, err
	}

	// Define filter to identify the section to update
	filter := bson.M{"_id": objectID}

	// Define update operation
	updateDoc := bson.M{}

	// Update all fields from the updatedSection
	updateDoc["$set"] = bson.M{
		"sectionType":  updatedSection.SectionType,
		"contentType":  updatedSection.ContentType,
		"background":   updatedSection.Background,
		"title":        updatedSection.Title,
		"text":         updatedSection.Text,
		"leftCol":      updatedSection.LeftCol,
		"rightCol":     updatedSection.RightCol,
		"images":       updatedSection.Images,
		"exhibitionID": updatedSection.ExhibitionID,
	}

	// Perform update operation
	result, err := r.Collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("no exhibition section updated")
	}

	return &objectID, nil
}
