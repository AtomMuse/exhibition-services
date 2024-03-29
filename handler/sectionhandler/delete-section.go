package sectionhandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//	@Summary		Delete Section by ID
//	@Description	Delete Section data by sectionID
//	@Tags			Sections
//	@ID				DeleteExhibitionSectionByID
//	@Produce		json
//	@Param			id	path		string							true	"Section ID"
//	@Success		200	{object}	model.ResponseGetExhibitionId	"Delete Section Success"
//	@Failure		500	{object}	helper.APIError					"Internal server error"
//	@Router			/api/sections/{id} [delete]
func (h *Handler) DeleteExhibitionSectionByID(c *gin.Context) {
	sectionID := c.Param("id")

	err := h.SectionService.DeleteExhibitionSectionByID(c.Request.Context(), sectionID)
	if err != nil {
		log.Printf("Error deleting exhibitionSection %s: %v", sectionID, err)

		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			c.JSON(http.StatusNotFound, gin.H{"error": "Exhibition section not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"_id": sectionID + " has been deleted."})
}
