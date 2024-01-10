package fake

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
)

// UseCase represents the business logic.
type UseCase interface {
	GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.Exhibition, error)
	GetAllExhibitions(ctx context.Context) ([]model.Exhibition, error)
}
