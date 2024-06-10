package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

func CreateServer(handler *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	route := r.PathPrefix("/leetcode.uz").Subrouter()

	StartUsersRoute(route, handler)
	StartProblemsRoute(route, handler)
	StartLanguagesRoute(route, handler)
	StartTopicsRoute(route, handler)
	StartTopicProblemsRoute(route, handler)
	StartSubmissionsRoute(route, handler)

	return r
}
