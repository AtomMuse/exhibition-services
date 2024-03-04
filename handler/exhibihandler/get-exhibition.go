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
//	@Router			/api/all-exhibitions [get]
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

	exhibition, err := h.ExhibitionService.GetExhibitionByID(c.Request.Context(), exhibitionID)
	if err != nil {
		log.Printf("Error retrieving exhibition %s: %v", exhibitionID, err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibition)
}

// GetExhibitions godoc
//
//	@Summary		Get all exhibitions is public
//	@Description	Get a list of all exhibitions is public
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

// @Summary		Get exhibitionSection by ID
// @Description	Get exhibition details by ID
// @Tags			Exhibitions
// @ID				GetExhibitionSectionByID
// @Produce		json
// @Param			id	path		string	true	"Exhibition Section ID"
// @Success		200	{object}	model.ResponseExhibitionSection
// @Failure		500	{object}	helper.APIError	"Internal server error"
// @Router			/api/exhibitionSection/{id} [get]
func (h *Handler) GetExhibitionSectionByID(c *gin.Context) {
	sectionID := c.Param("id")

	exhibitionSection, err := h.SectionService.GetExhibitionSectionByID(c.Request.Context(), sectionID)
	if err != nil {
		log.Printf("Error retrieving exhibition section  %s: %v", sectionID, err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibitionSection)
}

// @Summary		Get all exhibitions sections
// @Description	Get a list of all exhibition sections
// @Tags			Exhibitions
// @ID				GetAllExhibitionSections
// @Produce		json
// @Success		200	{object}	[]model.ResponseExhibitionSection
// @Failure		500	{object}	helper.APIError	"Internal server error"
// @Router			/api/all-sections [get]
func (h *Handler) GetAllExhibitionSections(c *gin.Context) {
	exhibitionSections, err := h.SectionService.GetAllExhibitionSections(c.Request.Context())
	if err != nil {
		log.Printf("Error retrieving exhibitions : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the exhibition details
	c.JSON(http.StatusOK, exhibitionSections)
}
