package main

// TODO: Separate into its own package

import (
	"html/template"
)

type FormValidationResult struct {
	Success        bool
	IsAgeIncorrect bool
}

func peopleHTMLTemplate() *template.Template {
	var tmpl = `<html>
<head>
    <title>Hello!</title>
</head>
<body>
	<ol>
	{{range .}}
	<li>{{.Name}}, Aged: {{.Age}}</li>
	{{end}}
	</ol>
</body>
</html>
`
	return createTemplate(tmpl)
}

func newPersonHTMLTemplate() *template.Template {
	// TODO: Break this out. Use different templates for post and get.
	var tmpl = `
	{{if .Success}}
	<h1>Thanks!</h1>
	{{else}}
	<h1>New Person{{if .IsAgeIncorrect}}😿{{end}}</h1>
	{{if .IsAgeIncorrect}}<h3>Please enter a valid age.</h3>{{end}}
	<form method="POST">
		<label>Name:</label><br />
		<input type="text" name="name"><br />
		<label>Age:</label><br />
		<input type="text" name="age"><br />
		<input type="submit">
	</form>
	{{end}}
	`
	return createTemplate(tmpl)
}

func createTemplate(templateString string) *template.Template {
	// TODO: Why is a name required?
	t := template.New("main")
	t, err := t.Parse(templateString)
	if err != nil {
		panic(err)
	}
	return t
}
