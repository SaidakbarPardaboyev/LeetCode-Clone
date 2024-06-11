package router

import (
	"leetcode/handler"

	"github.com/gin-gonic/gin"
)

type TopicRouter struct {
	Router  *gin.RouterGroup
	Handler *handler.Handler
}

func NewTopicsRouter(tr *gin.RouterGroup, handler *handler.Handler) *TopicRouter {
	return &TopicRouter{Router: tr, Handler: handler}
}

func StartTopicsRoute(mainRouter *gin.RouterGroup, handler *handler.Handler) {
	// Topic CRUD
	tr := mainRouter.Group("/Topics")
	r := NewTopicsRouter(tr, handler)

	r.Router.GET("/getaall/", r.Handler.GetTopics)
	r.Router.GET("/:id", r.Handler.GetTopicByID)
	r.Router.POST("/create", r.Handler.CreateTopic)
	r.Router.PUT("/update/:id", r.Handler.UpdateTopic)
	r.Router.DELETE("/delete/:id", r.Handler.DeleteTopic)
}
