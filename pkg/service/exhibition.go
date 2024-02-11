package service

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ExhibitionServices is the implementation of the Services interface.
type ExhibitionServices struct {
	Repository *exhibirepo.MongoDBRepository
}

func (service *ExhibitionServices) GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	return service.Repository.GetAllExhibitions(ctx)
}

func (service *ExhibitionServices) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error) {
	return service.Repository.GetExhibitionByID(ctx, exhibitionID)
}

func (service *ExhibitionServices) CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error) {
	return service.Repository.CreateExhibition(ctx, exhibition)
}

func (service *ExhibitionServices) DeleteExhibition(ctx context.Context, exhibitionID string) error {
	return service.Repository.DeleteExhibition(ctx, exhibitionID)
}
