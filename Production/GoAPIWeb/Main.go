package main

import (
	"log"
	"net/http"
	"github.com/rs/cors"
	"database/sql"
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

// ARGS: 1= db username, 2= db password, 3 = db ip, 4 = rabbit user, 5= rabbit password, 6= rabbit ip, 7 = emailuser, 8 = emailpassword
// go get "github.com/go-sql-driver/mysql"
// go get “github.com/gorilla/mux”
// go get get github.com/streadway/amqp
// go get github.com/rs/cors
// go get -u gopkg.in/russross/blackfriday.v2

var DB *sql.DB
var CONN *amqp.Connection
var CH *amqp.Channel
var Post_Q amqp.Queue
var User_Q amqp.Queue

func main() {

	log.Println("Initializing API.")
	router := NewRouter()
	handler := cors.Default().Handler(router)


	log.Println("Initializing Database Connection.")
	db, err := sql.Open("mysql", os.Args[1]+":"+os.Args[2]+"@tcp("+os.Args[3]+":3306)/HackerNewsDB?parseTime=True")
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
	conn, err := amqp.Dial("amqp://"+os.Args[4]+":"+os.Args[5]+"@"+os.Args[6]+""); if err != nil { panic(err) }
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

	q2, err := ch.QueueDeclare(
		"HNUser",
		true,
		false,
		false,
		false,
		nil,
	); if err != nil { panic(err)}

	User_Q = q2
	Post_Q = q1
	CH = ch
	CONN = conn

	log.Println("Initializing Server!.")
	log.Fatal(http.ListenAndServe(":9191", handler))
	log.Println("Server running on port 9191.")



}

