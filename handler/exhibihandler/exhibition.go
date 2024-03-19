package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/service/exhibisvc"
)

// Handler is responsible for handling HTTP requests.
type Handler struct {
	ExhibitionService exhibisvc.IExhibitionServices
}
