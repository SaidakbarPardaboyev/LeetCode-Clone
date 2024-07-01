package handler

import (
	"encoding/json"
	"leetcode/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTopics(c *gin.Context) {
	filter := models.TopicFilter{}
	name, hasName := c.GetQuery("name")
	if hasName {
		filter.Name = &name
	}

	topics, err := h.TopicRepo.GetTopics(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting topics",
		})
		return
	}
	c.JSON(http.StatusOK, topics)
}

func (h *Handler) GetTopicByID(c *gin.Context) {

	id := c.Query("id")
	if len(id) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id sent or invalid id",
		})
		log.Println("Error no Id sent or invalid id ", id)
		return
	}

	topic, err := h.TopicRepo.GetTopicById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting topic by Id",
		})
		return
	}

	c.JSON(http.StatusOK, topic)
}

func (h *Handler) CreateTopic(c *gin.Context) {
	newtopic := models.CreateTopic{}
	err := json.NewDecoder(c.Request.Body).Decode(&newtopic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding topic",
		})
		return
	}
	id, err := h.TopicRepo.CreateTopic(&newtopic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating topic",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *Handler) UpdateTopic(c *gin.Context) {
	topic := models.UpdateTopic{}

	err := json.NewDecoder(c.Request.Body).Decode(&topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding topic",
		})
		return
	}

	id := c.Query("id")
	if len(id) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id sent or invalid id",
		})
		log.Println("Error no Id sent or invalid id ", id)
		return
	}

	topic.Id = id
	err = h.TopicRepo.UpdateTopic(&topic)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while updating topic",
		})
		return
	}
	c.JSON(http.StatusOK, "Updated Succesfully")
}

func (h *Handler) DeleteTopic(c *gin.Context) {

	id := c.Query("id")
	if len(id) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id sent or invalid id",
		})
		log.Println("Error no Id sent or invalid id ", id)
		return
	}

	err := h.TopicRepo.DeleteTopic(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while deleting topic",
		})
		return
	}

	c.JSON(http.StatusOK, "Deleted Succesfully")
}
