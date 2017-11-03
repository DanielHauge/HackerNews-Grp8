package main

import (
	"log"
	"net/http"
	"github.com/rs/cors"
	"database/sql"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)


// go get "github.com/go-sql-driver/mysql"
// go get “github.com/gorilla/mux”
// go get get github.com/streadway/amqp

var DB *sql.DB
var CONN *amqp.Connection
var CH *amqp.Channel
var Post_Q amqp.Queue
var Hannest_id int
var Status string

var (

	promRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_since",
			Help: "The ammount of requests which the api has gotten since the last check.",
		},
	)

)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(promRequests)

}

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
	Hannest_id = 0
	go FindLatest()
	go SetStatus()
	go StartStatusTask()



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

func SetStatus(){

	_, err := http.Get("http://138.197.186.82:15672/api/overview")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		Status = "Down"
	} else{

		if SqlStatus(){
			Status = "Alive"
		} else
		{
			Status = "Update"
		}

	}

}

func StartStatusTask(){
	t := time.NewTicker(time.Second*5)
	for {
		SetStatus()
		<-t.C
	}
}
