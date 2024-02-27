package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) error {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTasks).Methods("GET")

	return nil

}

func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {

}

func (s *TasksService) handleGetTasks(w http.ResponseWriter, r *http.Request) {

}
