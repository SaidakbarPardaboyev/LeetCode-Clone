package router

import (
	"leetcode/handler"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Router  *gin.RouterGroup
	Handler *handler.Handler
}

func NewUsersRouter(ur *gin.RouterGroup, handler *handler.Handler) *UserRouter {
	return &UserRouter{Router: ur, Handler: handler}
}

func StartUsersRoute(LeetcodeGroup *gin.RouterGroup, handler *handler.Handler) {
	// User CRUD
	ur := LeetcodeGroup.Group("/users")
	r := NewUsersRouter(ur, handler)

	r.Router.GET("/getall/", r.Handler.GetUsers)
	r.Router.GET("/:id", r.Handler.GetUserByID)
	r.Router.POST("/create", r.Handler.CreateUser)
	r.Router.PUT("/update/:id", r.Handler.UpdateUser)
	r.Router.DELETE("/delete/:id", r.Handler.DeleteUser)
}
