package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateExhibition godoc
//
//	@Summary		Update exhibition by ID
//	@Description	Update exhibition details by ID
//	@Tags			Exhibitions
//	@ID				UpdateExhibition
//	@Produce		json
//	@Param			id				path		string							true	"Exhibition ID"
//
//	@Param			updateRequest	body		model.RequestUpdateExhibition	true	"Exhibition data to update"
//
//	@Success		200				{object}	model.ResponseExhibition
//	@Failure		500				{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/{id} [put]
func (h *Handler) UpdateExhibition(c *gin.Context) {
	exhibitionID := c.Param("id") // assuming exhibition ID is part of the URL

	var updateRequest model.RequestUpdateExhibition

	// Parse request body
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case to update exhibition
	objectID, err := h.ExhibitionService.UpdateExhibition(c.Request.Context(), exhibitionID, &updateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exhibition"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"_id": objectID.Hex(), "message": "Exhibition updated successfully"})
}
