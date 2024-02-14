package exhibihandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteExhibition godoc
// @Summary Delete exhibition by ID
// @Description Delete exhibition by ID
// @Tags Exhibitions
// @ID DeleteExhibition
// @Produce json
// @Param id path string true "Exhibition ID"
// @Success 204 "Delete Exhibition Success"
// @Failure 400 {object} string "Invalid request body"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Permission denied"
// @Failure 404 {object} string "Not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/exhibitions/{id} [delete]
func (h *Handler) DeleteExhibition(c *gin.Context) {
	exhibitionID := c.Param("id")

	err := h.Service.DeleteExhibition(c.Request.Context(), exhibitionID)
	if err != nil {
		log.Printf("Error deleting exhibition %s: %v", exhibitionID, err)

		// Check for specific errors and return appropriate responses
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Exhibition not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Return a successful response with no content (HTTP 204 No Content)
	c.JSON(http.StatusOK, nil)
}
