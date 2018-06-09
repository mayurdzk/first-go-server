package main

import (
	"log"
	"net/http"
)

func main() {
	people := BuildPeopleArray()
	router := firstGoServerRouter(people)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func firstGoServerRouter(people []Person) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		getAllPeople(w, people)
	})

	router.HandleFunc("/add-person", func(w http.ResponseWriter, r *http.Request) {
		people = addPerson(w, r, people)
	})

	return router
}