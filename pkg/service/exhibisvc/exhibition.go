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
	UpdateVisitedNumber(ctx context.Context, exhibitionID string, visitedNumber int) error
	LikeExhibition(ctx context.Context, exhibitionID string) error
	UnlikeExhibition(ctx context.Context, exhibitionID string) error
	GetExhibitionsByCategory(ctx context.Context, category string) ([]model.ResponseExhibition, error)
	GetCurrentlyExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetPreviouslyExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetUpcomingExhibitions(ctx context.Context) ([]model.ResponseExhibition, error)
	GetExhibitionsByFilter(ctx context.Context, category, status, sortOrder string) ([]model.ResponseExhibition, error)
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

func (service ExhibitionServices) UpdateVisitedNumber(ctx context.Context, exhibitionID string, visitedNumber int) error {
	return service.Repository.UpdateVisitedNumber(ctx, exhibitionID, visitedNumber)
}
func (service ExhibitionServices) LikeExhibition(ctx context.Context, exhibitionID string) error {
	return service.Repository.LikeExhibition(context.Background(), exhibitionID)
}

func (service ExhibitionServices) UnlikeExhibition(ctx context.Context, exhibitionID string) error {
	return service.Repository.UnlikeExhibition(context.Background(), exhibitionID)
}
func (service ExhibitionServices) GetExhibitionsByCategory(ctx context.Context, category string) ([]model.ResponseExhibition, error) {
	return service.Repository.GetExhibitionsByCategory(ctx, category)
}
func (service ExhibitionServices) GetCurrentlyExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	return service.Repository.GetCurrentlyExhibitions(ctx)
}

func (service ExhibitionServices) GetPreviouslyExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	return service.Repository.GetPreviouslyExhibitions(ctx)
}

func (service ExhibitionServices) GetUpcomingExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	return service.Repository.GetUpcomingExhibitions(ctx)
}
func (service ExhibitionServices) GetExhibitionsByFilter(ctx context.Context, category, status, sortOrder string) ([]model.ResponseExhibition, error) {
	return service.Repository.GetExhibitionsByFilter(ctx, category, status, sortOrder)
}
