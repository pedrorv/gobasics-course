package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl-slice.gohtml", "tpl-slice2.gohtml"))
}

func main() {
	office := []string{"Dwight", "Jim", "Michael", "Pam"}

	err := tpl.Execute(os.Stdout, office)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-slice2.gohtml", office)
	if err != nil {
		log.Fatalln(err)
	}
}