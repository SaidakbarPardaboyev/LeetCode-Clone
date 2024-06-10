package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type ProblemRouter struct{
	Router *mux.Router
	Handler *handler.Handler
}

func StartProblemsRoute(mainRoute *mux.Router, handler *handler.Handler){
	ProblemRoute := mainRoute.PathPrefix("/problems").Subrouter()
	ur := ProblemRouter{ProblemRoute, handler}
	ur.HandleFunctions()
}


func (ur *ProblemRouter) HandleFunctions() {
	// Problem CRUD
	ur.Router.HandleFunc("/getall/", ur.Handler.GetProblems).Methods("GET")
	ur.Router.HandleFunc("/{id}", ur.Handler.GetProblemByID).Methods("GET")
	ur.Router.HandleFunc("/create", ur.Handler.CreateProblem).Methods("POST")
	ur.Router.HandleFunc("/update/{id}", ur.Handler.UpdateProblem).Methods("PUT")
	ur.Router.HandleFunc("/delete/{id}", ur.Handler.DeleteProblem).Methods("DELETE")
}
