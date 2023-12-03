package intline

import (
	"bytes"
	"fmt"
	"text/template"
)

type UserProfile struct {
	FirstName string
	Lastname  string
}

func (s UserProfile) DisplayName(firstName string, lastname string) string {
	return fmt.Sprintf("%v %v", firstName, lastname)
}

func Example_createUserProfile() {

	tmplVars := `FirstName: {{.FirstName}}
Lastname: {{.Lastname}}`

	tmpl := template.Must(template.New("userprofile").Parse(tmplVars))

	data := struct {
		FirstName string
		Lastname  string
	}{
		FirstName: "John",
		Lastname:  "Doe",
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	fmt.Println(buf.String())
	// Output:
	// FirstName: John
	// Lastname: Doe

}

func Example_createUserProfileWithDisplayName() {

	tmplVars := `FirstName: {{.FirstName}}
Lastname: {{.Lastname}}
Displayname: {{displayName .FirstName .Lastname}}`

	var tmpl = template.Must(template.New("userprofile").Funcs(template.FuncMap{
		"displayName": func(firstName string, lastname string) string { return fmt.Sprintf("%v %v", firstName, lastname) }}).Parse(tmplVars))

	data := struct {
		FirstName string
		Lastname  string
	}{
		FirstName: "John",
		Lastname:  "Doe",
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	fmt.Println(buf.String())
	// Output:
	// FirstName: John
	// Lastname: Doe
	// Displayname: John Doe

}

func Example_createUserProfileMemberMethod() {

	tmplVars := `FirstName: {{.FirstName}}
Lastname: {{.Lastname}}
DisplayName: {{.DisplayName .FirstName .Lastname}}`

	tmpl, err := template.New("userprofile").Parse(tmplVars)
	if err != nil {
		fmt.Println("Not expected")
	}

	data := UserProfile{
		FirstName: "John",
		Lastname:  "Doe",
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	fmt.Println(buf.String())
	// Output:
	// FirstName: John
	// Lastname: Doe
	// DisplayName: John Doe

}

func Example_templateRanging() {

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
