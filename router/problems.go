package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type ProblemRouter struct {
	Router  *mux.Router
	Handler *handler.Handler
}

func NewProblemRouter(pr *mux.Router, handler *handler.Handler) *ProblemRouter {
	return &ProblemRouter{Router: pr, Handler: handler}
}

func StartProblemsRoute(mainRouter *mux.Router, handler *handler.Handler) {
	// Problem CRUD
	p := mainRouter.PathPrefix("/Problems").Subrouter()
	r := NewProblemRouter(p, handler)

	r.Router.HandleFunc("/getall/", r.Handler.GetProblems).Methods("GET")
	r.Router.HandleFunc("/{id}", r.Handler.GetProblemByID).Methods("GET")
	r.Router.HandleFunc("/create", r.Handler.CreateProblem).Methods("POST")
	r.Router.HandleFunc("/update/{id}", r.Handler.UpdateProblem).Methods("PUT")
	r.Router.HandleFunc("/delete/{id}", r.Handler.DeleteProblem).Methods("DELETE")
}
