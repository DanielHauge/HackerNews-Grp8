package main

import (
	"database/sql"
	"os"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"encoding/json"

	"time"
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
var Error_q amqp.Queue
var Users map[string]int
var ThreadIDs map[int]int64



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
	Users = SetupUserDictionairy()
	ThreadIDs = SetupThreadIDdictionairy()


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
		"HNError",
		true,
		false,
		false,
		false,
		nil,
	); if err != nil { panic(err)}

	Post_Q = q1
	Error_q = q2
	CH = ch
	CONN = conn



	msgs, err := ch.Consume(Post_Q.Name,"",false,false,false,false,nil,); if err != nil { panic(err)}

	forever := make(chan bool)

	ticker := time.NewTicker(10 * time.Minute)
	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
			forever<-true

			}
		}
	}(ticker)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var request JsonMessage

			if err := json.Unmarshal(d.Body, &request); err != nil {
				log.Println("Couldn't marshal")
				go RabbitMessage("HNError", string(d.Body), "Unmarshal of rabbitmq response", err.Error())
			} else {
				SaveData(FillInBlanks(request))
				log.Printf("Done")
				d.Ack(false)
			}
		}
	}()
	<-forever


}

