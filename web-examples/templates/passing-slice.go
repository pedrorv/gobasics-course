package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl-slice.gohtml"))
}

func main() {
	office := []string{"Dwight", "Jim", "Michael", "Pam"}

	err := tpl.Execute(os.Stdout, office)
	if err != nil {
		log.Fatalln(err)
	}
}