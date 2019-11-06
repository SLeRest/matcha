package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user", getAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getOneUser).Methods("GET")
	router.HandleFunc("/user/{id}", updateUser).Methods("PATCH")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
