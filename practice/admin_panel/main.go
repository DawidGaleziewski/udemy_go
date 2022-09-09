package main

import (
	"html/template"
	"net/http"
)

var templates template.Template
var tmplError error

func init() {
	templates, tmpError := template.ParseGlob("templates/*.go")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

		}
	})
}
