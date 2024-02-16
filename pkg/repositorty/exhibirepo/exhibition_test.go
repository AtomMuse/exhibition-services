package exhibirepo_test

import (
	"atommuse/backend/exhibition-service/internal/fake"
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/service"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllExhibitions(t *testing.T) {
	// Create a mock repository
	mockRepo := &fake.MockRepository{}

	// Set expectations for the mock repository
	expectedExhibitions := []model.ResponseExhibition{
		// Define your expected data here
	}

	mockRepo.On("GetAllExhibitions", mock.Anything).Return(expectedExhibitions, nil)

	// Create the service with the mock repository
	exhibitionService := service.ExhibitionServices{
		Repository: mockRepo,
	}

	// Call the service method
	exhibitions, err := exhibitionService.GetAllExhibitions(context.Background())

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedExhibitions, exhibitions)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetExhibitionByID(t *testing.T) {
	// Create a mock repository
	mockRepo := &fake.MockRepository{}

	// Set expectations for the mock repository
	expectedExhibition := &model.ResponseExhibition{
		// Define your expected data here
	}

	mockRepo.On("GetExhibitionByID", mock.Anything, mock.AnythingOfType("string")).Return(expectedExhibition, nil)

	// Create the service with the mock repository
	exhibitionService := service.ExhibitionServices{
		Repository: mockRepo,
	}

	// Call the service method
	exhibition, err := exhibitionService.GetExhibitionByID(context.Background(), "validID")

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedExhibition, exhibition)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestCreateExhibition(t *testing.T) {
	// Create a mock repository
	mockRepo := &fake.MockRepository{}

	// Set expectations for the mock repository
	expectedObjectID := &primitive.ObjectID{}

	mockRepo.On("CreateExhibition", mock.Anything, mock.AnythingOfType("*model.RequestCreateExhibition")).Return(expectedObjectID, nil)

	// Create the service with the mock repository
	exhibitionService := service.ExhibitionServices{
		Repository: mockRepo,
	}

	// Call the service method
	objectID, err := exhibitionService.CreateExhibition(context.Background(), &model.RequestCreateExhibition{})

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedObjectID, objectID)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDeleteExhibition(t *testing.T) {
	// Create a mock repository
	mockRepo := &fake.MockRepository{}

	// Set expectations for the mock repository
	mockRepo.On("DeleteExhibition", mock.Anything, mock.AnythingOfType("string")).Return(nil)

	// Create the service with the mock repository
	exhibitionService := service.ExhibitionServices{
		Repository: mockRepo,
	}

	// Call the service method
	err := exhibitionService.DeleteExhibition(context.Background(), "validID")

	// Assert the results
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
