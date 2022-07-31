package main

import (
	"log"
	"os"
	"text/template"
)

// one thing we can do to optimize our templates is to make sure they are parsed only once
var initTpl *template.Template
func init(){
	// Musat will do erro checking.
	initTpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

// go standard library has two kinds of templates. Text templates which is the fundation and HTML templates which builds on top of it but has more thigs like security baked into it.
// go encourages to use less abstractions and think more like a programer

// we will be parsing a file. By reading a html file and using it to bring it into our projct
// we will use tpl.gohtml file. Extension .gohtml is not anything like html tmpl files. This is just text, we use it to hold a string
func main(){
	// tpl is a container to all the templates that were parsed
	tpl, err := template.ParseFiles("templates/tpl.gohtml")
	if err != nil {
		log.Println(err)
	}

	nf, err := os.Create("templates/index.html")
	if err != nil {
		log.Println("error creating a file", err)
	}
	defer nf.Close()

	// we use method execute. Execute passes data to writer interface. We can use os.Stdout to just print the template
	//err = tpl.Execute(os.Stdout, nil)

	// we can also create a new file and pass it in. As it implements writter interface we can do this. In this case the data will be passed to the file we created by its implementation of the writter interface, new file will be created
	err = tpl.Execute(nf, nil)

	if err != nil {
		log.Fatalln(err)
	}

	// we can also add more files to the template container. One way of getting all files is using parseGlob
	tplGroup, err := tpl.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Println(err)
	}

	// we can select a single template from the struct
	err = tplGroup.ExecuteTemplate(os.Stdout, "about_us_tmpl.gohtml", nil)
	if err != nil {
		log.Panicln(err)
	}

}