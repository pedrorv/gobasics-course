package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl-data.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl-data.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}