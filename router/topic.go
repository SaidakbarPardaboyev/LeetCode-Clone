package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type TopicRouter struct {
	Router  *mux.Router
	Handler *handler.Handler
}

func NewTopicsRouter(tr *mux.Router, handler *handler.Handler) *TopicRouter {
	return &TopicRouter{Router: tr, Handler: handler}
}

func StartTopicsRoute(mainRouter *mux.Router, handler *handler.Handler) {
	// Topic CRUD
	tr := mainRouter.PathPrefix("/Topics").Subrouter()
	r := NewTopicsRouter(tr, handler)

	r.Router.HandleFunc("/getaall/", r.Handler.GetTopics).Methods("GET")
	r.Router.HandleFunc("/{id}", r.Handler.GetTopicByID).Methods("GET")
	r.Router.HandleFunc("/create", r.Handler.CreateTopic).Methods("POST")
	r.Router.HandleFunc("/update/{id}", r.Handler.UpdateTopic).Methods("PUT")
	r.Router.HandleFunc("/delete/{id}", r.Handler.DeleteTopic).Methods("DELETE")
}
