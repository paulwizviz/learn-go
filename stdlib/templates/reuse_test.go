package templates

import (
	"fmt"
	"os"
	"text/template"
)

// This example use online
// template definition
func Example_reuseEx1() {
	headerTemplate := `
	{{define "header"}}
Header: {{.Header}}
{{end}}`

	bodyTemplate := `
	{{template "header" .}}
FirstName: {{.FirstName}}
Surname: {{.Surname}}
	`

	combinedTemplate := headerTemplate + bodyTemplate

	t := template.Must(template.New("userprofile").Parse(combinedTemplate))

	data := map[string]string{
		"Header":    "User Profile",
		"FirstName": "John",
		"Surname":   "Doe",
	}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Header: User Profile
	//
	// FirstName: John
	// Surname: Doe
}

func Example_resuseEx2() {
	data := map[string]string{
		"Title":     "User Profile",
		"FirstName": "John",
		"Surname":   "Doe",
	}
	t, err := template.ParseFiles("./ex2/header.tmpl", "./ex2/body.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "body.tmpl", data)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Title: User Profile
	// FirstName: John
	// Surname: Doe
}

func Example_reuseEx3() {
	tmpl := template.New("userprofile")

	tmpl, err := tmpl.Parse(`
	{{define "header"}}
Title: {{.Title}}{{end}}`)
	if err != nil {
		fmt.Println(err)
	}

	tmpl, err = tmpl.Parse(`
	{{template "header" .}}
FirstName: {{.FirstName}}
Surname: {{.Surname}}
	`)

	data := map[string]string{
		"Title":     "User Profile",
		"FirstName": "John",
		"Surname":   "Doe",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Title: User Profile
	// FirstName: John
	// Surname: Doe
}
