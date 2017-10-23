package main

import (
	"github.com/streadway/amqp"
	"log"
)



func InitRabbit(){

}

func SendToRabbit(properties amqp.Publishing, qname string){


	err := CH.Publish(
		"",     // exchange
		qname, // routing key
		false,  // mandatory
		false,  // immediate
		properties, // Properties
	)
	log.Printf(" [x] Sent %s", string(properties.Body))
	if err != nil { panic(err)}

}
