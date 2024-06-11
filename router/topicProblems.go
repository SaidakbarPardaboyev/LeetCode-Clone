package router

import (
	"leetcode/handler"

	"github.com/gin-gonic/gin"
)

type TopicProblemRouter struct {
	Router  *gin.RouterGroup
	Handler *handler.Handler
}

func NewTopicsProblemsRouter(pt *gin.RouterGroup, handler *handler.Handler) *TopicProblemRouter {
	return &TopicProblemRouter{Router: pt, Handler: handler}
}

func StartTopicProblemsRoute(mainRouter *gin.RouterGroup, handler *handler.Handler) {
	// TopicProblem CRUD
	tp := mainRouter.Group("/ProblemsTopics")
	r := NewTopicsProblemsRouter(tp, handler)

	r.Router.GET("/getall/", r.Handler.GetTopicProblems)
	r.Router.GET("/:id", r.Handler.GetTopicProblemByID)
	r.Router.GET("/:id", r.Handler.GetProblemsByTopicId)
	r.Router.GET("/:id", r.Handler.GetTopicsByProblemId)
	r.Router.POST("/create", r.Handler.CreateTopicProblem)
	r.Router.PUT("/update/:id", r.Handler.UpdateTopicProblem)
	r.Router.DELETE("/delete/:id", r.Handler.DeleteTopicProblem)
}
