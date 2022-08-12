package main

import (
	"fmt"
	"net/http"
)

type hotdogKind string
func(h hotdogKind) ServeHTTP(writer http.ResponseWriter, r *http.Request){
	fmt.Fprintln(writer, "any code handling this func. This will be written to the connection i.e from a browser")
}

type toolbox struct{}
func (t toolbox) ServeHTTP(writer http.ResponseWriter, r *http.Request){
	fmt.Fprintln(writer, "any code here")
}

func main() {

	// as long as a type is also of type interface handler we can pass it to listen and serve interface
	var hd hotdogKind
	var t toolbox

	http.ListenAndServe(":9090", hd)
	http.ListenAndServe(":9091", t)
}