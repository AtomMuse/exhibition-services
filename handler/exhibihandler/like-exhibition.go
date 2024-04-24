package exhibihandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary		Like exhibition by ID
// @Description	Like exhibition by exhibitionID
// @Tags			Like & Unlike
//
// @Security		BearerAuth
//
// @ID				LikeExhibition
// @Produce		json
// @Param			id	path		string	true	"Exhibition ID"
// @Success		200	{object}	model.ResponseExhibition
// @Failure		500	{object}	helper.APIError	"Internal server error"
// @Router			/api/exhibitions/{id}/like [put]
func (h *Handler) LikeExhibition(c *gin.Context) {
	exhibitionID := c.Param("id")

	// Check if user_id exists in the context
	userIDInterface, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		return
	}

	// Convert userIDInterface to string
	userIDObjectID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is not a primitive.ObjectID"})
		return
	}

	// Convert ObjectID to string
	userID := userIDObjectID.Hex()

	if err := h.ExhibitionService.LikeExhibition(c, exhibitionID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like exhibition", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exhibition liked successfully"})
}

// @Summary		Unlike exhibition by ID
// @Description	unlike exhibition by exhibitionID
// @Tags			Like & Unlike
//
// @Security		BearerAuth
//
// @ID				UnlikeExhibition
// @Produce		json
// @Param			id	path		string	true	"Exhibition ID"
// @Success		200	{object}	model.ResponseExhibition
// @Failure		500	{object}	helper.APIError	"Internal server error"
// @Router			/api/exhibitions/{id}/unlike [put]
func (h *Handler) UnlikeExhibition(c *gin.Context) {
	exhibitionID := c.Param("id")

	// Check if user_id exists in the context
	userIDInterface, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		return
	}

	// Convert userIDInterface to string
	userIDObjectID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is not a primitive.ObjectID"})
		return
	}

	// Convert ObjectID to string
	userID := userIDObjectID.Hex()

	// Unlike the exhibition
	if err := h.ExhibitionService.UnlikeExhibition(c, exhibitionID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike exhibition", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exhibition unliked successfully"})
}
