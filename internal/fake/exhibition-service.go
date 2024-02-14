package fake

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockService struct {
	Repository *MockRepository
}

func (s *MockService) GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	args := s.Repository.Called(ctx)
	return args.Get(0).([]model.ResponseExhibition), args.Error(1)
}

func (s *MockService) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error) {
	args := s.Repository.Called(ctx, exhibitionID)
	return args.Get(0).(*model.ResponseExhibition), args.Error(1)
}

func (s *MockService) CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error) {
	args := s.Repository.Called(ctx, exhibition)
	return args.Get(0).(*primitive.ObjectID), args.Error(1)
}

func (s *MockService) DeleteExhibition(ctx context.Context, exhibitionID string) error {
	args := s.Repository.Called(ctx, exhibitionID)
	return args.Error(0)
}
