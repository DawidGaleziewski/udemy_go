package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
func init() {
	var err error
	tpl, err = template.ParseGlob("templates/*.gohtml")

	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", "")
		
		if err != nil {
			log.Panicln(err)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "about.gohtml", "")

		if err != nil {
			log.Panicln(err)
		}
	})

	http.ListenAndServe(":9090", nil)
}