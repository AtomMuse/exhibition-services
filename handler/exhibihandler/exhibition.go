package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/service/exhibisvc"
	"atommuse/backend/exhibition-service/pkg/service/sectionsvc"
)

// Handler is responsible for handling HTTP requests.
type Handler struct {
	ExhibitionService exhibisvc.IExhibitionServices
	SectionService    sectionsvc.ISectionServices
}
