package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template
func init() {
	var err error
	tpl, err = template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Println("Error parsing templates: ", err)
	}
}

func main() {
	// By using define keyword we can declate a template.
	// By using template keyword we can call our defined template
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "test data passed from main.go")
	if err != nil {
		log.Println(err)
	}
}