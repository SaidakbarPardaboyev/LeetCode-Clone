package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Read
func (h *Handler) GetLanguages(c *gin.Context) {
	filter := model.LanguageFilter{}
	if name, has := c.GetQuery("name"); has {
		filter.Name = &name
	}

	languages, err := h.LanguageRepo.GetLanguages(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting languages by filter",
		})
		log.Println("Error getting languages by filter", err)
		return
	}

	c.JSON(200, languages)

}

func (h *Handler) GetLanguageByID(c *gin.Context) {
	id, hasKey := c.GetQuery("id")
	if !hasKey {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id sent",
		})
		log.Println("Error no Id sent")
		return
	}

	language, err := h.LanguageRepo.GetLanguageById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting language by Id",
		})
		log.Println("Error getting language by Id", err)
		return
	}
	c.JSON(200, language)
}

// Create
func (h *Handler) CreateLanguage(c *gin.Context) {
	newlanguage := model.Language{}
	err := json.NewDecoder(c.Request.Body).Decode(&newlanguage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding language",
		})
		log.Println("Error while decoding language", err)
		return
	}
	err = h.LanguageRepo.CreateLanguage(&newlanguage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating language",
		})
		log.Println("Error while creating language", err)
		return
	}
}

// Update
func (h *Handler) UpdateLanguage(c *gin.Context) {
	language := model.Language{}
	id, hasKey := c.GetQuery("id")
	if !hasKey {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id sent",
		})
		log.Println("Error no Id sent")
		return
	}
	language.Id = id
	err := json.NewDecoder(c.Request.Body).Decode(&language)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding language",
		})
		log.Println("Error while decoding language", err)
		return
	}

	err = h.LanguageRepo.UpdateLanguage(&language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while updating language",
		})
		log.Println("Error while updating language", err)
		return
	}
}

// Delete
func (h *Handler) DeleteLanguage(c *gin.Context) {

	id, hasKey := c.GetQuery("id")
	if !hasKey {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id sent",
		})
		log.Println("Error no Id sent")
		return
	}
	err := h.LanguageRepo.DeleteLanguage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while deleting language",
		})
		log.Println("Error while deleting language", err)
		return
	}
}
