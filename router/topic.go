package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type TopicRouter struct{
	Router *mux.Router
	Handler *handler.Handler
}

func StartTopicsRoute(mainRoute *mux.Router, handler *handler.Handler){
	TopicRoute := mainRoute.PathPrefix("/topics").Subrouter()
	ur := TopicRouter{TopicRoute, handler}
	ur.HandleFunctions()
}


func (ur *TopicRouter) HandleFunctions() {
	// Topic CRUD
	ur.Router.HandleFunc("/getaall/", ur.Handler.GetTopics).Methods("GET")
	ur.Router.HandleFunc("/{id}", ur.Handler.GetTopicByID).Methods("GET")
	ur.Router.HandleFunc("/create", ur.Handler.CreateTopic).Methods("POST")
	ur.Router.HandleFunc("/update/{id}", ur.Handler.UpdateTopic).Methods("PUT")
	ur.Router.HandleFunc("/delete/{id}", ur.Handler.DeleteTopic).Methods("DELETE")
}
