package main

import (
	"database/sql"
	"os"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"encoding/json"

)

//dbusername = args[1];
//dbpassword = args[2];
//dbip = args[3];
//rabbituser = args[4];
//rabbitpassword = args[5];
//rabbitip = args[6];

var DB *sql.DB
var CONN *amqp.Connection
var CH *amqp.Channel
var Post_Q amqp.Queue



func main() {


	log.Println("Initializing Database Connection.")
	db, err := sql.Open("mysql", os.Args[1]+":"+os.Args[2]+"@tcp("+os.Args[3]+":3306)/HackerNewsDB")
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
		"TestPost",
		true,
		false,
		false,
		false,
		nil,
	); if err != nil { panic(err)}

	Post_Q = q1
	CH = ch
	CONN = conn


	msgs, err := ch.Consume(
		Post_Q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	); if err != nil { panic(err)}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var request JsonMessage

			if err := json.Unmarshal(d.Body, &request); err != nil {
				log.Println("Couldn't marshal")
			}
			SaveData(request)

			log.Printf("Done")
			d.Ack(false)
		}
	}()
	<-forever


}
