package fake

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockRepository is a mock implementation of the Repository interface.
type MockRepository struct {
	mock.Mock
}

// GetAllExhibitions is a mock implementation for testing.
func (m *MockRepository) GetAllExhibitions(ctx context.Context) ([]model.ResponseExhibition, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.ResponseExhibition), args.Error(1)
}

// GetExhibitionByID is a mock implementation for testing.
func (m *MockRepository) GetExhibitionByID(ctx context.Context, exhibitionID string) (*model.ResponseExhibition, error) {
	args := m.Called(ctx, exhibitionID)
	return args.Get(0).(*model.ResponseExhibition), args.Error(1)
}

// GetExhibitionsIsPublic is a mock implementation for testing.
func (m *MockRepository) GetExhibitionsIsPublic(ctx context.Context) ([]model.ResponseExhibition, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.ResponseExhibition), args.Error(1)
}

// CreateExhibition is a mock implementation for testing.
func (m *MockRepository) CreateExhibition(ctx context.Context, exhibition *model.RequestCreateExhibition) (*primitive.ObjectID, error) {
	args := m.Called(ctx, exhibition)
	return args.Get(0).(*primitive.ObjectID), args.Error(1)
}

// DeleteExhibition is a mock implementation for testing.
func (m *MockRepository) DeleteExhibition(ctx context.Context, exhibitionID string) error {
	args := m.Called(ctx, exhibitionID)
	return args.Error(0)
}
