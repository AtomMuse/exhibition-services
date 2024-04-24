package roomhandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//	@Summary		Delete Room by ID
//	@Description	Delete Room data by RoomID
//	@Tags			Rooms
//	@Security		BearerAuth
//	@ID				DeleteExhibitionRoomByID
//	@Produce		json
//	@Param			id	path		string							true	"Room ID"
//	@Success		200	{object}	model.ResponseGetExhibitionId	"Delete Room Success"
//	@Failure		401
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/rooms/{id} [delete]
func (h *Handler) DeleteExhibitionRoomByID(c *gin.Context) {
	RoomID := c.Param("id")

	err := h.RoomService.DeleteExhibitionRoomByID(c.Request.Context(), RoomID)
	if err != nil {
		log.Printf("Error deleting exhibitionRoom %s: %v", RoomID, err)

		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			c.JSON(http.StatusNotFound, gin.H{"error": "Exhibition Room not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"_id": RoomID + " has been deleted."})
}
