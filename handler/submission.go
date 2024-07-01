package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"leetcode/models"
	"log"
	"net/http"
	"strconv"
)

// Read
func (h *Handler) GetSubmissions(c *gin.Context) {
	filter := models.SubmissionFilter{}
	problemId := c.Query("problem_id")
	if len(problemId) != 32 {
		filter.ProblemId = &problemId
	}

	userId := c.Query("user_id")
	if len(userId) != 32 {
		filter.UserId = &userId
	}

	languageId := c.Query("language_id")
	if len(languageId) != 32 {
		filter.LanguageId = &languageId
	}

	code, hasCode := c.GetQuery("code")
	if hasCode {
		filter.Code = &code
	}

	submissionStatus, hasSubmissionStatus := c.GetQuery("submission_status")
	if hasSubmissionStatus {
		filter.SubmissionStatus = &submissionStatus
	}

	limit, hasLimit := c.GetQuery("limit")
	if hasLimit {
		l, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "StatusBadRequest",
				"measage": "Error while converting limit",
			})
			log.Println("Error while converting limit ", err)
			return
		}
		filter.Limit = &l
	}

	offset, hasOffset := c.GetQuery("offset")
	if hasOffset {
		of, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "StatusBadRequest",
				"measage": "Error while converting offset",
			})
			log.Println("Error while converting offset ", err)
			return
		}
		filter.Limit = &of
	}

	submissions, err := h.SubmissionRepo.GetSubmissions(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting submissions",
		})
		log.Println("Error getting submissions ", err)
		return
	}

	c.JSON(http.StatusOK, submissions)
}

func (h *Handler) GetSubmissionByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: id not found or invalid id in params",
		})
		log.Println("Error: id not found or invalid id in params ", id)
		return
	}

	submission, err := h.SubmissionRepo.GetSubmissionById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting submission by Id",
		})
		log.Println("Error getting submission by Id ", err)
		return
	}

	c.JSON(http.StatusOK, submission)
}

func (h *Handler) GetActiveDays(c *gin.Context) {
	userId := c.Param("user_id")
	if len(userId) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: user_id not found or invalid user_id in params",
		})
		log.Println("Error: user_id not found or invalid user_id in params ", userId)
		return
	}

	year := c.Param("year")
	if len(year) <= 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: year not found or invalid year in params",
		})
		log.Println("Error: year not found or invalid year in params ", year)
		return
	}

	y, err := strconv.Atoi(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error while converting year",
		})
		log.Println("Error while converting year ", err)
		return
	}

	activity, err := h.SubmissionRepo.GetActiveDays(userId, y)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting active days",
		})
		log.Println("Error getting sactive days ", err)
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (h *Handler) GetRecentlyAcceptedSubmissionsByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	if len(userId) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: user_id not found or invalid user_id in params",
		})
		log.Println("Error: user_id not found or invalid user_id in params ", userId)
		return
	}

	recentAC, err := h.SubmissionRepo.GetRecentlyAcceptedSubmissionsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting RecentlyAcceptedSubmissions by user_id",
		})
		log.Println("Error getting RecentlyAcceptedSubmissions by user_id ", err)
		return
	}

	c.JSON(http.StatusOK, recentAC)
}

func (h *Handler) GetLastSubmittedCodeByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	if len(userId) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: user_id not found or invalid user_id in params",
		})
		log.Println("Error: user_id not found or invalid user_id in params ", userId)
		return
	}
	problemId := c.Param("problem_id")
	if len(problemId) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: problem_id not found or invalid problem_id in params",
		})
		log.Println("Error: problem_id not found or invalid problem_id in params ", userId)
		return
	}

	code, err := h.SubmissionRepo.GetLastSubmittedCodeByUserId(userId, problemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting GetLastSubmittedCodeByUserId by user_id",
		})
		log.Println("Error getting GetLastSubmittedCodeByUserId by user_id ", err)
		return
	}

	c.JSON(http.StatusOK, code)
}

// Create
func (h *Handler) CreateSubmission(c *gin.Context) {
	newSubmission := models.CreateSubmission{}
	err := json.NewDecoder(c.Request.Body).Decode(&newSubmission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error while decoding submission",
		})
		log.Println("Error while decoding submission ", err)
		return
	}

	id, err := h.SubmissionRepo.CreateSubmission(&newSubmission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error while creating submission",
		})
		log.Println("Error while creating submission ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

// Update
func (h *Handler) UpdateSubmission(c *gin.Context) {
	submission := models.UpdateSubmission{}

	err := json.NewDecoder(c.Request.Body).Decode(&submission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error while decoding submission",
		})
		log.Println("Error while decoding submission ", err)
		return
	}

	id := c.Param("id")
	if len(id) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: id not found in params",
		})
		log.Println("Error: id not found in params ", id)
		return
	}
	submission.Id = id

	err = h.SubmissionRepo.UpdateSubmission(&submission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error while updating submission",
		})
		log.Println("Error while updating submission ", err)
		return
	}
	c.JSON(http.StatusOK, "Updated Succesfully")
}

// Delete
func (h *Handler) DeleteSubmission(c *gin.Context) {

	id := c.Param("id")
	if len(id) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: id not found in params",
		})
		log.Println("Error: id not found in params ", id)
		return
	}

	err := h.SubmissionRepo.DeleteSubmission(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error while deleting Submission",
		})
		log.Println("Error while deleting Submission ", err)
		return
	}
	c.JSON(http.StatusOK, "Deleted Succesfully")
}
