package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type TopicProblemRouter struct{
	Router *mux.Router
	Handler *handler.Handler
}

func StartTopicProblemsRoute(mainRoute *mux.Router, handler *handler.Handler){
	TopicProblemRoute := mainRoute.PathPrefix("/topicproblems").Subrouter()
	ur := TopicProblemRouter{TopicProblemRoute, handler}
	ur.HandleFunctions()
}


func (ur *TopicProblemRouter) HandleFunctions() {
	// TopicProblem CRUD
	ur.Router.HandleFunc("/getall/", ur.Handler.GetTopicProblems).Methods("GET")
	ur.Router.HandleFunc("/{id}", ur.Handler.GetTopicProblemByID).Methods("GET")
	ur.Router.HandleFunc("/create", ur.Handler.CreateTopicProblem).Methods("POST")
	ur.Router.HandleFunc("/update/{id}", ur.Handler.UpdateTopicProblem).Methods("PUT")
	ur.Router.HandleFunc("/delete/{id}", ur.Handler.DeleteTopicProblem).Methods("DELETE")
}
