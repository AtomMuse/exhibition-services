package sectionsvc

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"atommuse/backend/exhibition-service/pkg/repositorty/sectionrepo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SectionServices defines the interface for exhibition section services.
type ISectionServices interface {
	CreateExhibitionSection(ctx context.Context, section *model.RequestCreateExhibitionSection) (*primitive.ObjectID, error)
	DeleteExhibitionSectionByID(ctx context.Context, sectionID string) error
	GetExhibitionSectionByID(ctx context.Context, sectionID string) (*model.ResponseExhibitionSection, error)
	GetAllExhibitionSections(ctx context.Context) ([]model.ResponseExhibitionSection, error)
}

// SectionServices is the implementation of the IExhibitionSectionServices interface.
type SectionServices struct {
	Repository sectionrepo.ISectionRepository
}

func (service SectionServices) CreateExhibitionSection(ctx context.Context, section *model.RequestCreateExhibitionSection) (*primitive.ObjectID, error) {
	return service.Repository.CreateExhibitionSection(ctx, section)
}

func (service SectionServices) DeleteExhibitionSectionByID(ctx context.Context, sectionID string) error {
	return service.Repository.DeleteExhibitionSectionByID(ctx, sectionID)
}

func (service SectionServices) GetExhibitionSectionByID(ctx context.Context, sectionID string) (*model.ResponseExhibitionSection, error) {
	return service.Repository.GetExhibitionSectionByID(ctx, sectionID)
}

func (service SectionServices) GetAllExhibitionSections(ctx context.Context) ([]model.ResponseExhibitionSection, error) {
	return service.Repository.GetAllExhibitionSections(ctx)
}
