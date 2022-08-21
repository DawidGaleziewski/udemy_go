package main

import (
	"fmt"
	"net/http"
)

func main() {
	// there are couple of redirect codes and they will differ in significant ways
	// 301 - moved permanently. If we give this code, many browsers will remember this and NEVER hit that servers url again. It will go straignt to the redirected url ALWAYS
	// 303 - see other. Will change method to GET
	// 307 temporary redirects - keeps same method
	// ultimate source for definitive information on anything http protocol related https://www.rfc-editor.org/rfc/rfc7231

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "redirected")
	})

	http.HandleFunc("/move-perm", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})

	http.HandleFunc("/see-other", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w,r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/temp-redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w,r, "/", http.StatusTemporaryRedirect)
	})


	// second way of doing this is by setting up headers on response request
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/") // set new location for redirect
		w.WriteHeader(301) // will set it as permanent moved asset. Browser will cache it and not event try to get the resource next time
	})

	http.ListenAndServe(":9090", nil)
}
