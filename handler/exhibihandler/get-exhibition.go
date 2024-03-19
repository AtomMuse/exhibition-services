package exhibihandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		Get all exhibitions
//	@Description	Get a list of all exhibitions data
//	@Tags			Exhibitions
//	@ID				GetAllExhibitions
//	@Produce		json
//	@Success		200	{object}	[]model.ResponseExhibition
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/all [get]
func (h *Handler) GetAllExhibitions(c *gin.Context) {
	exhibitions, err := h.ExhibitionService.GetAllExhibitions(c.Request.Context())
	if err != nil {
		log.Printf("Error retrieving exhibitions : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibitions)
}

//	@Summary		Get exhibition by ID
//	@Description	Get exhibition data by exhibitionID
//	@Tags			Exhibitions
//	@ID				GetExhibitionByID
//	@Produce		json
//	@Param			id	path		string	true	"Exhibition ID"
//	@Success		200	{object}	model.ResponseExhibition
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/{id} [get]
func (h *Handler) GetExhibitionByID(c *gin.Context) {
	exhibitionID := c.Param("id")

	exhibition, err := h.ExhibitionService.GetExhibitionByID(c.Request.Context(), exhibitionID)
	if err != nil {
		log.Printf("Error retrieving exhibition %s: %v", exhibitionID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Increment the visited number
	exhibition.VisitedNumber++

	// Update the visited number in the database
	if err := h.ExhibitionService.UpdateVisitedNumber(c.Request.Context(), exhibitionID, exhibition.VisitedNumber); err != nil {
		log.Printf("Error updating visited number for exhibition %s: %v", exhibitionID, err)
		// Handle the error accordingly
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibition)
}

// GetExhibitions godoc
//
//	@Summary		Get all exhibitions is public
//	@Description	Get a list of all exhibitions data is public only
//	@Tags			Exhibitions
//	@ID				GetExhibitionsIsPublic
//	@Produce		json
//	@Success		200	{object}	[]model.ResponseExhibition
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions [get]
func (h *Handler) GetExhibitionsIsPublic(c *gin.Context) {
	exhibitions, err := h.ExhibitionService.GetExhibitionsIsPublic(c.Request.Context())
	if err != nil {
		log.Printf("Error retrieving exhibitions : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibitions)
}
