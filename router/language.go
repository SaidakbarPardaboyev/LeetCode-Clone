package router

import (
	"leetcode/handler"

	"github.com/gin-gonic/gin"
)

type LanguageRouter struct {
	Router  *gin.RouterGroup
	Handler *handler.Handler
}

func NewLanguageRouter(l *gin.RouterGroup, handler *handler.Handler) *LanguageRouter {
	return &LanguageRouter{Router: l, Handler: handler}
}

func StartLanguagesRoute(mainRouter *gin.RouterGroup, handler *handler.Handler) {
	// Language CRUD
	languageGroup := mainRouter.Group("/language")
	r := NewLanguageRouter(languageGroup, handler)

	r.Router.GET("/get", r.Handler.GetLanguages)
	r.Router.GET("/{id}", r.Handler.GetLanguageByID)
	r.Router.POST("/create", r.Handler.CreateLanguage)
	r.Router.PUT("/update/{id}", r.Handler.UpdateLanguage)
	r.Router.DELETE("/delete/{id}", r.Handler.DeleteLanguage)
}
