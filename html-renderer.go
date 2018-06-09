package main
// TODO: Separate into its own package

import (
	"fmt"
	"html/template"
)

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
	t := template.New("main")
	t, err := t.Parse(tmpl)
	fmt.Println(t)
	if err != nil {
		panic(err)
	}
	return t
}
