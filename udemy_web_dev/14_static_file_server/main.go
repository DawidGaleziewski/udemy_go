package main

import "net/http"

func main() {
	http.ListenAndServe(":9090", http.FileServer(http.Dir("."))) // normally it would serve our whole directory. But this is a special case: if only index.html is at this location only it will be sevred by default
}