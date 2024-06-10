package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type LanguageRouter struct{
	Router *mux.Router
	Handler *handler.Handler
}

func StartLanguagesRoute(mainRoute *mux.Router, handler *handler.Handler){
	LanguageRoute := mainRoute.PathPrefix("/Languages").Subrouter()
	ur := LanguageRouter{LanguageRoute, handler}
	ur.HandleFunctions()
}


func (ur *LanguageRouter) HandleFunctions() {
	// Language CRUD
	ur.Router.HandleFunc("/getall/", ur.Handler.GetLanguages).Methods("GET")
	ur.Router.HandleFunc("/{id}", ur.Handler.GetLanguageByID).Methods("GET")
	ur.Router.HandleFunc("/create", ur.Handler.CreateLanguage).Methods("POST")
	ur.Router.HandleFunc("/update/{id}", ur.Handler.UpdateLanguage).Methods("PUT")
	ur.Router.HandleFunc("/delete/{id}", ur.Handler.DeleteLanguage).Methods("DELETE")
}
