package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// we can always pass values

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})
	http.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		// make sure method is post
		if r.Method == http.MethodPost {
			file, fileMetadata, err := r.FormFile("dick-pics")
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

			//creating new file
			dst, err := os.Create(filepath.Join(filepath.Join("./user/", fileMetadata.Filename)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dst.Close()

			// write to a file we have created
			_, err = dst.Write(bs)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
	})
	http.ListenAndServe(":9090", nil)
}