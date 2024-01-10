package service

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"context"
)

// ExhibitionUseCase is the implementation of the UseCase interface.
type ExhibitionUseCase struct {
	Repository *exhibirepo.MongoDBRepository
}

func (uc *ExhibitionUseCase) GetAllExhibitions(ctx context.Context) ([]model.Exhibition, error) {
	return uc.Repository.GetAllExhibitions(ctx)
}

func (uc *ExhibitionUseCase) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.Exhibition, error) {
	return uc.Repository.GetExhibitionByID(ctx, exhibitionID)
}
