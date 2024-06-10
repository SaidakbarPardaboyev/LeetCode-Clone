package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type SubmissionRouter struct{
	Router *mux.Router
	Handler *handler.Handler
}

func StartSubmissionsRoute(mainRoute *mux.Router, handler *handler.Handler){
	SubmissionRoute := mainRoute.PathPrefix("/submissions").Subrouter()
	ur := SubmissionRouter{SubmissionRoute, handler}
	ur.HandleFunctions()
}


func (ur *SubmissionRouter) HandleFunctions() {
	// Submission CRUD
	ur.Router.HandleFunc("/getall/", ur.Handler.GetSubmissions).Methods("GET")
	ur.Router.HandleFunc("/{id}", ur.Handler.GetSubmissionByID).Methods("GET")
	ur.Router.HandleFunc("/getsubmissions", ur.Handler.GetSubmissionsOfUserForProblem).Methods("GET")
	ur.Router.HandleFunc("/getrecentac", ur.Handler.GetRecentAcceptedSubmissions).Methods("GET")
	ur.Router.HandleFunc("/create", ur.Handler.CreateSubmission).Methods("POST")
	ur.Router.HandleFunc("/update/{id}", ur.Handler.UpdateSubmission).Methods("PUT")
	ur.Router.HandleFunc("/delete/{id}", ur.Handler.DeleteSubmission).Methods("DELETE")
}
