package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type LanguageRouter struct {
	Router  *mux.Router
	Handler *handler.Handler
}

func NewLanguageRouter(l *mux.Router, handler *handler.Handler) *LanguageRouter {
	return &LanguageRouter{Router: l, Handler: handler}
}

func StartLanguagesRoute(mainRouter *mux.Router, handler *handler.Handler) {
	// Language CRUD
	l := mainRouter.PathPrefix("/Language").Subrouter()
	r := NewLanguageRouter(l, handler)

	r.Router.HandleFunc("/get", r.Handler.GetLanguages).Methods("GET")
	r.Router.HandleFunc("/{id}", r.Handler.GetLanguageByID).Methods("GET")
	r.Router.HandleFunc("/create", r.Handler.CreateLanguage).Methods("POST")
	r.Router.HandleFunc("/update/{id}", r.Handler.UpdateLanguage).Methods("PUT")
	r.Router.HandleFunc("/delete/{id}", r.Handler.DeleteLanguage).Methods("DELETE")
}
