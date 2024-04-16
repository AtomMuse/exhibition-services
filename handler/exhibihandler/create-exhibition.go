package exhibihandler

import (
	_ "atommuse/backend/exhibition-service/pkg/helper"
	"atommuse/backend/exhibition-service/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary		Create a new exhibition
// @Description	Create a new exhibition data
// @Tags			Exhibitions
// @Security		BearerAuth
// @ID				CreateExhibition
// @Accept			json
// @Produce		json
// @Param			requestExhibition	body		model.RequestCreateExhibition	true	"Exhibition data to create"
// @Success		201					{object}	model.ResponseGetExhibitionId	"Success"
// @Failure		400					{object}	helper.APIError					"Invalid request body"
// @Router			/api/exhibitions [post]
func (h *Handler) CreateExhibition(c *gin.Context) {

	// Get user information from request context
	userID, _ := c.Get("user_id")
	firstName, _ := c.Get("user_first_name")
	lastName, _ := c.Get("user_last_name")
	profileImage, _ := c.Get("user_image")
	username, _ := c.Get("user_username")

	var requestExhibition model.RequestCreateExhibition
	var validate = validator.New()

	requestExhibition.UserID.UserID = userID.(primitive.ObjectID)
	requestExhibition.UserID.FirstName = firstName.(string)
	requestExhibition.UserID.LastName = lastName.(string)
	requestExhibition.UserID.ProfileImage = profileImage.(string)
	requestExhibition.UserID.Username = username.(string)

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
