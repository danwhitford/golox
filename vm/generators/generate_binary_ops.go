package main

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed binary_ops.go.tmpl
var tmplString string

func main() {
	t := template.Must(template.New("binaryops").Parse(tmplString))

	f, err := os.Create("binaryops.go")
	if err != nil {
		panic(err)
	}

	data := []struct {
		OpName string
		Op     string
	}{
		{
			"Add",
			"+",
		},
		{
			"Sub",
			"-",
		},
		{
			"Mult",
			"*",
		},
		{
			"Div",
			"/",
		},
	}

	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
}
