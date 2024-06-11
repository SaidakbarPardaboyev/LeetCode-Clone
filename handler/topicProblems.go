package handler

import (
	"encoding/json"
	"leetcode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Read
func (h *Handler) GetTopicProblems(c *gin.Context) {
	filter := model.TopicProblemFilter{}
	TopicId, hasTopicId := c.GetQuery("topic_id")
	if hasTopicId {
		filter.TopicId = &TopicId
	}
	ProblemId, hasProblemId := c.GetQuery("problem_id")
	if hasProblemId {
		filter.ProblemId = &ProblemId
	}

	TopicProblems, err := h.TopicProblemRepo.GetTopicProblems(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting TopicProblems",
		})
		return
	}

	c.JSON(http.StatusOK, TopicProblems)
}

func (h *Handler) GetTopicProblemByID(c *gin.Context) {
	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error getting Id from params",
		})
	}

	TopicProblem, err := h.TopicProblemRepo.GetTopicProblemById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting TopicProblem by Id",
		})
		return
	}

	c.JSON(http.StatusOK, TopicProblem)
}

func (h *Handler) GetProblemsByTopicId(c *gin.Context) {
	topicId, hasTopicId := c.Params.Get("topic_id")
	if !hasTopicId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in params",
		})
		return
	}

	problems, err := h.TopicProblemRepo.GetTopicProblemById(topicId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting problems by topic_id",
		})
		return
	}

	c.JSON(http.StatusOK, problems)
}

func (h *Handler) GetTopicsByProblemId(c *gin.Context) {
	problemId, hasProblemId := c.Params.Get("problem_id")
	if !hasProblemId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in params",
		})
		return
	}

	problems, err := h.TopicProblemRepo.GetTopicsByProblemId(problemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting topics by problem_id",
		})
		return
	}

	c.JSON(http.StatusOK, problems)
}

// create
func (h *Handler) CreateTopicProblem(c *gin.Context) {
	newTopicProblem := model.TopicProblem{}
	err := json.NewDecoder(c.Request.Body).Decode(&newTopicProblem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding TopicProblem",
		})
		return
	}
	err = h.TopicProblemRepo.CreateTopicProblem(&newTopicProblem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating TopicProblem",
		})
		return
	}
}

// Update
func (h *Handler) UpdateTopicProblem(c *gin.Context) {
	TopicProblem := model.TopicProblem{}

	err := json.NewDecoder(c.Request.Body).Decode(&TopicProblem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding TopicProblem",
		})
		return
	}

	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in params",
		})
		return
	}
	TopicProblem.Id = id
	err = h.TopicProblemRepo.UpdateTopicProblem(&TopicProblem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while updating TopicProblem",
		})
		return
	}
	c.JSON(http.StatusOK, "Updated Succesfully")
}

// Delete
func (h *Handler) DeleteTopicProblem(c *gin.Context) {

	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in params",
		})
		return
	}
	err := h.TopicProblemRepo.DeleteTopicProblem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while deleting TopicProblem",
		})
		return
	}
	c.JSON(http.StatusOK, "Deleted Succesfully")
}
