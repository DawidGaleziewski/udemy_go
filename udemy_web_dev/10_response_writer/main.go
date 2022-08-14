package main

import (
	"fmt"
	"net/http"
)

type app struct{}

func (a app) ServeHTTP(responseWriter http.ResponseWriter, req *http.Request){
	responseWriter.Header().Set("random-key", "random value here") // set headers before sending response back
	responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(responseWriter, "<h1>This is some random response</h1>")
}

func main(){
	var myApp app
	http.ListenAndServe(":9090", myApp)
}