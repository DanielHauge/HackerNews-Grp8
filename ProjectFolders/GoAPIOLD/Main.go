package main

import (
	"log"
	"net/http"


	// go get "github.com/go-sql-driver/mysql"
	//// go get "github.com/gin-gonic/gin" // Only for specific routing not using this
	// go get “github.com/gorilla/mux”

	"database/sql"
	"fmt"
)

var DB *sql.DB

func main() {

	router := NewRouter()

	db, err := sql.Open("mysql", "myuser:HackerNews8@tcp(46.101.103.163:3306)/HackerNewsDB")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	DB = db;

	log.Fatal(http.ListenAndServe(":8787", router))
}
