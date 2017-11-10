package main

import "github.com/streadway/amqp"
import "log"

func RabbitMessage(qname string, message string){

	props := amqp.Publishing{
		ContentType: "text/plain; charset=UTF-8",
		Body:        []byte(message),
	}

	err := CH.Publish(
		"",     // exchange
		qname, // routing key
		false,  // mandatory
		false,  // immediate
		props, // Properties
	)
	log.Printf(" [x] Sent %s", string(props.Body))
	if err != nil { log.Println(err.Error())}

}


