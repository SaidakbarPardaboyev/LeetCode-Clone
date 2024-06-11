package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProblems(c *gin.Context) {
	filter := model.ProblemFilter{}

	if qs, hasKey := c.GetQuery("question_number"); hasKey {
		qn, err := strconv.Atoi(qs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Error converting question_number",
			})
			log.Println("Error converting question_number:", err)
			return
		}
		filter.QuestionNumber = &qn
	}
	if title, hasKey := c.GetQuery("title"); hasKey {
		filter.Title = &title
	}
	if difficultyLevel, hasKey := c.GetQuery("difficulty_level"); hasKey {

		filter.DifficultyLevel = &difficultyLevel
	}

	problems, err := h.ProblemRepo.GetProblems(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error getting problems by filter",
		})
		log.Println("Error getting problems by filter", err)
		return
	}
	c.JSON(200, problems)

}

func (h *Handler) GetProblemByID(c *gin.Context) {
	id, hasKey := c.GetQuery("id")
	if !hasKey {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id sent",
		})
		log.Println("Error no Id sent")
		return
	}

	problem, err := h.ProblemRepo.GetProblemById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while getting problem by id",
		})
		log.Println("Error getting problem by Id", err)
		return
	}
	c.JSON(200, problem)
}

func (h *Handler) CreateProblem(c *gin.Context) {
	newproblem := model.Problem{}
	err := json.NewDecoder(c.Request.Body).Decode(&newproblem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding",
		})
		log.Println("Error while decoding problem", err)
		return
	}
	err = h.ProblemRepo.CreateProblem(&newproblem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while creating problem",
		})
		log.Println("Error while creating problem", err)
		return
	}
}

func (h *Handler) UpdateProblem(c *gin.Context) {
	problem := model.Problem{}

	err := json.NewDecoder(c.Request.Body).Decode(&problem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error while decoding problem",
		})
		log.Println("Error while decoding problem", err)
		return
	}

	id, hasId := c.GetQuery("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id",
		})
		log.Println("Error no Id")
		return
	}
	problem.Id = id
	err = h.ProblemRepo.UpdateProblem(&problem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while updating problem",
		})
		log.Println("Error while updating problem", err)
		return
	}
}

func (h *Handler) DeleteProblem(c *gin.Context) {

	id, hasId := c.GetQuery("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error no Id",
		})
		log.Println("Error no Id")
		return
	}
	err := h.ProblemRepo.DeleteProblem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": "Error while deleting problem",
		})
		log.Println("Error while deleting problem", err)
		return
	}
}
