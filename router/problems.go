package router

import (
	"leetcode/handler"

	"github.com/gin-gonic/gin"
)

type ProblemRouter struct {
	Router  *gin.RouterGroup
	Handler *handler.Handler
}

func NewProblemRouter(pr *gin.RouterGroup, handler *handler.Handler) *ProblemRouter {
	return &ProblemRouter{Router: pr, Handler: handler}
}

func StartProblemsRoute(mainRouter *gin.RouterGroup, handler *handler.Handler) {
	// Problem CRUD
	p := mainRouter.Group("/problems")
	r := NewProblemRouter(p, handler)

	r.Router.GET("/getall/", r.Handler.GetProblems)
	r.Router.GET("/:id", r.Handler.GetProblemByID)
	r.Router.POST("/create", r.Handler.CreateProblem)
	r.Router.PUT("/update/:id", r.Handler.UpdateProblem)
	r.Router.DELETE("/delete/:id", r.Handler.DeleteProblem)
}
