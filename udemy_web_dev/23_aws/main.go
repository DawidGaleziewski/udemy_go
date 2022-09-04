package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("starting app")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("recived a http connection")
		fmt.Fprint(w, "howdy there")
	})

	http.ListenAndServe(":80", nil)
}