package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// TODO: Why is there a requirement to export names?
// TODO: What is the `code` syntax?
type Person struct {
	// Notice custom JSON fields in response.
	Name string `json:"persons_name,omitempty"`
	Age  int8   `json:"persons_age,omitempty"`
}

// Creates an array of the Person object that's used throughout the app.
// The array is populated with default values. It's ready to be shown to a user.
func BuildPeopleArray() []Person {
	people := make([]Person, 2)
	people[0] = Person{Name: "Jane Doe", Age: 20}
	people[1] = Person{Name: "John Doe", Age: 24}

	return people
}

func getAllPeople(w http.ResponseWriter, people []Person) {
	t := peopleHTMLTemplate()
	t.Execute(w, people)
}

func addPerson(w http.ResponseWriter, r *http.Request, people []Person) []Person {
	t := newPersonHTMLTemplate()

	if r.Method != http.MethodPost {
		// Serve the user the form
		formSubmitionNeededResult := FormValidationResult{Success: false, IsAgeIncorrect: false}
		t.Execute(w, formSubmitionNeededResult)
		return people
	}

	ageString := r.FormValue("age")
	ageNum, err := strconv.ParseInt(ageString, 10, 8)
	if err != nil || ageNum < 1 || ageNum > 150 {
		invalidAgeFormResult := FormValidationResult{Success: false, IsAgeIncorrect: true}
		t.Execute(w, invalidAgeFormResult)
		return people
	}
	age := int8(ageNum)

	newPerson := Person{
		Name: r.FormValue("name"),
		Age:  age,
	}
	fmt.Println(people)
	people = append(people, newPerson)
	fmt.Println(people)

	validFormResult := FormValidationResult{Success: true, IsAgeIncorrect: false}
	t.Execute(w, validFormResult)
	return people
}