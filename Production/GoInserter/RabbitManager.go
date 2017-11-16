package main

import "github.com/streadway/amqp"
import "log"

func RabbitMessage(qname string, message string, headermessage string, errormessage string){

	props := amqp.Publishing{
		ContentType: "text/plain; charset=UTF-8",
		Body:        []byte(message),
		Headers:amqp.Table{"Error-Message":errormessage,"Error-in":headermessage},
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


