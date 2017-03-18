package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl-nesting.gohtml", "tpl-inner.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl-nesting.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}