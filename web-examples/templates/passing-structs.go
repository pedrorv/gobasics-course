package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type employee struct {
	Name string
	Job  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl-struct.gohtml"))
}

func main() {
	dwight := employee{
		Name: "Dwight",
		Job: "Salesman",
	}

	err := tpl.Execute(os.Stdout, dwight)
	if err != nil {
		log.Fatalln(err)
	}
}