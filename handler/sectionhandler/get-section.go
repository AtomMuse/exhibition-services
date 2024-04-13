package sectionhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		Get exhibitionSection by ID
//	@Description	Get exhibition data by sectionID
//	@Tags			Sections
//
//	@Security		BearerAuth
//
//	@ID				GetExhibitionSectionByID
//	@Produce		json
//	@Param			id	path		string	true	"Exhibition Section ID"
//	@Success		200	{object}	model.ResponseExhibitionSection
//	@Failure		401
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/sections/{id} [get]
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

//	@Summary		Get all exhibitions sections
//	@Description	Get a list of all exhibition sections data
//	@Tags			Sections
//
//	@Security		BearerAuth
//
//	@ID				GetAllExhibitionSections
//	@Param			id	path	string	true	"Exhibition ID"
//	@Produce		json
//	@Success		200	{object}	[]model.ResponseExhibitionSection
//	@Failure		401
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/sections/all [get]
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

//	@Summary		Get Sections By exhibitionID
//	@Description	Get Sections By exhibitionID
//	@Tags			Sections
//	@Security		BearerAuth
//	@ID				GetSectionsByExhibitionID
//	@Produce		json
//	@Success		200	{object}	[]model.ResponseExhibitionSection
//	@Failure		401
//	@Failure		500	{object}	helper.APIError	"Internal server error"
//	@Router			/api/exhibitions/{id}/sections [get]
func (h *Handler) GetSectionsByExhibitionID(c *gin.Context) {
	// Extract the exhibition ID from the request
	exhibitionID := c.Param("id")

	// Call the service to get sections by exhibition ID
	sections, err := h.SectionService.GetSectionsByExhibitionID(c.Request.Context(), exhibitionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the sections as JSON response
	c.JSON(http.StatusOK, sections)
}
