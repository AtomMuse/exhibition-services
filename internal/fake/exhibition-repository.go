package fake

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
)

// Repository interface defines the methods for data access.
type Repository interface {
	GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.Exhibition, error)
	GetAllExhibitions(ctx context.Context) ([]model.Exhibition, error)
}
