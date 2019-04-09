package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Masterminds/sprig"
)

type Person struct {
	Id    int
	Count int
	Name  string
}

func main() {
	t := template.New("test").Funcs(sprig.FuncMap())
	t, _ = t.Parse(`{{range $j := until .Count}}
{{ . }}
{{end}}
{{ .Name }}
{{ .Id }}
`)
	p := Person{
		Name:  "Utkarsh",
		Id:    1234,
		Count: 10,
	}
	err := t.Execute(os.Stdout, p)
	fmt.Println(err)
}
