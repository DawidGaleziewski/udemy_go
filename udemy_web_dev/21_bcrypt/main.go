package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Email string;
	Fname string;
	Password []byte; // use slice of byte as we will encrypt the password
}

var userDB = map[string]user{}
var sessionDB = map[string]string{} 
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

				// we can validate the pasword using bcrypt to compare entered password and the stored hash
				err := bcrypt.CompareHashAndPassword([]byte(passwd), val.Password) 
				isValidPasswd := err != nil;

				if email == val.Email && isValidPasswd  {
					validated = true;
					userID = key
				}
			}

			if validated {
				id := uuid.NewV4().String() 
				sessionDB[id] = userID 

				cookie := http.Cookie{
					Name: "session",
					Value: id,
				}

				http.SetCookie(w, &cookie)


				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}


			tpl.ExecuteTemplate(w, "login.gohtml", "wrong credentials")
		}
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session");
		if err != nil {
			http.Redirect(w,r, "/login", http.StatusSeeOther)
			return
		}

		if _, ok := sessionDB[cookie.Value]; ok {
			delete(sessionDB, cookie.Value)
		}

		http.SetCookie(w, &http.Cookie{
			Name: "session",
			MaxAge: -1,			
		})

		http.Redirect(w,r, "/login", http.StatusSeeOther)

	})


	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tpl.ExecuteTemplate(w, "register.gohtml", "")
		}

		if r.Method == http.MethodPost {
			r.ParseForm()
			email := r.PostForm.Get("email")
			fname := r.PostForm.Get("fname")
			passwd := r.PostForm.Get("passwd")

			// we convert password string into a slice of byte.
			encryptedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			id := uuid.NewV4().String()


			userDB[id] = user{
				Fname: fname,
				Email: email,
				Password: encryptedPasswd,
			}

			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	})

	// page that is behind a auth	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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