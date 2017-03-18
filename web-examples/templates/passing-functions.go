package main

import (
	"log"
	"os"
	"text/template"
	"strings"
)

var tpl *template.Template

type employee struct {
	Name string
	Job  string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl-functions.gohtml"))
}

func main() {
	dwight := employee{
		Name: "Dwight",
		Job: "Salesman",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-functions.gohtml", dwight)
	if err != nil {
		log.Fatalln(err)
	}
}