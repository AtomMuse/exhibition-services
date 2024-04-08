package sectionhandler

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//	@Summary		Update exhibitionSection by sectionID
//	@Description	Update exhibitionSection data by sectionID
//	@Tags			Sections
//	@Security		BearerAuth
//	@ID				UpdateExhibitionSection
//	@Produce		json
//	@Param			id				path		string									true	"ExhibitionSection ID"
//
//	@Param			updateRequest	body		model.RequestUpdateExhibitionSection	true	"ExhibitionSection data to update"
//
//	@Success		200				{object}	model.ResponseExhibition
//	@Failure		500				{object}	helper.APIError	"Internal server error"
//	@Router			/api/sections/{id} [put]
func (h *Handler) UpdateExhibitionSection(c *gin.Context) {
	var requestUpdateExhibitionSection model.RequestUpdateExhibitionSection
	var validate = validator.New()

	// Parse request body
	if err := c.BindJSON(&requestUpdateExhibitionSection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid request body"})
		return
	}

	// Validate the request body
	if err := validate.Struct(requestUpdateExhibitionSection); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": validationErrors})
		return
	}

	// Get section ID from the URL parameter
	sectionID := c.Param("id")

	// Call use case to update exhibition
	objectID, err := h.SectionService.UpdateExhibitionSection(c.Request.Context(), sectionID, &requestUpdateExhibitionSection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exhibition"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"_id": objectID.Hex(), "message": "Section updated successfully"})
}
