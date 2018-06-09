package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	people := BuildPeopleArray()

	router := http.NewServeMux()
	router.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		// TODO, Use `getAllPeople()` instead.
		json.NewEncoder(w).Encode(people)
	})
	log.Fatal(http.ListenAndServe(":8000", router))
}