package main

import (
	"os"
	"text/template"
)

//go:generate go run main.go

var tpl = template.Must(template.New("").Parse(`Hello World!`))

func main() {
	data := struct{}{}
	out, _ := os.Create("hello.txt")
	defer out.Close()
	tpl.Execute(out, data)
}
