package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"

	"net/http"
	"time"

	"admin_panel/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var tpl *template.Template
var err error
var db *sql.DB

func init() {
	tpl, err = template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Println(err)
	}

}

const PORT = "8080"

func main() {
	fmt.Println("server starting at port", PORT)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tpl.ExecuteTemplate(w, "homepage.gohtml", "")
			return
		}

	})

	http.HandleFunc("/api/user/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			return
		}
	})

	http.HandleFunc("/user/new", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tpl.ExecuteTemplate(w, "user_new.gohtml", "")
			return
		}
	})

	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")

			fmt.Println("form name", name)

			ID, err := uuid.NewUUID()
			if err != nil {
				log.Println(err)
			}

			user := user.User{
				ID:           ID,
				Name:         name,
				Email:        email,
				Password:     password,
				CreationDate: time.Now(),
			}

			userData, validationErrors, err := user.Register(db)

			if err != nil {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			if len(validationErrors) > 0 {
				errMsg := ""
				for _, value := range validationErrors {
					errMsg += value + "\n"
				}
				http.Error(w, errMsg, http.StatusBadRequest)
				return
			}

			err = tpl.ExecuteTemplate(w, "user_detail.gohtml", userData)
			if err != nil {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			return
		}

		return
	})

	http.ListenAndServe(":"+PORT, nil)
}
