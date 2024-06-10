package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type TopicProblemRouter struct {
	Router  *mux.Router
	Handler *handler.Handler
}

func NewTopicsProblemsRouter(pt *mux.Router, handler *handler.Handler) *TopicProblemRouter {
	return &TopicProblemRouter{Router: pt, Handler: handler}
}

func StartTopicProblemsRoute(mainRouter *mux.Router, handler *handler.Handler) {
	// TopicProblem CRUD
	tp := mainRouter.PathPrefix("/ProblemsTopics").Subrouter()
	r := NewTopicsProblemsRouter(tp, handler)

	r.Router.HandleFunc("/getall/", r.Handler.GetTopicProblems).Methods("GET")
	r.Router.HandleFunc("/{id}", r.Handler.GetTopicProblemByID).Methods("GET")
	r.Router.HandleFunc("/create", r.Handler.CreateTopicProblem).Methods("POST")
	r.Router.HandleFunc("/update/{id}", r.Handler.UpdateTopicProblem).Methods("PUT")
	r.Router.HandleFunc("/delete/{id}", r.Handler.DeleteTopicProblem).Methods("DELETE")
}
