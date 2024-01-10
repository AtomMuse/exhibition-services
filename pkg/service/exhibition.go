package service

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"context"
)

// UseCase represents the business logic.
type UseCase interface {
	GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.Exhibition, error)
}

// ExhibitionUseCase is the implementation of the UseCase interface.
type ExhibitionUseCase struct {
	Repository exhibirepo.Repository
}

func (uc *ExhibitionUseCase) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.Exhibition, error) {
	return uc.Repository.GetExhibitionByID(ctx, exhibitionID)
}
