package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := os.Open("test.html")
		if err != nil {
			// http error will write a response with error.
			// We can also use constants on the http package to provide status codes. Or just pass 404 etc.
			http.Error(w, "error msg to the user here", http.StatusNotFound)
			return // important thing is to remember that http.Error will not break out the function, we still need to end it with a guard clause here
		}
	})

	log.Fatal(http.ListenAndServe(":9090", nil)) // log fata for cases where it stops working. As ListenAndServe returns a error we can pass it right to log, in case of a error it will be logged to a file.
}