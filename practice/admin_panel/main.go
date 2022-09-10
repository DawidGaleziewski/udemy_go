package main

import (
	"admin_panel/user"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"strings"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

	// template views
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tpl.ExecuteTemplate(w, "homepage.gohtml", "")
			return
		}

	})
	http.HandleFunc("/user/new", func(w http.ResponseWriter, r *http.Request) {
		onGET(w, r, func(w http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(w, "user_new.gohtml", "")
			return
		})
	})

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		onGET(w, r, func(w http.ResponseWriter, r *http.Request) {
			userID := strings.Split(r.URL.Path, "/user/")[1]
			fmt.Println(userID)
			var user user.User
			userData, err := user.Find(db, map[string]string{})
			if err != nil {
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			tpl.ExecuteTemplate(w, "user_detail.gohtml", userData)
			return
		})
	})

	// api
	http.HandleFunc("/api/user/auth", func(w http.ResponseWriter, r *http.Request) {
		// if isPOST(r) {
		// 	return
		// }
	})

	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		// onGET(w, r, func(w http.ResponseWriter, r *http.Request) {
		// 	user.GetUserDetail(w, r, *db)
		// })

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

func onPOST(w http.ResponseWriter, r *http.Request, cb func(w http.ResponseWriter, r *http.Request)) {
	if r.Method != http.MethodPost {
		return
	}

	cb(w, r)
}