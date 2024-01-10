package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/service"
	"net/http"
)

// Handler is responsible for handling HTTP requests.
type Handler struct {
	UseCase *service.ExhibitionUseCase
}

type IExhibitionHandle interface {
	GetAllExhibitions(w http.ResponseWriter, r *http.Request)
	GetExhibitionByID(w http.ResponseWriter, r *http.Request)
}
