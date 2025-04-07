package templates

import (
	"os"
	"strings"
	"text/template"
)

func Example_pipesEx1() {
	tmpl := `
        {{.Text | printf "Text: %s"}}
        `

	t := template.Must(template.New("pipe").Parse(tmpl))

	data := map[string]string{
		"Text": "hello world",
	}

	t.Execute(os.Stdout, data)

	// Output:
	// Text: hello world
}

func Example_pipesEx2() {

	tmpl := `
	{{.Text | toupper | printf "Text: %s"}}
	`

	t := template.Must(template.New("pipe").Funcs(template.FuncMap{"toupper": strings.ToUpper}).Parse(tmpl))

	data := map[string]string{
		"Text": "hello world",
	}

	t.Execute(os.Stdout, data)

	// Output:
	// Text: HELLO WORLD

}
