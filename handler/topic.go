package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetTopics(w http.ResponseWriter, r *http.Request) {
	filter := model.TopicFilter{}
	query := r.URL.Query()
	if query.Has("name"){
		name := query.Get("name")
		filter.Name = &name
	}

	topics, err := h.TopicRepo.GetTopics(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting topics",err)
		return
	} 

	err = json.NewEncoder(w).Encode(topics)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding topics", err)
		return
	}

}

func (h *Handler) GetTopicByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	topic, err := h.TopicRepo.GetTopicById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting topic by Id", err)
		return
	}
	err = json.NewEncoder(w).Encode(topic)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding topic", err)
		return
	}
}

func (h *Handler) CreateTopic(w http.ResponseWriter, r *http.Request) {
	newtopic := model.Topic{}
	err := json.NewDecoder(r.Body).Decode(&newtopic)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding topic", err)
		return
	}
	err = h.TopicRepo.CreateTopic(newtopic)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while creating topic", err)
		return
	}
}

func (h *Handler) UpdateTopic(w http.ResponseWriter, r *http.Request) {
	topic := model.Topic{}
	vars := mux.Vars(r)
	
	err := json.NewDecoder(r.Body).Decode(&topic)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding topic", err)
		return
	}

	topic.Id = vars["id"]
	err = h.TopicRepo.UpdateTopic(topic)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while updating topic", err)
		return
	}
}

func (h *Handler) DeleteTopic(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	err := h.TopicRepo.DeleteTopic(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while deleting topic", err)
		return
	}
}
