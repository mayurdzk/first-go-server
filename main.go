package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	people := BuildPeopleArray()

	router := http.NewServeMux()
	router.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		// TODO, Use `getAllPeople()` instead.
		t := peopleHTMLTemplate()
		t.Execute(w, people)
	})

	router.HandleFunc("/add-person", func(w http.ResponseWriter, r *http.Request) {
		t := newPersonHTMLTemplate()

		if r.Method != http.MethodPost {
			// Serve the user the form
			formSubmitionNeededResult := FormValidationResult{Success: false, IsAgeIncorrect: false}
			t.Execute(w, formSubmitionNeededResult)
			return
		}

		ageString := r.FormValue("age")
		ageNum, err := strconv.ParseInt(ageString, 10, 8)
		if err != nil || ageNum < 1 || ageNum > 150 {
			invalidAgeFormResult := FormValidationResult{Success: false, IsAgeIncorrect: true}
			t.Execute(w, invalidAgeFormResult)
			return
		}
		age := int8(ageNum)

		newPerson := Person{
			Name: r.FormValue("name"),
			Age:  age,
		}
		people = append(people, newPerson)

		validFormResult := FormValidationResult{Success: true, IsAgeIncorrect: false}
		t.Execute(w, validFormResult)
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}