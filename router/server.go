package router

import (
	"leetcode/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateServer(handler *handler.Handler) *http.Server {
	LeetcodeRouter := mux.NewRouter().PathPrefix("/leetcode.uz").Subrouter()

	StartUsersRoute(LeetcodeRouter, handler)
	StartProblemsRoute(LeetcodeRouter, handler)
	StartLanguagesRoute(LeetcodeRouter, handler)
	StartTopicsRoute(LeetcodeRouter, handler)
	StartTopicProblemsRoute(LeetcodeRouter, handler)
	StartSubmissionsRoute(LeetcodeRouter, handler)

	return &http.Server{
		Addr:    ":8080",
		Handler: LeetcodeRouter,
	}
}
