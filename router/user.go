package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type UserRouter struct{
	Router *mux.Router
	Handler *handler.Handler
}

func StartUsersRoute(mainRoute *mux.Router, handler *handler.Handler){
	userRoute := mainRoute.PathPrefix("/users").Subrouter()
	ur := UserRouter{userRoute, handler}
	ur.HandleFunctions()
}


func (ur *UserRouter) HandleFunctions() {
	// User CRUD
	ur.Router.HandleFunc("/getall/", ur.Handler.GetUsers).Methods("GET")
	ur.Router.HandleFunc("/{id}", ur.Handler.GetUserByID).Methods("GET")
	ur.Router.HandleFunc("/create", ur.Handler.CreateUser).Methods("POST")
	ur.Router.HandleFunc("/update/{id}", ur.Handler.UpdateUser).Methods("PUT")
	ur.Router.HandleFunc("/delete/{id}", ur.Handler.DeleteUser).Methods("DELETE")
}
