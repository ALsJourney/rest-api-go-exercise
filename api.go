package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewApiServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	subrouter := s.createSubRouter()

	// registering our services
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting the API Server at ", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}

func (*APIServer) createSubRouter() *mux.Router {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	return subrouter
}
