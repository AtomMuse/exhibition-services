package exhibihandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//	@Summary		Delete exhibition by ID
//	@Description	Delete exhibition data by exhibitionID
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

		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			c.JSON(http.StatusNotFound, gin.H{"error": "Exhibition not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"_id": exhibitionID + " has been deleted."})
}
