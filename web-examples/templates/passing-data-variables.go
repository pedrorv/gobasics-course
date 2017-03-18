package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl-variable.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl-variable.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}