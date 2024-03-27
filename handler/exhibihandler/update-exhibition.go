package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UpdateExhibition godoc
//
//	@Summary		Update exhibition by ID
//	@Description	Update exhibition data by exhibitionID
//	@Tags			Exhibitions
//	@Security		BearerAuth
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
	var validate = validator.New()
	var updateRequest model.RequestUpdateExhibition

	// Parse request body
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate the request body
	if err := validate.Struct(updateRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": validationErrors})
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
