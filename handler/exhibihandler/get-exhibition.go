package exhibihandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetExhibitions godoc
//
//	@Summary		Get all exhibitions
//	@Description	Get a list of all exhibitions
//	@Tags			Exhibitions
//	@ID				GetAllExhibitions
//	@Produce		json
//	@Success		200	{object}	[]model.ResponseExhibition
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions [get]
func (h *Handler) GetAllExhibitions(c *gin.Context) {
	exhibitions, err := h.Service.GetAllExhibitions(c.Request.Context())
	if err != nil {
		log.Printf("Error retrieving exhibitions : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibitions)
}

// GetExhibition godoc
//
//	@Summary		Get exhibition by ID
//	@Description	Get exhibition details by ID
//	@Tags			Exhibitions
//	@ID				GetExhibitionByID
//	@Produce		json
//	@Param			id	path		string	true	"Exhibition ID"
//	@Success		200	{object}	model.ResponseExhibition
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/{id} [get]
func (h *Handler) GetExhibitionByID(c *gin.Context) {
	exhibitionID := c.Param("id")

	exhibition, err := h.Service.GetExhibitionByID(c.Request.Context(), exhibitionID)
	if err != nil {
		log.Printf("Error retrieving exhibition %s: %v", exhibitionID, err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibition)
}
