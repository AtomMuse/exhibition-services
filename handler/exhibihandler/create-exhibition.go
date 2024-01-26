package exhibihandler

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handler struct definition

// CreateExhibition godoc
// @Summary Create a new exhibition
// @Description Create a new exhibition
// @Tags Exhibitions
// @ID CreateExhibition
// @Produce json
// @Consumes json
// @Param requestExhibition body model.ResponseExhibition true "Exhibition details"
// @Success 201 {object} object "Created"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/exhibitions [post]
func (h *Handler) CreateExhibition(c *gin.Context) {
	var requestExhibition model.RequestCreateExhibition

	// Parse request body
	if err := c.ShouldBindJSON(&requestExhibition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "validationErrors": err.Error()})
		return
	}

	// Check for required fields in JSON
	requiredFields := []string{"ExhibitionName", "StartDate", "EndDate", "UserID", "LayoutUsed"}
	missingFields := getMissingFields(requestExhibition, requiredFields)
	if len(missingFields) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields", "missingFields": missingFields})
		return
	}

	// Validate required and non-empty fields
	validate := validator.New()
	validate.RegisterValidation("customDate", customDateValidator) // Register custom date validator
	if err := validate.Struct(requestExhibition); err != nil {
		// Map validation errors for better response
		validationErrors := make(map[string]string)
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				validationErrors[e.Field()] = e.Tag()
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "validationErrors": validationErrors})
		return
	}

	// Call use case to create exhibition
	objectID, err := h.UseCase.CreateExhibition(c.Request.Context(), &requestExhibition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create exhibition"})
		return
	}

	// Return the created exhibition ID
	c.JSON(http.StatusCreated, gin.H{"id": objectID.Hex()})
}

// getMissingFields checks for missing required fields in the request
func getMissingFields(exhibition model.RequestCreateExhibition, requiredFields []string) []string {
	var missingFields []string

	// Check for missing required fields
	for _, field := range requiredFields {
		switch field {
		case "ExhibitionName":
			if exhibition.ExhibitionName == "" {
				missingFields = append(missingFields, field)
			}
		case "StartDate":
			if exhibition.StartDate == "" {
				missingFields = append(missingFields, field)
			}
		case "EndDate":
			if exhibition.EndDate == "" {
				missingFields = append(missingFields, field)
			}
		case "UserID":
			if exhibition.UserID.UserID == 0 {
				missingFields = append(missingFields, field)
			}
		case "LayoutUsed":
			if exhibition.LayoutUsed == "" {
				missingFields = append(missingFields, field)
			}
		}
	}

	return missingFields
}

// Custom date validator
func customDateValidator(fl validator.FieldLevel) bool {
	dateString := fl.Field().String()
	_, err := time.Parse("2006-01-02", dateString)
	return err == nil
}
