package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	Router  *mux.Router
	Handler *handler.Handler
}

func NewUsersRouter(ur *mux.Router, handler *handler.Handler) *UserRouter {
	return &UserRouter{Router: ur, Handler: handler}
}

func StartUsersRoute(LeetcodeRouter *mux.Router, handler *handler.Handler) {
	// User CRUD
	ur := LeetcodeRouter.PathPrefix("/Users").Subrouter()
	r := NewUsersRouter(ur, handler)

	r.Router.HandleFunc("/getall/", r.Handler.GetUsers).Methods("GET")
	r.Router.HandleFunc("/{id}", r.Handler.GetUserByID).Methods("GET")
	r.Router.HandleFunc("/create", r.Handler.CreateUser).Methods("POST")
	r.Router.HandleFunc("/update/{id}", r.Handler.UpdateUser).Methods("PUT")
	r.Router.HandleFunc("/delete/{id}", r.Handler.DeleteUser).Methods("DELETE")
}
