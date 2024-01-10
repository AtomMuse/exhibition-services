package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is responsible for handling HTTP requests.
type Handler struct {
	UseCase service.UseCase
}

// GetExhibitionHandler handles HTTP requests for retrieving exhibition by ID.
func (h *Handler) GetExhibitionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exhibitionID := vars["id"]

	exhibition, err := h.UseCase.GetExhibitionByID(r.Context(), exhibitionID)
	if err != nil {
		log.Printf("Error retrieving exhibition %s: %v", exhibitionID, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exhibition)
}
