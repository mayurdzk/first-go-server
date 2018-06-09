package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/people", getAllPeople)
	log.Fatal(http.ListenAndServe(":8000", router))
}