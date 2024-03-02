package exhibihandler

import (
	_ "atommuse/backend/exhibition-service/pkg/helper"
	"atommuse/backend/exhibition-service/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	var validate = validator.New()

	// Parse request body
	if err := c.BindJSON(&requestExhibition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid request body"})
		return
	}

	// Validate the request body
	if err := validate.Struct(requestExhibition); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": validationErrors})
		return
	}

	// Call use case to create exhibition
	objectID, err := h.ExhibitionService.CreateExhibition(c.Request.Context(), &requestExhibition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Failed to create exhibition"})
		return
	}

	// Return the created exhibition ID
	c.JSON(http.StatusCreated, gin.H{"_id": objectID.Hex()})
}

//	@Summary		Create a new exhibitionSection
//	@Description	Create a new exhibitionSection
//	@Tags			Exhibitions
//	@Accept			json
//	@Produce		json
//	@Param			requestExhibitionSection	body		model.RequestCreateExhibitionSection	true	"ExhibitionSection data to create"
//	@Success		201							{object}	model.ResponseGetExhibitionSectionId	"Success"
//	@Failure		400							{object}	helper.APIError							"Invalid request body"
//	@Router			/api/sections [post]
func (h *Handler) CreateExhibitionSection(c *gin.Context) {
	var requestExhibitionSection model.RequestCreateExhibitionSection
	var validate = validator.New()

	// Parse request body
	if err := c.BindJSON(&requestExhibitionSection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid request body"})
		return
	}

	// Validate the request body
	if err := validate.Struct(requestExhibitionSection); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": validationErrors})
		return
	}

	// Call use case to create exhibition
	objectID, err := h.SectionService.CreateExhibitionSection(c.Request.Context(), &requestExhibitionSection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Failed to create exhibition"})
		return
	}

	// Return the created exhibition ID
	c.JSON(http.StatusCreated, gin.H{"_id": objectID.Hex()})
}
