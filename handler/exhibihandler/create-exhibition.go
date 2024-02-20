package exhibihandler

import (
	_ "atommuse/backend/exhibition-service/pkg/helper"
	"atommuse/backend/exhibition-service/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		Create a new exhibition
//	@Description	Create a new exhibition
//	@Tags			Exhibitions
//	@Accept			json
//	@Produce		json
//	@Param			requestExhibition	body		model.RequestCreateExhibition	true	"Exhibition data to create"
//	@Success		201					{object}	model.ResponseGetExhibitionId	"Success"
//	@Failure		400					{object}	helper.APIError					"Invalid request body"
//	@Router			/api/exhibitions [post]
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
	c.JSON(http.StatusCreated, gin.H{"_id": objectID.Hex()})
}
