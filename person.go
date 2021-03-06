package main

import (
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

func getAllPeople(w http.ResponseWriter, db *DB) {
	t := peopleHTMLTemplate()
	var people []Person
	rows, err := db.Query("SELECT * FROM People")
        assertNotError(err)

        for rows.Next() {
            var name string
			var age int8

            err = rows.Scan(&name, &age)
            assertNotError(err)

			person := Person{Name: name, Age: age}
			people = append(people, person)
        }
	t.Execute(w, people)
}

func addPerson(w http.ResponseWriter, r *http.Request, db *DB) {
	t := newPersonHTMLTemplate()

	if r.Method != http.MethodPost {
		// Serve the user the form
		formSubmitionNeededResult := FormValidationResult{IsAgeIncorrect: false}
		t.Execute(w, formSubmitionNeededResult)
		return
	}

	ageString := r.FormValue("age")
	ageNum, err := strconv.ParseInt(ageString, 10, 8)
	if err != nil || ageNum < 1 || ageNum > 150 {
		invalidAgeFormResult := FormValidationResult{IsAgeIncorrect: true}
		t.Execute(w, invalidAgeFormResult)
		return
	}
	age := int8(ageNum)
	name := r.FormValue("name")

	stmt, err := db.Prepare("insert into People values(?, ?)")
	assertNotError(err)

	_, e := stmt.Exec(name, age)
	assertNotError(e)

	http.Redirect(w, r, "people", http.StatusFound)
}
