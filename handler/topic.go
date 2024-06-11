package handler

import (
	"encoding/json"
	"leetcode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTopics(c *gin.Context) {
	filter := model.TopicFilter{}
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
	id, hasId := c.Params.Get("id")

	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Id URL da berilmagan",
		})
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
	newtopic := model.Topic{}
	err := json.NewDecoder(c.Request.Body).Decode(&newtopic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding topic",
		})
		return
	}
	err = h.TopicRepo.CreateTopic(&newtopic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating topic",
		})
		return
	}
	c.JSON(http.StatusOK, "Created Succesfully")
}

func (h *Handler) UpdateTopic(c *gin.Context) {
	topic := model.Topic{}

	err := json.NewDecoder(c.Request.Body).Decode(&topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding topic",
		})
		return
	}

	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating topic",
		})
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

	params := c.Params
	id, hasId := params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: Id is not found in URL",
		})
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
