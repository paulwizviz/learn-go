package templates

import (
	"os"
	"text/template"
)

func Example_conditionalEx1() {
	tmpl := `
{{if .IsAdmin}}Welcome, Admin {{.Name}}!{{else}}Welcome, User {{.Name}}!{{end}}`

	t := template.Must(template.New("user").Parse(tmpl))

	adminUser := map[string]any{"IsAdmin": true, "Name": "Alice"}
	regularUser := map[string]any{"IsAdmin": false, "Name": "Bob"}

	t.Execute(os.Stdout, adminUser)
	t.Execute(os.Stdout, regularUser)

	// Output:
	// Welcome, Admin Alice!
	// Welcome, User Bob!
}

func Example_conditionalEx2() {
	tmpl := `
{{if eq .Status "completed"}}Item is completed.{{else if eq .Status "pending"}}Item is pending.{{else}}Item status is unknown.{{end}}`

	t := template.Must(template.New("item").Parse(tmpl))

	completedItem := map[string]string{"Status": "completed"}
	pendingItem := map[string]string{"Status": "pending"}
	unknownItem := map[string]string{"Status": "unknown"}

	t.Execute(os.Stdout, completedItem)
	t.Execute(os.Stdout, pendingItem)
	t.Execute(os.Stdout, unknownItem)

	// Output:
	// Item is completed.
	// Item is pending.
	// Item status is unknown.

}

// Check empty slices
func Example_conditionalEx3() {
	tmpl := `
{{if .Items}} Items: {{range .Items}}{{.}} {{end}}{{else}}No items.{{end}}`

	t := template.Must(template.New("items").Parse(tmpl))

	withItems := map[string][]string{"Items": {"apple", "banana"}}
	noItems := map[string][]string{"Items": {}}

	t.Execute(os.Stdout, withItems)
	t.Execute(os.Stdout, noItems)

	// Output:
	// Items: apple banana
	// No items.
}
