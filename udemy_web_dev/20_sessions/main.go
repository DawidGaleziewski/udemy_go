package main

import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
)

func main(){
	var userDB = make(map[string]int32)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session") // we can create a basic session cookie with uuid

		if err != nil { // for when no session was found. We set a new session. Normally this would be done by using login
			id := uuid.NewV4().String()
			cookie = &http.Cookie{
				Name: "session",
				Value: id,
				// Secure: true, // In normal session we would probably want to use secure conncetion
				HttpOnly: true,
			}
			userDB[id] = 0 // we set number of visits for each id

			http.SetCookie(w, cookie) // set the cookie so the session with the uuid is stored in browser
		}

		id := cookie.Value // we can take the id from the cookie to recognise who is performing a action
		visits, ok := userDB[id]
		if !ok {
			fmt.Println("no such session")
		}

		userDB[id] = visits + 1
		fmt.Println(userDB)
	})

	http.ListenAndServe(":9090", nil)
}