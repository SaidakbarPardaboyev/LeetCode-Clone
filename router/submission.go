package router

import (
	"leetcode/handler"

	"github.com/gin-gonic/gin"
)

type SubmissionRouter struct {
	Router  *gin.RouterGroup
	Handler *handler.Handler
}

func NewSubmissionsRouter(sb *gin.RouterGroup, handler *handler.Handler) *SubmissionRouter {
	return &SubmissionRouter{Router: sb, Handler: handler}
}

func StartSubmissionsRoute(mainRouter *gin.RouterGroup, handler *handler.Handler) {
	// Submission CRUD
	sb := mainRouter.Group("/Submissions")
	r := NewSubmissionsRouter(sb, handler)

	r.Router.GET("/getall/", r.Handler.GetSubmissions)
	r.Router.GET("/:id", r.Handler.GetSubmissionByID)
	r.Router.GET("/getsubmissions", r.Handler.GetSubmissionsOfUserForProblem)
	r.Router.GET("/getrecentac/:id", r.Handler.GetRecentAcceptedSubmissions)
	r.Router.POST("/create", r.Handler.CreateSubmission)
	r.Router.PUT("/update/:id", r.Handler.UpdateSubmission)
	r.Router.DELETE("/delete/:id", r.Handler.DeleteSubmission)
}

