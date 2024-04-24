package roomhandler

import (
	"atommuse/backend/exhibition-service/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//	@Summary		Update exhibitionRoom by RoomID
//	@Description	Update exhibitionRoom data by RoomID
//	@Tags			Rooms
//	@Security		BearerAuth
//	@ID				UpdateExhibitionRoom
//	@Produce		json
//	@Param			id				path		string								true	"ExhibitionRoom ID"
//
//	@Param			updateRequest	body		model.RequestUpdateExhibitionRoom	true	"ExhibitionRoom data to update"
//
//	@Success		200				{object}	model.ResponseExhibition
//	@Failure		401
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/rooms/{id} [put]
func (h *Handler) UpdateExhibitionRoom(c *gin.Context) {
	var requestUpdateExhibitionRoom model.RequestUpdateExhibitionRoom
	var validate = validator.New()

	// Parse request body
	if err := c.BindJSON(&requestUpdateExhibitionRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid request body"})
		return
	}

	// Validate the request body
	if err := validate.Struct(requestUpdateExhibitionRoom); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": validationErrors})
		return
	}

	// Get Room ID from the URL parameter
	RoomID := c.Param("id")

	// Call use case to update exhibition
	objectID, err := h.RoomService.UpdateExhibitionRoom(c.Request.Context(), RoomID, &requestUpdateExhibitionRoom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exhibition"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"_id": objectID.Hex(), "message": "Room updated successfully"})
}
