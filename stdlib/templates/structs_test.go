package templates

import (
	"fmt"
	"os"
	"text/template"
)

type Address struct {
	City     string
	Country  string
	PostCode string
}

type User struct {
	Name    string
	Age     int
	Address Address
	Hobbies []string
}

func Example_structEx1() {
	tmpl := `
User Profile
Name: {{.Name}}
Age: {{.Age}}
Address:
City: {{.Address.City}}
Country: {{.Address.Country}}
Post Code: {{.Address.PostCode}}
Hobbies: {{range .Hobbies}}<li>{{.}}</li>{{end}}`

	t := template.Must(template.New("user").Parse(tmpl))

	user := User{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:     "London",
			Country:  "UK",
			PostCode: "SW1A 0AA",
		},
		Hobbies: []string{"reading", "hiking", "coding"},
	}

	err := t.Execute(os.Stdout, user)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// User Profile
	// Name: Alice
	// Age: 30
	// Address:
	// City: London
	// Country: UK
	// Post Code: SW1A 0AA
	// Hobbies: <li>reading</li><li>hiking</li><li>coding</li>
}

func Example_structEx2() {
	tmpl := `
User Profile
Name: {{.Name}}
Age: {{.Age}}
Address:{{with .Address}}
City: {{.City}}
Country: {{.Country}}
Post Code: {{.PostCode}}{{end}}
Hobbies: {{range .Hobbies}} {{.}}{{end}}`

	t := template.Must(template.New("user").Parse(tmpl))

	user := User{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:     "London",
			Country:  "UK",
			PostCode: "SW1A 0AA",
		},
		Hobbies: []string{"reading", "hiking", "coding"},
	}

	err := t.Execute(os.Stdout, user)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// User Profile
	// Name: Alice
	// Age: 30
	// Address:
	// City: London
	// Country: UK
	// Post Code: SW1A 0AA
	// Hobbies:  reading hiking coding
}
