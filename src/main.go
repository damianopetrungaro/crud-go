package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", ListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}