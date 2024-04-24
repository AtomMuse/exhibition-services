package roomhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		Get exhibitionRoom by ID
//	@Description	Get exhibition data by RoomID
//	@Tags			Rooms
//
//	@Security		BearerAuth
//
//	@ID				GetExhibitionRoomByID
//	@Produce		json
//	@Param			id	path		string	true	"Exhibition Room ID"
//	@Success		200	{object}	model.ResponseExhibitionRoom
//	@Failure		401
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/rooms/{id} [get]
func (h *Handler) GetExhibitionRoomByID(c *gin.Context) {
	RoomID := c.Param("id")

	exhibitionRoom, err := h.RoomService.GetExhibitionRoomByID(c.Request.Context(), RoomID)
	if err != nil {
		log.Printf("Error retrieving exhibition Room  %s: %v", RoomID, err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibitionRoom)
}

//	@Summary		Get all exhibitions Rooms
//	@Description	Get a list of all exhibition Rooms data
//	@Tags			Rooms
//
//	@Security		BearerAuth
//
//	@ID				GetAllExhibitionRooms
//	@Produce		json
//	@Success		200	{object}	[]model.ResponseExhibitionRoom
//	@Failure		401
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/rooms/all [get]
func (h *Handler) GetAllExhibitionRooms(c *gin.Context) {
	exhibitionRooms, err := h.RoomService.GetAllExhibitionRooms(c.Request.Context())
	if err != nil {
		log.Printf("Error retrieving exhibitions : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibitionRooms)
}

//	@Summary		Get Rooms By exhibitionID
//	@Description	Get Rooms By exhibitionID
//	@Tags			Rooms
//	@Security		BearerAuth
//	@ID				GetRoomsByExhibitionID
//	@Produce		json
//	@Param			id	path		string	true	"Exhibition ID"
//	@Success		200	{object}	[]model.ResponseExhibitionRoom
//	@Failure		401
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/{id}/rooms [get]
func (h *Handler) GetRoomsByExhibitionID(c *gin.Context) {
	// Extract the exhibition ID from the request
	exhibitionID := c.Param("id")

	// Call the service to get Rooms by exhibition ID
	Rooms, err := h.RoomService.GetRoomsByExhibitionID(c.Request.Context(), exhibitionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the Rooms as JSON response
	c.JSON(http.StatusOK, Rooms)
}
