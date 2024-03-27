package exhibihandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		Like exhibition by ID
//	@Description	Like exhibition by exhibitionID
//	@Tags			Like & Unlike
//
//	@Security		BearerAuth
//
//	@ID				LikeExhibition
//	@Produce		json
//	@Param			id	path		string	true	"Exhibition ID"
//	@Success		200	{object}	model.ResponseExhibition
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/{id}/like [put]
func (h *Handler) LikeExhibition(c *gin.Context) {
	exhibitionID := c.Param("id")

	if err := h.ExhibitionService.LikeExhibition(c.Request.Context(), exhibitionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like exhibition"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exhibition liked successfully"})
}

//	@Summary		Unlike exhibition by ID
//	@Description	unlike exhibition by exhibitionID
//	@Tags			Like & Unlike
//
//	@Security		BearerAuth
//
//	@ID				UnlikeExhibition
//	@Produce		json
//	@Param			id	path		string	true	"Exhibition ID"
//	@Success		200	{object}	model.ResponseExhibition
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/{id}/unlike [put]
func (h *Handler) UnlikeExhibition(c *gin.Context) {
	exhibitionID := c.Param("id")

	if err := h.ExhibitionService.UnlikeExhibition(c.Request.Context(), exhibitionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike exhibition"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exhibition unliked successfully"})
}
