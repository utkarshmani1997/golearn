package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type TemplateData struct {
	Name string
}

func splitListTrim(sep, orig string) []string {
	processedStr := strings.TrimRight(strings.TrimLeft(orig, sep), sep)
	return strings.Split(processedStr, sep)
}

func listLength(sep, orig string) int {
	return len(splitListTrim(sep, orig))
}

func main() {
	funcMap := template.FuncMap{
		"Length":        listLength,
		"SplitListTrim": splitListTrim,
	}

	// output of .Name is input to Length as a second argument
	tmpl, _ := template.New("myTemplate").Funcs(funcMap).Parse(string(`{{ .Name | Length " " | ge 3  }}`))

	templateData := TemplateData{
		Name: `"Hello" "there" "Hi"`,
	}

	var result bytes.Buffer

	tmpl.Execute(&result, templateData)
	fmt.Printf("%q", result)
}
