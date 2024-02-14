package service_test

import (
	"atommuse/backend/exhibition-service/internal/fake"
	"atommuse/backend/exhibition-service/pkg/model"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllExhibitions(t *testing.T) {
	// Create a mock repository
	mockRepo := &fake.MockRepository{}
	service := &fake.MockService{Repository: mockRepo}

	// Mock the GetAllExhibitions method in the mock repository
	mockExhibitions := []model.ResponseExhibition{
		{
			ID:                    primitive.NewObjectID(),
			ExhibitionName:        "",
			ExhibitionDescription: "",
			ThumbnailImg:          "",
			StartDate:             "",
			EndDate:               "",
			IsPublic:              false,
			ExhibitionCategories:  []string{},
			ExhibitionTags:        []string{},
			UserID:                model.UserID{},
			LayoutUsed:            "",
			ExhibitionSections:    []model.ExhibitionSection{},
		},
		{
			ID:                    primitive.NewObjectID(),
			ExhibitionName:        "",
			ExhibitionDescription: "",
			ThumbnailImg:          "",
			StartDate:             "",
			EndDate:               "",
			IsPublic:              false,
			ExhibitionCategories:  []string{},
			ExhibitionTags:        []string{},
			UserID:                model.UserID{},
			LayoutUsed:            "",
			ExhibitionSections:    []model.ExhibitionSection{},
		},
	}
	mockRepo.On("GetAllExhibitions", mock.Anything).Return(mockExhibitions, nil)

	// Perform the actual method call
	result, err := service.GetAllExhibitions(context.Background())

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, mockExhibitions, result)

	// Verify that the mock method was called
	mockRepo.AssertExpectations(t)
}

func TestGetExhibitionByID(t *testing.T) {
	// Create a mock repository
	mockRepo := &fake.MockRepository{}
	service := &fake.MockService{Repository: mockRepo}

	// Mock the GetExhibitionByID method in the mock repository
	exhibitionID := "some-exhibition-id"
	mockExhibition := &model.ResponseExhibition{
		ID:                    primitive.NewObjectID(),
		ExhibitionName:        "",
		ExhibitionDescription: "",
		ThumbnailImg:          "",
		StartDate:             "",
		EndDate:               "",
		IsPublic:              false,
		ExhibitionCategories:  []string{},
		ExhibitionTags:        []string{},
		UserID:                model.UserID{},
		LayoutUsed:            "",
		ExhibitionSections:    []model.ExhibitionSection{},
	}
	mockRepo.On("GetExhibitionByID", mock.Anything, exhibitionID).Return(mockExhibition, nil)

	// Perform the actual method call
	result, err := service.GetExhibitionByID(context.Background(), exhibitionID)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, mockExhibition, result)

	// Verify that the mock method was called
	mockRepo.AssertExpectations(t)
}

func TestCreateExhibition(t *testing.T) {
	// Create a mock repository
	mockRepo := &fake.MockRepository{}
	service := &fake.MockService{Repository: mockRepo}

	// Mock the CreateExhibition method in the mock repository
	mockExhibitionID := primitive.NewObjectID()
	mockRequestExhibition := &model.RequestCreateExhibition{
		ExhibitionName:        "",
		ExhibitionDescription: "",
		ThumbnailImg:          "",
		StartDate:             "",
		EndDate:               "",
		IsPublic:              false,
		ExhibitionCategories:  []string{},
		ExhibitionTags:        []string{},
		UserID:                model.UserID{},
		LayoutUsed:            "",
		ExhibitionSections:    []model.ExhibitionSection{},
	}
	mockRepo.On("CreateExhibition", mock.Anything, mockRequestExhibition).Return(&mockExhibitionID, nil)

	// Perform the actual method call
	result, err := service.CreateExhibition(context.Background(), mockRequestExhibition)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, &mockExhibitionID, result)

	// Verify that the mock method was called
	mockRepo.AssertExpectations(t)
}

func TestDeleteExhibition(t *testing.T) {
	// Create a mock repository
	mockRepo := &fake.MockRepository{}
	service := &fake.MockService{Repository: mockRepo}

	// Mock the DeleteExhibition method in the mock repository
	exhibitionID := "some-exhibition-id"
	mockRepo.On("DeleteExhibition", mock.Anything, exhibitionID).Return(nil)

	// Perform the actual method call
	err := service.DeleteExhibition(context.Background(), exhibitionID)

	// Assertions
	assert.Nil(t, err)

	// Verify that the mock method was called
	mockRepo.AssertExpectations(t)
}
