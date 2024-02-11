package fake

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Services represents the business logic.
type Service interface {
	GetExhibitionByID(ctx context.Context, exhibitionID primitive.ObjectID) (*model.ResponseExhibition, error)
	GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
}
