package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // we use alias here. This is a throw away alist as we just import this for setup. We dont need no more code from this package
)

var db *sql.DB
var err error
var tpl *template.Template

func init() {
	tpl, err = template.ParseGlob("usersgo.html")
	if err != nil {
		log.Panicln(err)
	}
}

const port = ":8080"

func prepareTables(db *sql.DB) {
	// we can create a table using db.Prepare and sql statment
	// we can also use prepare for any other sql statment
	stmt, err := db.Prepare(`CREATE TABLE blog_users (
		id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(20) NOT NULL,
		email VARCHAR(20) NOT NULL,
		password VARCHAR(50) NOT NULL
	)`)

	if err != nil {
		log.Println("#", err)
	}

	result, err := stmt.Exec()

	if err != nil {
		log.Println("##", err)
	}

	numberOfRowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("lines affected: ", numberOfRowsAffected)

}

func main() {
	fmt.Println("boooting up server on port", port)
	// we just need to provide the name of the driver to the first param and this ugly address into the second one. This is a aws address example
	db, err := sql.Open("mysql", "test_mysql:Test123!@tcp(127.0.0.1:3306)/gosql?charset=utf8")
	prepareTables(db)

	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Ping()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rows, err := db.Query(`SELECT first_name FROM blog_users`)
			if err != nil {
				log.Println(err)
			}

			var cellValue string
			var users []string

			for rows.Next() { // this uses Next same as scanner to go line by line
				err = rows.Scan(&cellValue) // we save the result from each scan
				if err != nil {
					log.Println(err)
				}

				users = append(users, cellValue)

			}

			tpl.ExecuteTemplate(w, "usersgo.html", users)
		}

		if r.Method == http.MethodPost {
			// id := uuid.New().ID()
			stmt, err := db.Prepare(fmt.Sprintf(`INSERT INTO gosql.blog_users VALUES ("Mark", "dawid@gmail.com", "test123!");`))
			if err != nil {
				log.Println(err)
			}

			_, err = stmt.Exec()
			if err != nil {
				log.Panicln(err)
			}

			fmt.Fprintln(w, "user created")
		}
	})

	http.ListenAndServe(port, nil)
}
