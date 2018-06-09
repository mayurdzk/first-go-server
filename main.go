package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// TODO: Why do we need MUX here? We can't we use the regular `http.HandleFunc()` method here, instead?
	router := mux.NewRouter()
	router.HandleFunc("/people", getAllPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}