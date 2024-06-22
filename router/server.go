package router

import (
	"leetcode/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(handler *handler.Handler) *http.Server {
	router := gin.Default()
	LeetcodeGroup := router.Group("leetcode.uz")

	StartUsersRoute(LeetcodeGroup, handler)
	StartLanguagesRoute(LeetcodeGroup, handler)
	StartTopicsRoute(LeetcodeGroup, handler)
	StartSubmissionsRoute(LeetcodeGroup, handler)

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

}
