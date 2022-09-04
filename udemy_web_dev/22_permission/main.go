package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	viewer    = 1
	editor    = 2
	moderator = 3
	admin     = 4
)

type user struct {
	Email    string
	Fname    string
	Password []byte
	Roles    []int
}

type session struct {
	userID string;
	lastActive time.Time
}

var userDB = map[string]user{}
var sessionDB = map[string]session{}

var tpl *template.Template

func init() {
	// create some moc data
	pwd1, _ := bcrypt.GenerateFromPassword([]byte("test1"), bcrypt.MinCost)
	userDB[uuid.NewV4().String()] = user{
		Fname:    "Mike",
		Email:    "MikeAwesome@gmail.com",
		Roles: []int{viewer},
		Password: pwd1,
	}
	userDB[uuid.NewV4().String()] = user{
		Fname:    "Dave",
		Email:    "dave@gmail.com",
		Roles: []int{admin},
		Password: pwd1,
	}

	userDB[uuid.NewV4().String()] = user{
		Fname:    "George",
		Email:    "George@gmail.com",
		Roles: []int{viewer, editor},
		Password: pwd1,
	}
	userDB[uuid.NewV4().String()] = user{
		Fname:    "Jenny",
		Email:    "jenny@gmail.com",
		Roles: []int{viewer, editor},
		Password: pwd1,
	}
	var err error
	tpl, err = template.ParseGlob("templates/*.gohtml")

	if err != nil {
		log.Panicln(err)
	}
}

func main() {

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

				err := bcrypt.CompareHashAndPassword([]byte(passwd), val.Password)
				isValidPasswd := err != nil

				if email == val.Email && isValidPasswd {
					validated = true
					userID = key
				}
			}

			if validated {
				id := uuid.NewV4().String()
				sessionDB[id] = session{
					userID: userID,
					lastActive: time.Now(),
				}

				cookie := http.Cookie{
					Name:  "session",
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
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if _, ok := sessionDB[cookie.Value]; ok {
			delete(sessionDB, cookie.Value)
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			MaxAge: -1,
		})

		http.Redirect(w, r, "/login", http.StatusSeeOther)

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

			encryptedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			id := uuid.NewV4().String()

			userDB[id] = user{
				Fname:    fname,
				Email:    email,
				Password: encryptedPasswd,
				Roles:    []int{viewer},
			}

			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			sessionID, err := r.Cookie("session")

			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			session, ok := sessionDB[sessionID.Value]
			if !ok {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			userData, ok := userDB[session.userID]
			if !ok {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			tpl.ExecuteTemplate(w, "my_acc.gohtml", userData)
			return
		}
	})

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "admin_page.gohtml", userDB)
	})

	http.ListenAndServe(":9090", nil)
}
