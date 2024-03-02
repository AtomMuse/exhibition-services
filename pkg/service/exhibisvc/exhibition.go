package exhibisvc

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/repositorty/exhibirepo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IExhibitionServices defines the interface for exhibition services.
type IExhibitionServices interface {
	GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error)
	GetExhibitionsIsPublic(context.Context) ([]model.ResponseExhibition, error)
	CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error)
	DeleteExhibition(ctx context.Context, exhibitionID string) error
	UpdateExhibition(ctx context.Context, exhibitionID string, update *model.RequestUpdateExhibition) (*primitive.ObjectID, error)
	DeleteExhibitionSectionID(ctx context.Context, exhibitionID string, sectionID string) error
}

// ExhibitionServices is the implementation of the IExhibitionServices interface.
type ExhibitionServices struct {
	Repository exhibirepo.IExhibitionRepository
}

func (service ExhibitionServices) GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	return service.Repository.GetAllExhibitions(ctx)
}

func (service ExhibitionServices) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error) {
	return service.Repository.GetExhibitionByID(ctx, exhibitionID)
}

func (service ExhibitionServices) GetExhibitionsIsPublic(ctx context.Context) ([]model.ResponseExhibition, error) {
	return service.Repository.GetExhibitionsIsPublic(ctx)
}

func (service ExhibitionServices) CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error) {
	return service.Repository.CreateExhibition(ctx, exhibition)
}

func (service ExhibitionServices) DeleteExhibition(ctx context.Context, exhibitionID string) error {
	return service.Repository.DeleteExhibition(ctx, exhibitionID)
}

func (service ExhibitionServices) UpdateExhibition(ctx context.Context, exhibitionID string, update *model.RequestUpdateExhibition) (*primitive.ObjectID, error) {
	return service.Repository.UpdateExhibition(ctx, exhibitionID, update)
}

func (service ExhibitionServices) DeleteExhibitionSectionID(ctx context.Context, exhibitionID string, sectionID string) error {
	return service.Repository.DeleteExhibitionSectionID(ctx, exhibitionID, sectionID)
}
