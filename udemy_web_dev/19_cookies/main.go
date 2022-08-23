package main

import (
	"fmt"
	"net/http"
	"strconv"
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

		// we can set multiple cookies this way
		http.SetCookie(w, &http.Cookie{
			Name: "spy-1",
			Value: "spy on this guy",
		})
		// cookie will be present on our response and will be set to this value
		fmt.Fprintln(w, "Cookie has been set")
	})

	// we want to handle favicon as some browsers will request for it on start. We have a build in method for handling not found
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		// to get any cookie will simply use Cookie method on the request
		cookie, err := r.Cookie("tracking-go-id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, "YOUR COOKIE: ", cookie)
	})

	// tracking number of visits in the application
	http.HandleFunc("/visit", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("no-visits")

		if err == http.ErrNoCookie {
			http.SetCookie(w, &http.Cookie{
				Name: "no-visits",
				Value: "0",
			})

			return
		}

		visitNumber, err := strconv.Atoi(cookie.Value);
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		currentVisit := strconv.Itoa(visitNumber + 1)
		http.SetCookie(w, &http.Cookie{
			Name: "no-visits",
			Value: currentVisit,
		})
		fmt.Println("cookie is", cookie)
		fmt.Fprintln(w, "howdy this is your " + currentVisit + " visit")

	})

	// delete a cookie
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("no-visits")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther) // redirect user if we have not cookie to set
			return
		}

		c.MaxAge = -1 // this deletes the cookie in the browser as it expires
		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	})

	http.ListenAndServe(":9090", nil)
}