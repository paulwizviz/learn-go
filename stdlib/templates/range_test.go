package templates

import (
	"bytes"
	"fmt"
	"text/template"
)

func Example_range() {
	tmplVars := `{{range .Profiles}}
------------------
FirstName: {{.FirstName}}
Lastname: {{.Lastname}}
{{end}}`

	tmpl, err := template.New("listprofiles").Parse(tmplVars)
	if err != nil {
		fmt.Println("Not expected")
	}

	profiles := []UserProfile{
		{"Alice", "Doe"},
		{"Bob", "Doe"},
		{"Charlie", "Doe"},
	}

	listOfProfiles := struct {
		Profiles []UserProfile
	}{
		Profiles: profiles,
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, listOfProfiles)
	fmt.Println(buf.String())

	// Output:
	// ------------------
	// FirstName: Alice
	// Lastname: Doe
	//
	// ------------------
	// FirstName: Bob
	// Lastname: Doe
	//
	// ------------------
	// FirstName: Charlie
	// Lastname: Doe
}
