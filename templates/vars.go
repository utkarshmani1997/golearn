package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Masterminds/sprig"
)

func main() {
	t := template.New("test").Funcs(sprig.FuncMap())
	t, _ = t.Parse(`{{- $ns := `)
	err := t.Execute(os.Stdout, p)
	fmt.Println(err)
}
