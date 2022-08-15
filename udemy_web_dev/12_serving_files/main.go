package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./assets"))) // simplest way to serve each file in directory

	// other ways of serving files. Info only as this is not used in the wild
	// http.HandleFunc("/hill.jpg", hillPic) // serving a picture on this route
	// http.HandleFunc("/hill-4.jpg", secondHillPic)
	// http.HandleFunc("hill-5.jpg", thirdPic)


	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 	io.WriteString(w, `
	// 		<img src="hill.jpg">
	// 	`)
	// })

	http.ListenAndServe(":9090", nil)
}

// info only!
func hillPic(w http.ResponseWriter, req *http.Request){
	file , err := os.Open("/hill.jpg") // we open a file which returns a pointer. This pointer implements the writer interface
	if err != nil {
		http.Error(w, "file not found", 404) // way of handling error responses
	}

	defer file.Close()
	io.Copy(w, file) // we can pas our picture file here as it implements the writer interface. Therefore the file is copied and send via the responseWriter method
}

// info only!
func secondHillPic(w http.ResponseWriter, req *http.Request){
	// serve content will also provide information like content name, file size etc.
	file, err := os.Open("hill.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	
	// serve content will return more info regarding the metadata. 
	// this method is rarly used and its just for information purposes
	http.ServeContent(w, req, file.Name(), fileStat.ModTime(), file)
}

// info only!
func thirdPic(w http.ResponseWriter, req *http.Request){
	// rarly used as well
	http.ServeFile(w, req, "hill.jpg")
}