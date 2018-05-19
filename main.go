package main

import (
	"encoding/json"
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
