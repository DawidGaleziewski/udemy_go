package main

import (
	"admin_panel/session"
	"admin_panel/user"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var tpl *template.Template
var err error
var db *sql.DB

func init() {
	tpl, err = template.ParseGlob("templates/*.gohtml")
	//tpl.ParseGlob("templates/partials/*.gohtml")
	if err != nil {
		log.Println(err)
	}

}

const PORT = "8080"

func main() {
	fmt.Println("server starting at port", PORT)

	// template views
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tpl.ExecuteTemplate(w, "homepage.gohtml", "")
			return
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tpl.ExecuteTemplate(w, "login.gohtml", "")
			return
		}
	})

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		onGET(w, r, func(w http.ResponseWriter, r *http.Request) {
			userID := strings.Split(r.URL.Path, "/user/")[1]
			fmt.Println(userID)
			var user user.User
			users, err := user.FindBy(db, map[string]string{
				"id": userID,
			})
			if err != nil {
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			fmt.Println("data returned ", users[0])
			tpl.ExecuteTemplate(w, "user_detail.gohtml", users[0])
			return
		})
	})

	http.HandleFunc("/user/new", func(w http.ResponseWriter, r *http.Request) {
		onGET(w, r, func(w http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(w, "user_new.gohtml", "")
			return
		})
	})

	// api
	http.HandleFunc("/api/user/auth", func(w http.ResponseWriter, r *http.Request) {
		// if isPOST(r) {
		// 	return
		// }
		onPOST(w, r, func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			email := r.FormValue("email")
			password := r.FormValue("password")
			// verifiedUser, ok, err :=
			verifiedUser, isAuthn, err := user.User.VerifyCredentials(user.User{}, db, email, password)

			if err != nil {
				log.Println(err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}

			if !isAuthn {
				http.Error(w, "wront credentials", http.StatusUnauthorized)
				return
			}

			id := uuid.New()
			newSession := session.Session{
				ID:               id,
				UserID:           verifiedUser.ID,
				CreationDate:     time.Now(),
				LastActivityDate: time.Now(),
			}

			sessionRecord, err := newSession.Create()

			if err != nil {
				log.Println(err)
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "session_id",
				Value: sessionRecord.ID.String(),
				Path:  "/",
			})

			return
		})
	})

	http.HandleFunc("/api/user/logout", func(w http.ResponseWriter, r *http.Request) {
		onPOST(w, r, func(w http.ResponseWriter, r *http.Request) {
			sessionCookie, err := r.Cookie("session_id")

			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			s := session.Session{}
			sessionDBRecords, err := s.FindBy(map[string]string{
				"id": sessionCookie.Value,
			})

			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if len(sessionDBRecords) > 1 {
				log.Println("db should not return more then one session record")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if len(sessionDBRecords) == 1 {
				_, err := sessionDBRecords[0].Delete()
				if err != nil {
					log.Panicln(err)
				}
			}

			sessionCookie.MaxAge = -1
			http.SetCookie(w, sessionCookie)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		})
	})

	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		onPOST(w, r, func(w http.ResponseWriter, r *http.Request) {
			user.CreateAdminUser(w, r, db)
		})

		return
	})

	http.ListenAndServe(":"+PORT, nil)
}

func onGET(w http.ResponseWriter, r *http.Request, cb func(w http.ResponseWriter, r *http.Request)) {
	if r.Method != http.MethodGet {
		return
	}

	cb(w, r)
}

func onDelete(w http.ResponseWriter, r *http.Request, cb func(w http.ResponseWriter, r *http.Request)) {
	if r.Method != http.MethodDelete {
		return
	}

	cb(w, r)
}

func onPOST(w http.ResponseWriter, r *http.Request, cb func(w http.ResponseWriter, r *http.Request)) {
	if r.Method != http.MethodPost {
		return
	}

	cb(w, r)
}
