package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new exhibition
// @Description Create a new exhibition
// @Tags Exhibitions
// @Accept json
// @Produce json
// @Param requestExhibition body string true "Exhibition data to create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Invalid request body"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Permission denied"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/exhibitions [post]
func (h *Handler) CreateExhibition(c *gin.Context) {
	var requestExhibition model.RequestCreateExhibition

	// Parse request body
	if err := c.BindJSON(&requestExhibition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case to create exhibition
	objectID, err := h.Service.CreateExhibition(c.Request.Context(), &requestExhibition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create exhibition"})
		return
	}

	// Return the created exhibition ID
	c.JSON(http.StatusCreated, gin.H{"id": objectID.Hex()})
}
