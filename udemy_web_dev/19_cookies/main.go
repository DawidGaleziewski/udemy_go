package main

import (
	"fmt"
	"net/http"
)

// cookie is a file, that server can use to write to client browser, if its allowed to by the machine
// cookie will be send to the server on each request by the server
// cookies are domain specific
// when using https, all data, including that passed by the cookie and the https params, is encrypted
// when passing the data by url, if someone was looking by your shoulder they could hajiack the session by getting that id from the url
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// to set a cookie we just need to use SetCookie method
		// we pass in a pointer to a cookie struct
		// cookie will be available only in this domain
		http.SetCookie(w, &http.Cookie{
			Name: "tracking-go-id",
			Value: "user-no-123",
		})

		// cookie will be present on our response and will be set to this value
		fmt.Fprintln(w, "Cookie has been set")
	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		// to get any cookie will simply use Cookie method on the request
		cookie, err := r.Cookie("tracking-go-id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, "YOUR COOKIE: ", cookie)
	})

	http.ListenAndServe(":9090", nil)
}