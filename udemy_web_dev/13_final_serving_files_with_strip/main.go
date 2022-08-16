package main

import "net/http"

func main() {
	// we use handle. Which accepts a route and a type implementing a handler interface
	// we serve whole folder /assets at route /static
	// we strip the static (url name) from the url 
	http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":9090", nil)
}