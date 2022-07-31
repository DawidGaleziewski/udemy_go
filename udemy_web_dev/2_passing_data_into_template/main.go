package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

// you can do this but you should not most of the time. As this is against separation of concerns. We have a type called FuncMap that is a map where each value is a interfcae{}
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init(){
	var err error
	// we can pass the functions to a template when we initialize it
	tpl, err = template.New("").Funcs(fm).ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Println(err)
	}
}

func main(){
	// pass data
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "test input")
	if err != nil {
		log.Fatalln(err)
	}

	// pass data and use it in template variable
	err = tpl.ExecuteTemplate(os.Stdout, "data_in_template.gohtml", "##hello from widsom variable inside template")
	if err != nil {
		log.Fatalln(err)
	}

	// passing composite data types
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err = tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}

	person := map[string]string {
		"name": "Ben",
		"surname": "Savege",
		"age": "18",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "map.gohtml", person)
	if err != nil {
		log.Fatalln(err)
	}

	// using functions inside the templates. We can use helpers like UC now and manipulate the templates inside. Do not over do this as logic should not be to heavy inside the templates
	err = tpl.ExecuteTemplate(os.Stdout, "tpl_funcs.gohtml", person)

}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}
