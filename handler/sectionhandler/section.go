package sectionhandler

import "atommuse/backend/exhibition-service/pkg/service/sectionsvc"

// Handler is responsible for handling HTTP requests.
type Handler struct {
	SectionService sectionsvc.ISectionServices
}
