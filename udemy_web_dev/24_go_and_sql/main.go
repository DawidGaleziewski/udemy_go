package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql/" // we use alias here. This is a throw away alist as we just import this for setup. We dont need no more code from this package
)

var db *sql.DB
var err error

func main(){
	// we just need to provide the name of the driver to the first param and this ugly address into the second one. This is a aws address example
	db, err := sql.Open("mysql", "user:password@tcp(endpoint.address.amazonaws.con:3306)/databaseName?charset=utf8")

	if err != nil {
		log.Println(err)
	}

	defer db.Close()
}