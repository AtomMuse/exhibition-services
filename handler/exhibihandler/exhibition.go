package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/service"
)

// Handler is responsible for handling HTTP requests.
type Handler struct {
	Service *service.ExhibitionServices
}
