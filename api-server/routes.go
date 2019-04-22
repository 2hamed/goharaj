package main

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login", LoginHandler).Methods("POST")

	return r
}
