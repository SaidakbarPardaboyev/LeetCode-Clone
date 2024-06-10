package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Read
func (h *Handler) GetLanguages(w http.ResponseWriter, r *http.Request) {
	filter := model.LanguageFilter{}
	query := r.URL.Query()
	if query.Has("name"){
		name := query.Get("name")
		filter.Name = &name
	}

	languages, err := h.LanguageRepo.GetLanguages(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting languages",err)
		return
	} 

	err = json.NewEncoder(w).Encode(languages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding languages", err)
		return
	}

}

func (h *Handler) GetLanguageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	language, err := h.LanguageRepo.GetLanguageById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting language by Id", err)
		return
	}
	err = json.NewEncoder(w).Encode(language)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding language", err)
		return
	}
}

// Create
func (h *Handler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	newlanguage := model.Language{}
	err := json.NewDecoder(r.Body).Decode(&newlanguage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding language", err)
		return
	}
	err = h.LanguageRepo.CreateLanguage(newlanguage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while creating language", err)
		return
	}
}

// Update
func (h *Handler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	language := model.Language{}
	vars := mux.Vars(r)
	
	err := json.NewDecoder(r.Body).Decode(&language)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding language", err)
		return
	}

	language.Id = vars["id"]
	err = h.LanguageRepo.UpdateLanguage(language)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while updating language", err)
		return
	}
}

// Delete
func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	err := h.LanguageRepo.DeleteLanguage(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while deleting language", err)
		return
	}
}
