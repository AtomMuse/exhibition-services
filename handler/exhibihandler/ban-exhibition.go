package exhibihandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		BanExhibition
//	@Description	BanExhibition by exhibitionID
//	@Tags			Ban
//	@ID				BanExhibition
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Exhibition ID"
//	@Success		200	{object}	model.ResponseExhibition
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/{id}/ban [post]
func (h *Handler) BanExhibition(c *gin.Context) {
	exhibitionID := c.Param("id")
	if err := h.ExhibitionService.BanExhibition(c.Request.Context(), exhibitionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to ban exhibition"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Exhibition banned successfully"})
}
