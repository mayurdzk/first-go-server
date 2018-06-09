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

	router.HandleFunc("/add-person", func(w http.ResponseWriter, r *http.Request) {
		newPerson := NewPerson()
		people = append(people, newPerson)

		http.Redirect(w, r, "people", http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}