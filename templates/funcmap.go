package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type TemplateData struct {
	Name   string
	Result string
}

func splitListTrim(sep, orig string) []string {
	processedStr := strings.TrimRight(strings.TrimLeft(orig, sep), sep)
	return strings.Split(processedStr, sep)
}

func listLength(sep, orig string) int {
	return len(splitListTrim(sep, orig))
}

func notAllowedErr(errMessage string, isTrue bool) (err error) {
	if !isTrue {
		// no error if not empty
		return
	}
	fmt.Println(isTrue)

	if len(errMessage) == 0 {
		errMessage = "operation is not allowed"
	}

	err = fmt.Errorf("%v", errMessage)
	return
}

func main() {
	funcMap := template.FuncMap{
		"Length":        listLength,
		"SplitListTrim": splitListTrim,
		"NotAllowedErr": notAllowedErr,
	}

	// output of .Name is input to Length as a second argument
	tmpl, _ := template.New("myTemplate").Funcs(funcMap).Parse(string(`{{ .Name | Length " " | ne 4 | NotAllowedErr "not allowed err" | printf "namespace: %q" }}`))

	templateData := TemplateData{
		Name:   `"Hello" "there" "Hi"`,
		Result: "",
	}

	var result bytes.Buffer

	err := tmpl.Execute(&result, templateData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", result.String())
}
