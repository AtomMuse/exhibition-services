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
	GetExhibitionByUserID(ctx context.Context, userID string) ([]*model.ResponseExhibition, error)
	CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error)
	DeleteExhibition(ctx context.Context, exhibitionID string) error
	UpdateExhibition(ctx context.Context, exhibitionID string, update *model.RequestUpdateExhibition) (*primitive.ObjectID, error)
	UpdateVisitedNumber(ctx context.Context, exhibitionID string, visitedNumber int) error
	LikeExhibition(ctx context.Context, exhibitionID string) error
	UnlikeExhibition(ctx context.Context, exhibitionID string) error
	GetExhibitionsByCategory(ctx context.Context, category string) ([]model.ResponseExhibition, error)
	GetCurrentlyExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetPreviouslyExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetUpcomingExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetExhibitionsByFilter(ctx context.Context, category, status, sortOrder string) ([]model.ResponseExhibition, error)
	GetExhibitionSectionInfo(ctx context.Context, exhibitionID string) ([]model.ExhibitionSectionInfo, error)
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
	// Convert exhibitionID to ObjectID
	objID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return nil, err
	}

	// Find exhibition by ID
	var exhibition model.ResponseExhibition
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&exhibition)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("exhibition not found")
		}
		return nil, err
	}

	// if exhibition.LayoutUsed == "blogLayout" {

	// Find sections related to the exhibition
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

	// Loop through exhibition section IDs
	var sections []model.ExhibitionSection
	fmt.Println("sec1", sections)
	for _, sectionID := range exhibition.ExhibitionSectionsID {
		// Convert sectionID to ObjectID
		sectionObjID, err := primitive.ObjectIDFromHex(sectionID)
		if err != nil {
			return nil, err
		}

		// Find section by ID
		var section model.ExhibitionSection
		fmt.Println("sec2", section)
		err = sectionCollection.FindOne(ctx, bson.M{"_id": sectionObjID}).Decode(&section)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("section not found")
			}
			return nil, err
		}

		sections = append(sections, section)
	}

	// Assign sections to the exhibition
	exhibition.ExhibitionSections = sections
	// }
	return &exhibition, nil
}

func (r *ExhibitionRepository) GetSectionsByExhibitionID(ctx context.Context, exhibitionID primitive.ObjectID) ([]model.ExhibitionSection, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return nil, errors.New("MONGO_URI environment variable not set")
	}

	client, err := connectToMongoDB(mongoURI)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// Specify the collection names
	sectionCollection := client.Database("atommuse").Collection("exhibitionSections")

	// Find sections by exhibition ID
	cursor, err := sectionCollection.Find(ctx, bson.M{"exhibitionID": exhibitionID})
	if err != nil {
		return nil, fmt.Errorf("error finding sections for exhibition: %w", err)
	}
	defer cursor.Close(ctx)

	var sections []model.ExhibitionSection
	for cursor.Next(ctx) {
		var section model.ExhibitionSection
		if err := cursor.Decode(&section); err != nil {
			return nil, fmt.Errorf("error decoding section: %w", err)
		}
		sections = append(sections, section)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	if len(sections) == 0 {
		return nil, errors.New("no sections found for the exhibition")
	}

	return sections, nil
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

	mongoURIcomment := os.Getenv("MONGO_URI_COMMENT")
	if mongoURIcomment == "" {
		log.Fatal("MONGO_URI environment variable not set.")
	}
	log.Println("MongoURI:", mongoURIcomment)

	clientComment, err := connectToMongoDB(mongoURIcomment)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := clientComment.Disconnect(context.Background()); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// Specify the collection names
	commentCollection := clientComment.Database("atommuse-comment").Collection("comments")

	// Now, call DeleteCommentsByExhibitionID from commentrepo
	exhibitionObjectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	// Check if there are any comments to delete
	numComments, err := commentCollection.CountDocuments(ctx, bson.M{"exhibitionID": exhibitionObjectID})
	if err != nil {
		return fmt.Errorf("error counting comments for exhibition: %v", err)
	}

	fmt.Println(numComments)

	if numComments > 0 {
		// If there are comments, delete them
		if _, err := commentCollection.DeleteMany(ctx, bson.M{"exhibitionID": exhibitionObjectID}); err != nil {
			return fmt.Errorf("error deleting comments for exhibition: %v", err)
		}
		fmt.Println("Deleted comments for exhibition ID:", exhibitionID)
	} else {
		// If there are no comments, continue execution without error
		log.Printf("No comments found for exhibition ID: %s", exhibitionID)
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

func (r *ExhibitionRepository) GetExhibitionByUserID(ctx context.Context, userID string) ([]*model.ResponseExhibition, error) {

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	// Define the filter for the query
	filter := bson.M{"userId.userId": objectID}

	// Execute the find query
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find error: %v", err)
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and decode each document
	var exhibitions []*model.ResponseExhibition
	for cursor.Next(ctx) {
		var exhibition model.ResponseExhibition
		if err := cursor.Decode(&exhibition); err != nil {
			return nil, fmt.Errorf("decoding error: %v", err)
		}
		exhibitions = append(exhibitions, &exhibition)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return exhibitions, nil
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

func (r *ExhibitionRepository) GetExhibitionsByCategory(ctx context.Context, category string) ([]model.ResponseExhibition, error) {
	// Define the match stage for the aggregation pipeline
	matchStage := bson.M{"$match": bson.M{"exhibitionCategories": category, "isPublic": true}}

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

func (r *ExhibitionRepository) GetCurrentlyExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	matchStage := bson.M{"$match": bson.M{
		"isPublic": true,
		"$expr": bson.M{
			"$and": []bson.M{
				bson.M{"$lte": []interface{}{"$startDate", time.Now()}},
				bson.M{"$gte": []interface{}{"$endDate", time.Now()}},
			},
		},
	}}
	sortStage := bson.M{"$sort": bson.M{"startDate": 1}}
	pipeline := []bson.M{matchStage, sortStage}
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)
	var exhibitions []model.ResponseExhibition
	if err := cursor.All(ctx, &exhibitions); err != nil {
		return nil, err
	}
	return exhibitions, nil
}

func (r *ExhibitionRepository) GetPreviouslyExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	matchStage := bson.M{"$match": bson.M{
		"isPublic": true,
		"endDate":  bson.M{"$lt": time.Now().Format("2006-01-02T15:04:05.000Z")},
	}}
	sortStage := bson.M{"$sort": bson.M{"startDate": 1}}
	pipeline := []bson.M{matchStage, sortStage}
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)
	var exhibitions []model.ResponseExhibition
	if err := cursor.All(ctx, &exhibitions); err != nil {
		return nil, err
	}
	return exhibitions, nil
}

func (r *ExhibitionRepository) GetUpcomingExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	matchStage := bson.M{"$match": bson.M{
		"isPublic":  true,
		"startDate": bson.M{"$gt": time.Now().Format("2006-01-02T15:04:05.000Z")},
	}}
	sortStage := bson.M{"$sort": bson.M{"startDate": 1}}
	pipeline := []bson.M{matchStage, sortStage}
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)
	var exhibitions []model.ResponseExhibition
	if err := cursor.All(ctx, &exhibitions); err != nil {
		return nil, err
	}
	return exhibitions, nil
}

func (r *ExhibitionRepository) GetExhibitionsByFilter(ctx context.Context, category, status, sortOrder string) ([]model.ResponseExhibition, error) {
	// Define the match stage for the aggregation pipeline
	matchStage := bson.M{"$match": bson.M{
		"exhibitionCategories": category,
		"isPublic":             true,
	}}

	// Add status filtering
	switch status {
	case "current":
		matchStage["$match"].(bson.M)["$expr"] = bson.M{
			"$and": []bson.M{
				bson.M{"$lte": []interface{}{"$startDate", time.Now()}},
				bson.M{"$gte": []interface{}{"$endDate", time.Now()}},
			},
		}
	case "previous":
		matchStage["$match"].(bson.M)["endDate"] = bson.M{"$lt": time.Now().Format("2006-01-02T15:04:05.000Z")}
	case "upcoming":
		matchStage["$match"].(bson.M)["startDate"] = bson.M{"$gt": time.Now().Format("2006-01-02T15:04:05.000Z")}
	}

	// Define the sort stage for the aggregation pipeline
	sortStage := bson.M{"$sort": bson.M{"startDate": 1}}
	if sortOrder == "desc" {
		sortStage["$sort"].(bson.M)["startDate"] = -1
	}

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

// GetExhibitionSectionInfo retrieves the section IDs of an exhibition along with its index
func (r *ExhibitionRepository) GetExhibitionSectionInfo(ctx context.Context, exhibitionID string) ([]model.ExhibitionSectionInfo, error) {
	// Convert the string ID to ObjectId
	objectID, err := primitive.ObjectIDFromHex(exhibitionID)
	if err != nil {
		return nil, fmt.Errorf("invalid exhibition ID format: %v", err)
	}

	// Define the match stage for the aggregation pipeline
	matchStage := bson.D{{"$match", bson.D{
		{"_id", objectID},
	}}}

	// Define the project stage to extract only the exhibitionSectionsID field
	projectStage := bson.D{{"$project", bson.D{
		{"_id", 0}, // Exclude _id field
		{"exhibitionSectionsID", 1},
	}}}

	// Aggregate pipeline
	pipeline := []bson.D{matchStage, projectStage}

	// Execute the aggregation
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %v", err)
	}
	defer cursor.Close(ctx)

	// Decode the results into a slice of strings
	var result []struct {
		ExhibitionSectionsID []string `bson:"exhibitionSectionsID"`
	}
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("decoding error: %v", err)
	}

	// Check if any result is found
	if len(result) == 0 {
		return nil, fmt.Errorf("exhibition not found for ID %s", exhibitionID)
	}

	// Construct ExhibitionSectionInfo array
	var infos []model.ExhibitionSectionInfo
	for i, sectionID := range result[0].ExhibitionSectionsID {
		info := model.ExhibitionSectionInfo{
			Index:                i,
			ExhibitionSectionsID: sectionID,
		}
		infos = append(infos, info)
	}

	return infos, nil
}
