package templates

import (
	"fmt"
	"os"
	"text/template"
)

func renderTemplate(tmpl string, data map[string]string) {
	t, err := template.ParseFiles("ex3/base.tmpl", "ex3/home.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	err = t.ExecuteTemplate(os.Stdout, tmpl, data)
	if err != nil {
		fmt.Println(err)
	}
}

func handle1() {
	data := map[string]string{"Head": "This is a home page", "Content": "This is a content"}
	renderTemplate("home", data)
}

func Example_blockEx1() {
	handle1()

	// Output:
	// Title Head: This is a home page
	// Content Content: This is a content
}
