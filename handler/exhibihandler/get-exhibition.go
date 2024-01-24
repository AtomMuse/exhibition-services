package exhibihandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllExhibitions godoc
// @Summary Get all exhibitions
// @Description Get a list of all exhibitions
// @Tags Exhibitions
// @ID GetAllExhibitions
// @Produce json
// @Success 200 {array} string "Success"
// @Failure 400 {string} string "Invalid request body"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Permission denied"
// @Failure 500 {string} string "Internal server error"
// @Router /exhibitions [get]
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

// GetExhibition godoc
// @Summary Get exhibition by ID
// @Description Get exhibition details by ID
// @Tags Exhibitions
// @ID GetExhibitionByID
// @Produce json
// @Param id path string true "Exhibition ID"
// @Success 200 {object} object "Success"
// @Failure 400 {string} string "Invalid request body"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Permission denied"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /exhibitions/{id} [get]
func (h *Handler) GetExhibitionByID(c *gin.Context) {
	exhibitionID := c.Param("id")

	exhibition, err := h.UseCase.GetExhibitionByID(c.Request.Context(), exhibitionID)
	if err != nil {
		log.Printf("Error retrieving exhibition %s: %v", exhibitionID, err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibition)
}
