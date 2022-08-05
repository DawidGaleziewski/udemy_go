package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template
func init() {
	var err error
	tpl, err = template.ParseFiles("templates/global.gohtml")
	if err != nil {
		log.Println("error creating template",err)
	}
}

func main() {
	data := struct{
		Words []string
		Label string
	}{
		Words: []string{"test1", "test2", "test3"},
		Label: "ABC1",
	}

	// we can access struct properties by defining the label after .Words
	// templates have build in functions like index, if, and, gt (greater then)
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Println(err)
	}
}