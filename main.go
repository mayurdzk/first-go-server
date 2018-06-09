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
			t.Execute(w, nil)
			return
		}

		ageString := r.FormValue("age")
		ageNum, err := strconv.ParseInt(ageString, 10, 8)
		if err != nil {
			// TODO: Handle this better. Age needs to be a number between 0-150
			panic(err)
		}
		age := int8(ageNum)

		newPerson := Person{
			Name: r.FormValue("name"),
			Age:  age,
		}
		people = append(people, newPerson)

		t.Execute(w, struct{ Success bool }{true})
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}
