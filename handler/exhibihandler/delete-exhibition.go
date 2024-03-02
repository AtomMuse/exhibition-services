package exhibihandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteExhibition godoc
//
//	@Summary		Delete exhibition by ID
//	@Description	Delete exhibition by ID
//	@Tags			Exhibitions
//	@ID				DeleteExhibition
//	@Produce		json
//	@Param			id	path		string							true	"Exhibition ID"
//	@Success		200	{object}	model.ResponseGetExhibitionId	"Delete Exhibition Success"
//	@Failure		500	{object}	helper.APIError					"Internal server error"
//	@Router			/api/exhibitions/{id} [delete]
func (h *Handler) DeleteExhibition(c *gin.Context) {
	exhibitionID := c.Param("id")

	err := h.ExhibitionService.DeleteExhibition(c.Request.Context(), exhibitionID)
	if err != nil {
		log.Printf("Error deleting exhibition %s: %v", exhibitionID, err)

		// Check for specific errors and return appropriate responses
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Exhibition not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Return a successful response with no content (HTTP 204 No Content)
	c.JSON(http.StatusOK, gin.H{"_id": exhibitionID + " has been deleted."})
}

// @Summary		Delete Section by ID
// @Description	Delete Section by ID
// @Tags			Exhibitions
// @ID				DeleteExhibitionSectionByID
// @Produce		json
// @Param			id	path		string							true	"Section ID"
// @Success		200	{object}	model.ResponseGetExhibitionId	"Delete Section Success"
// @Failure		500	{object}	helper.APIError					"Internal server error"
// @Router			/api/sections/{id} [delete]
func (h *Handler) DeleteExhibitionSectionByID(c *gin.Context) {
	sectionID := c.Param("id")

	err := h.SectionService.DeleteExhibitionSectionByID(c.Request.Context(), sectionID)
	if err != nil {
		log.Printf("Error deleting exhibitionSection %s: %v", sectionID, err)

		// Check for specific errors and return appropriate responses
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Exhibition not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Return a successful response with no content (HTTP 204 No Content)
	c.JSON(http.StatusOK, gin.H{"_id": sectionID + " has been deleted."})
}
