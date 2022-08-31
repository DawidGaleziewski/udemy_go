package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	Email string;
	Fname string;
	Password string;
}

var userDB = map[string]user{}
var sessionDB = map[string]string{} // we use composite literal to create a empty map. This is a alternative for useing make map.
var tpl *template.Template

func init(){
	var err error
	tpl, err = template.ParseGlob("templates/*.gohtml")

	if err != nil {
		log.Panicln(err)
	}
}

func main(){

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if IsLoggedIn(r) {
			http.Redirect(w, r, "/my-account", http.StatusSeeOther)
			return
		}
		
		if r.Method == http.MethodGet {
			tpl.ExecuteTemplate(w, "login.gohtml", "")
			return
		}

		if r.Method == http.MethodPost {
			r.ParseForm()
			email := r.PostForm.Get("email")
			passwd := r.PostForm.Get("passwd")

			var validated = false
			var userID string

			for key, val := range userDB {
				fmt.Println("iterate", key, val)
				if email == val.Email && passwd == val.Password {
					validated = true;
					userID = key
				}
			}

			if validated {
				id := uuid.NewV4().String() // we create new id session
				sessionDB[id] = userID // store new session id with relation to the user

				// we create a session cookie with a id to that session
				cookie := http.Cookie{
					Name: "session",
					Value: id,
				}

				http.SetCookie(w, &cookie) // cookie is set in response


				http.Redirect(w, r, "/my-account", http.StatusSeeOther)
				return
			}


			tpl.ExecuteTemplate(w, "login.gohtml", "wrong credentials")
		}
	})


	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tpl.ExecuteTemplate(w, "register.gohtml", "")
		}

		// we get the data from form
		if r.Method == http.MethodPost {
			r.ParseForm()
			email := r.PostForm.Get("email")
			fname := r.PostForm.Get("fname")
			passwd := r.PostForm.Get("passwd")

			id := uuid.NewV4().String() // we create new id for a user

			// we create new entry for that use id with data provided
			userDB[id] = user{
				Fname: fname,
				Email: email,
				Password: passwd,
			}

			// now we can redirect user to login
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	})

	
	http.HandleFunc("/my-account", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			sessionID, err := r.Cookie("session")

			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			userID, ok := sessionDB[sessionID.Value]
			if !ok {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return	
			}

			userData, ok := userDB[userID]
			if !ok {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return	
			}

			tpl.ExecuteTemplate(w, "my_acc.gohtml", userData)
			return
		}
	})

	http.ListenAndServe(":9090", nil)
}