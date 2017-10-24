package main

import (
	"log"
	"net/http"
	"github.com/rs/cors"
	"database/sql"
	"fmt"
	"github.com/streadway/amqp"
)


// go get "github.com/go-sql-driver/mysql"
// go get “github.com/gorilla/mux”
// go get get github.com/streadway/amqp

var DB *sql.DB
var CONN *amqp.Connection
var CH *amqp.Channel
var Post_Q amqp.Queue

func main() {

	log.Println("Initializing API.")
	router := NewRouter()
	handler := cors.Default().Handler(router)


	log.Println("Initializing Database Connection.")
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

	log.Println("Initializing RabbitMQ Server Connections and Channels.")
	conn, err := amqp.Dial("amqp://admin:password@138.197.186.82"); if err != nil { panic(err) }
	defer conn.Close()

	ch, err := conn.Channel(); if err != nil { panic(err) }
	defer ch.Close()

	q1, err := ch.QueueDeclare(
		"HNPost",
		true,
		false,
		false,
		false,
		nil,
	); if err != nil { panic(err)}

	Post_Q = q1
	CH = ch
	CONN = conn


	log.Println("Initializing Server!.")
	log.Fatal(http.ListenAndServe(":8787", handler))
	log.Println("Server running on port 8787.")

}
