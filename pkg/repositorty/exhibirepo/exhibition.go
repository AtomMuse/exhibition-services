package exhibirepo

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IExhibitionRepository interface {
	GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error)
	GetExhibitionsIsPublic(ctx context.Context) ([]model.ResponseExhibition, error)
	CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error)
	DeleteExhibition(ctx context.Context, exhibitionID string) error
	UpdateExhibition(ctx context.Context, exhibitionID string, update *model.RequestUpdateExhibition) (*primitive.ObjectID, error)
	DeleteExhibitionSectionID(ctx context.Context, exhibitionID string, sectionID string) error
}

// ExhibitionRepository is the MongoDB implementation of the Repository interface.
type ExhibitionRepository struct {
	Collection *mongo.Collection
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

func (r *ExhibitionRepository) UpdateExhibition(ctx context.Context, exhibitionID string, update *model.RequestUpdateExhibition) (*primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	updateDoc := bson.M{"$set": bson.M{}}

	// Update only non-empty fields
	if update.ExhibitionName != "" {
		updateDoc["$set"].(bson.M)["exhibitionName"] = update.ExhibitionName
	}
	if update.ExhibitionDescription != "" {
		updateDoc["$set"].(bson.M)["exhibitionDescription"] = update.ExhibitionDescription
	}
	if update.ThumbnailImg != "" {
		updateDoc["$set"].(bson.M)["thumbnailImg"] = update.ThumbnailImg
	}
	if update.StartDate != "" {
		updateDoc["$set"].(bson.M)["startDate"] = update.StartDate
	}
	if update.EndDate != "" {
		updateDoc["$set"].(bson.M)["endDate"] = update.EndDate
	}
	if update.IsPublic {
		updateDoc["$set"].(bson.M)["isPublic"] = update.IsPublic
	}
	if len(update.ExhibitionCategories) > 0 {
		updateDoc["$set"].(bson.M)["exhibitionCategories"] = update.ExhibitionCategories
	}
	if len(update.ExhibitionTags) > 0 {
		updateDoc["$set"].(bson.M)["exhibitionTags"] = update.ExhibitionTags
	}
	if update.UserID != (model.UserID{}) {
		updateDoc["$set"].(bson.M)["userId"] = update.UserID
	}
	if update.LayoutUsed != "" {
		updateDoc["$set"].(bson.M)["layoutUsed"] = update.LayoutUsed
	}
	if len(update.ExhibitionSectionsID) > 0 {
		updateDoc["$set"].(bson.M)["exhibitionSectionsID"] = update.ExhibitionSectionsID
	}
	if update.VisitedNumber != 0 {
		updateDoc["$set"].(bson.M)["visitedNumber"] = update.VisitedNumber
	}
	if len(update.Room) > 0 {
		updateDoc["$set"].(bson.M)["rooms"] = update.Room
	}

	result, err := r.Collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("no exhibition updated")
	}

	return &objectID, nil
}

func (r *ExhibitionRepository) DeleteExhibitionSectionID(ctx context.Context, exhibitionID string, sectionID string) error {
	// Convert the string IDs to ObjectIds
	exhibitionObjectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	sectionObjectID, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return fmt.Errorf("invalid section ID format: %v", err)
	}

	// Define the filter for the UpdateOne operation
	filter := bson.M{"_id": exhibitionObjectID}

	// Define the update to pull the section ID from the array
	update := bson.M{"$pull": bson.M{"exhibitionSectionsID": sectionObjectID}}

	// Perform the update
	updateResult, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Check if any document is updated
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("exhibition section ID not deleted for exhibition ID %s and section ID %s", exhibitionID, sectionID)
	}

	return nil
}
