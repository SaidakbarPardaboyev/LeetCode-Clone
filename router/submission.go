package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type SubmissionRouter struct {
	Router  *mux.Router
	Handler *handler.Handler
}

func NewSubmissionsRouter(sb *mux.Router, handler *handler.Handler) *SubmissionRouter {
	return &SubmissionRouter{Router: sb, Handler: handler}
}

func StartSubmissionsRoute(mainRouter *mux.Router, handler *handler.Handler) {
	// Submission CRUD
	sb := mainRouter.PathPrefix("/Submissions").Subrouter()
	r := NewSubmissionsRouter(sb, handler)

	r.Router.HandleFunc("/getall/", r.Handler.GetSubmissions).Methods("GET")
	r.Router.HandleFunc("/{id}", r.Handler.GetSubmissionByID).Methods("GET")
	r.Router.HandleFunc("/getsubmissions", r.Handler.GetSubmissionsOfUserForProblem).Methods("GET")
	r.Router.HandleFunc("/getrecentac", r.Handler.GetRecentAcceptedSubmissions).Methods("GET")
	r.Router.HandleFunc("/create", r.Handler.CreateSubmission).Methods("POST")
	r.Router.HandleFunc("/update/{id}", r.Handler.UpdateSubmission).Methods("PUT")
	r.Router.HandleFunc("/delete/{id}", r.Handler.DeleteSubmission).Methods("DELETE")
}
