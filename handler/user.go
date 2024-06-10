package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	filter := model.UserFilter{}
	query := r.URL.Query()
	if query.Has("full_name"){
		fullName := query.Get("full_name")
		filter.FullName = &fullName
	}
	if query.Has("username"){
		username := query.Get("username")
		filter.FullName = &username
	}

	users, err := h.UserRepo.GetUsers(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting users",err)
		return
	} 
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding users", err)
		return
	}

}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.UserRepo.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting user by Id", err)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding user", err)
		return
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := model.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding user", err)
		return
	}
	err = h.UserRepo.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while creating user", err)
		return
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	vars := mux.Vars(r)
	
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding user", err)
		return
	}

	user.Id = vars["id"]
	err = h.UserRepo.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while updating user", err)
		return
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	err := h.UserRepo.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while deleting user", err)
		return
	}
}
