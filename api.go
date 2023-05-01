package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

// Helpers functions to handle errors in one place

type apiFunc func(http.ResponseWriter, *http.Request) error
type ApiError struct {
	Error string
}
type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {

	return &APIServer{

		listenAddr: listenAddr,
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("content-type", "Application/json")
	return json.NewEncoder(w).Encode(v)

}

// convert incoming request and
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)

		if err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{
				Error: err.Error()})
		}
	}

}

// START AND RUN THE SERVER
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Println("JSON API SERVER RUNNING ON PORT:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// Initial Entry Function for accounts
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)

	}
	if r.Method == "POST" {

		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)

	}
	return fmt.Errorf("Method Not Aloowed %s", r.Method)
}
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	// account := NewAccount("Pjd", "Swanepoel")
	return WriteJson(w, http.StatusOK, vars)
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}
