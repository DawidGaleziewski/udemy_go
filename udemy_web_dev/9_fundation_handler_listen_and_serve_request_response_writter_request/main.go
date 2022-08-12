package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template
func init(){
	var err error
	tpl, err = tpl.ParseGlob("templates/*gohtml")
	if err != nil {
		log.Panicln(err)
	}
}

type app struct{}

func (a app) ServeHTTP(resWritter http.ResponseWriter,req *http.Request){
	req.ParseForm() // we need to parse the form before using it

	// here we can see some useful data on the req type
	data := struct{
		Form map[string][]string; 
		Method string;
		URL *url.URL;
		Header http.Header;
		Host string;
		ContentLength int64;
	}{ 
		Form: req.Form, // Form allows us to get data from request that contains all inputs. // there is also field called PostForm for post body only. Form will also contain data from the url
		Method: req.Method,
		URL: req.URL,
		Header: req.Header,
		Host: req.Host,
		ContentLength: req.ContentLength,
	}

	err := tpl.ExecuteTemplate(resWritter, "index.gohtml", data)  // we can use response writter to display our tempplate
	if err != nil {
		log.Println(err)
	}
	
}

func main() {
	var myApp app
	http.ListenAndServe(":9090", myApp)
}