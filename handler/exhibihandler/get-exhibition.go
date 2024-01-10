package exhibihandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetAllExhibitions(w http.ResponseWriter, r *http.Request) {
	exhibitions, err := h.UseCase.GetAllExhibitions(r.Context())
	if err != nil {
		log.Printf("Error retrieving exhibitions : %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Respond with the list of exhibitions
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exhibitions)
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
