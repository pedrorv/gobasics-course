package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl-map.gohtml", "tpl-map2.gohtml"))
}

func main() {
	office := map[string]string{
		"Dwight": "Salesman", 
		"Jim": "Salesman",
		"Michael": "Regional Manager",
		"Pam": "Receptionist",
	}

	err := tpl.Execute(os.Stdout, office)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-map2.gohtml", office)
	if err != nil {
		log.Fatalln(err)
	}
}