package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"math/rand"
  	"time"
)

// TODO: Why is there a requirement to export names?
// TODO: What is the `code` syntax?
type Person struct {
	// Notice custom JSON fields in response.
	Name string `json:"persons_name,omitempty"`
	Age  int8    `json:"persons_age,omitempty"`
}

// Creates an array of the Person object that's used throughout the app.
// The array is populated with default values. It's ready to be shown to a user.
func BuildPeopleArray() []Person {
	people := make([]Person, 2)
	people[0] = Person{Name: "Jane Doe", Age: 20}
	people[1] = Person{Name: "John Doe", Age: 24}

	return people
}

// TODO: Why are we passing in r? Can we do something interesting with r by reading the path, etc?
func getAllPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person
	people = append(people, Person{Name: "Jane Doe", Age: 20})
	people = append(people, Person{Name: "John Doe", Age: 21})

	json.NewEncoder(w).Encode(people)
}

func NewPerson() Person {
	// TODO: Take these out of the function. Why initialise them each time the function is called?
	firstNames := [4]string{"Jack", "Joe", "Lisa", "Margaret"}
	lastNames := [4]string{"Smith", "Seinfeld", "Gruber", "Perry"}
	ages := [...]int8{26, 23, 30, 21, 19, 34}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator 
	
	firstName := firstNames[r.Intn(len(firstNames))]
	lastName := lastNames[r.Intn(len(lastNames))]
	age := ages[r.Intn(len(ages))]

	name := fmt.Sprintf("%s %s", firstName, lastName)

	return Person{Name: name, Age: age}
}