package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetProblems(w http.ResponseWriter, r *http.Request) {
	filter := model.ProblemFilter{}
	query := r.URL.Query()
	if query.Has("question_number") {
		qn, err := strconv.Atoi(query.Get("question_number"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Error converting id", err)
			return
		}
		filter.QuestionNumber = &qn
	}
	if query.Has("title") {
		title := query.Get("title")
		filter.Title = &title
	}
	if query.Has("difficulty_level") {
		DifficultyLevel := query.Get("difficulty_level")
		filter.Title = &DifficultyLevel
	}

	problems, err := h.ProblemRepo.GetProblems(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting problems", err)
		return
	}
	err = json.NewEncoder(w).Encode(problems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding problems", err)
		return
	}

}

func (h *Handler) GetProblemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	problem, err := h.ProblemRepo.GetProblemById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting problem by Id", err)
		return
	}
	err = json.NewEncoder(w).Encode(problem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding problem", err)
		return
	}
}

func (h *Handler) CreateProblem(w http.ResponseWriter, r *http.Request) {
	newproblem := model.Problem{}
	err := json.NewDecoder(r.Body).Decode(&newproblem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding problem", err)
		return
	}
	err = h.ProblemRepo.CreateProblem(newproblem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while creating problem", err)
		return
	}
}

func (h *Handler) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	problem := model.Problem{}
	vars := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding problem", err)
		return
	}

	problem.Id = vars["id"]
	err = h.ProblemRepo.UpdateProblem(problem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while updating problem", err)
		return
	}
}

func (h *Handler) DeleteProblem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	err := h.ProblemRepo.DeleteProblem(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while deleting problem", err)
		return
	}
}
