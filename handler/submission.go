package handler

import (
	"encoding/json"
	"leetcode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Read
func (h *Handler) GetSubmissions(c *gin.Context) {
	filter := model.SubmissionFilter{}
	problemId, hasProblemId := c.GetQuery("problem_id")
	if hasProblemId {
		filter.ProblemId = &problemId
	}

	userId, hasUserId := c.GetQuery("user_id")
	if hasUserId {
		filter.UserId = &userId
	}

	languageId, hasLanguageId := c.GetQuery("language_id")
	if hasLanguageId {
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

	submissions, err := h.SubmissionRepo.GetSubmissions(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting Submissions",
		})
		return
	}

	c.JSON(http.StatusOK, submissions)
}

func (h *Handler) GetSubmissionByID(c *gin.Context) {
	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: id not found in params",
		})
		return
	}

	submission, err := h.SubmissionRepo.GetSubmissionById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting Submission by Id",
		})
		return
	}

	c.JSON(http.StatusOK, submission)
}

func (h *Handler) GetSubmissionsOfUserForProblem(c *gin.Context) {
	userId, hasUserId := c.GetQuery("user_id")
	if !hasUserId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: userId not found in params",
		})
		return
	}

	problemId, hasProblemId := c.GetQuery("problem_id")
	if !hasProblemId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: problemId not found in params",
		})
		return
	}

	submissions, err := h.SubmissionRepo.GetSubmissionsOfUserForProblem(userId, problemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting Submission by Id",
		})
		return
	}
	c.JSON(http.StatusOK, submissions)
}

func (h *Handler) GetRecentAcceptedSubmissions(c *gin.Context) {
	userId, hasUserId := c.Params.Get("user_id")
	if !hasUserId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: userId not found in params",
		})
		return
	}

	submissions, err := h.SubmissionRepo.GetRecentAcceptedSubmissions(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error getting Submission by Id",
		})
		return
	}

	c.JSON(http.StatusOK, submissions)
}

// Create
func (h *Handler) CreateSubmission(c *gin.Context) {
	newSubmission := model.Submission{}
	err := json.NewDecoder(c.Request.Body).Decode(&newSubmission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error while decoding Submission",
		})
		return
	}

	err = h.SubmissionRepo.CreateSubmission(&newSubmission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error while creating Submission",
		})
		return
	}
}

// Update
func (h *Handler) UpdateSubmission(c *gin.Context) {
	Submission := model.Submission{}

	err := json.NewDecoder(c.Request.Body).Decode(&Submission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error while decoding Submission",
		})
		return
	}

	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: id not found in params",
		})
		return
	}
	Submission.Id = id

	err = h.SubmissionRepo.UpdateSubmission(&Submission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error while updating Submission",
		})
		return
	}
	c.JSON(http.StatusOK, "Updated Succesfully")
}

// Delete
func (h *Handler) DeleteSubmission(c *gin.Context) {

	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"measage": "Error: id not found in params",
		})
		return
	}

	err := h.SubmissionRepo.DeleteSubmission(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"measage": "Error while deleting Submission",
		})
		return
	}	
	c.JSON(http.StatusOK, "Deleted Succesfully")
}