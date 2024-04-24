package roomhandler

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//	@Summary		Create a new exhibitionRoom
//	@Description	Create a new exhibitionRoom data
//	@Tags			Rooms
//	@Security		BearerAuth
//
//	@ID				CreateExhibitionRoom
//
//	@Accept			json
//	@Produce		json
//	@Param			requestExhibitionRoom	body		model.RequestCreateExhibitionRoom	true	"ExhibitionRoom data to create"
//	@Success		201						{object}	model.ResponseExhibitionRoom		"Success"
//	@Failure		400						{object}	helper.APIError
//	@Failure		401
//	@Failure		500	"Invalid request body"
//	@Router			/api/rooms [post]
func (h *Handler) CreateExhibitionRoom(c *gin.Context) {
	var requestExhibitionRoom model.RequestCreateExhibitionRoom
	var validate = validator.New()

	// Parse request body
	if err := c.BindJSON(&requestExhibitionRoom); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid request body"})
		return
	}

	// Validate the request body
	if err := validate.Struct(requestExhibitionRoom); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": validationErrors})
		return
	}

	// Call use case to create exhibition
	objectID, err := h.RoomService.CreateExhibitionRoom(c.Request.Context(), &requestExhibitionRoom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Failed to create exhibition"})
		return
	}

	// Return the created exhibition ID
	c.JSON(http.StatusCreated, gin.H{"_id": objectID.Hex()})
}
