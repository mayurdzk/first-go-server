package main

import (
	"encoding/json"
	"net/http"
)

// TODO: Why is there a requirement to export names?
// TODO: What is the `code` syntax?
type Person struct {
	// Notice custom JSON fields in response.
	Name string `json:"persons_name,omitempty"`
	Age  int    `json:"persons_age,omitempty"`
}

// TODO: Why are we passing in r? Can we do something interesting with r by reading the path, etc?
func getAllPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person
	people = append(people, Person{Name: "Jane Doe", Age: 20})
	people = append(people, Person{Name: "John Doe", Age: 21})

	json.NewEncoder(w).Encode(people)
}