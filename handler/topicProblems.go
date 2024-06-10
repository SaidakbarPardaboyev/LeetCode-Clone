package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Read
func (h *Handler) GetTopicProblems(w http.ResponseWriter, r *http.Request) {
	filter := model.TopicProblemFilter{}
	query := r.URL.Query()
	if query.Has("topic_id"){
		TopicId := query.Get("topic_id")
		filter.TopicId = &TopicId
	}
	if query.Has("problem_id"){
		ProblemId := query.Get("problem_id")
		filter.ProblemId = &ProblemId
	}

	TopicProblems, err := h.TopicProblemRepo.GetTopicProblems(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting TopicProblems ",err)
		return
	} 

	err = json.NewEncoder(w).Encode(TopicProblems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding TopicProblems ", err)
		return
	}

}

func (h *Handler) GetTopicProblemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	TopicProblem, err := h.TopicProblemRepo.GetTopicProblemById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting TopicProblem by Id ", err)
		return
	}
	err = json.NewEncoder(w).Encode(TopicProblem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding TopicProblem ", err)
		return
	}
}

func (h *Handler) GetProblemsByTopicId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	topicId := vars["topic_id"]

	problems, err := h.TopicProblemRepo.GetTopicProblemById(topicId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting problems by topic_id ", err)
		return
	}
	err = json.NewEncoder(w).Encode(problems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding problems ", err)
		return
	}
}
func (h *Handler) GetTopicsByProblemId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problemId := vars["problem_id"]

	problems, err := h.TopicProblemRepo.GetTopicsByProblemId(problemId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting topics by problem_id ", err)
		return
	}
	err = json.NewEncoder(w).Encode(problems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding topics ", err)
		return
	}
}

// create
func (h *Handler) CreateTopicProblem(w http.ResponseWriter, r *http.Request) {
	newTopicProblem := model.TopicProblem{}
	err := json.NewDecoder(r.Body).Decode(&newTopicProblem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding TopicProblem", err)
		return
	}
	err = h.TopicProblemRepo.CreateTopicProblem(newTopicProblem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while creating TopicProblem", err)
		return
	}
}

// Update
func (h *Handler) UpdateTopicProblem(w http.ResponseWriter, r *http.Request) {
	TopicProblem := model.TopicProblem{}
	vars := mux.Vars(r)
	
	err := json.NewDecoder(r.Body).Decode(&TopicProblem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding TopicProblem", err)
		return
	}

	TopicProblem.Id = vars["id"]
	err = h.TopicProblemRepo.UpdateTopicProblem(TopicProblem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while updating TopicProblem", err)
		return
	}
}

// Delete
func (h *Handler) DeleteTopicProblem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	err := h.TopicProblemRepo.DeleteTopicProblem(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while deleting TopicProblem", err)
		return
	}
}
