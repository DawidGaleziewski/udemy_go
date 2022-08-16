package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// we can always pass values

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})
	http.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		var s string
		// make sure method is post
		if r.Method == http.MethodPost {
			file, _, err := r.FormFile("dick-pics")
			if err != nil {
				http.Error(w, err.Error() + " on form parse", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// read file
			bs, err := ioutil.ReadAll(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			s = string(bs)

			fmt.Println(s)
		}
	})
	http.ListenAndServe(":9090", nil)
}