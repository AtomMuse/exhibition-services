package service

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ExhibitionUseCase is the implementation of the UseCase interface.
type ExhibitionUseCase struct {
	Repository *exhibirepo.MongoDBRepository
}

func (uc *ExhibitionUseCase) GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	return uc.Repository.GetAllExhibitions(ctx)
}

func (uc *ExhibitionUseCase) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error) {
	return uc.Repository.GetExhibitionByID(ctx, exhibitionID)
}

func (uc *ExhibitionUseCase) CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error) {
	return uc.Repository.CreateExhibition(ctx, exhibition)
}

func (uc *ExhibitionUseCase) DeleteExhibition(ctx context.Context, exhibitionID string) error {
	return uc.Repository.DeleteExhibition(ctx, exhibitionID)
}
