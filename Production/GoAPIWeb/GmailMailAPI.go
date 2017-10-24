package main

import (

	"log"

	"net/smtp"
)


func SendEmail(email_addr string, pwd string){
	from := "hngroup8@gmail.com"
	pass := "HackerNews8"

	msg := "From: "+from+"\n"+
		"To: "+ email_addr + "\n"+
			"Subject: HackerNews Password Recovery\n\n"+
				"Your password has been recovered, your password is: "+pwd


	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("",from,pass,"smtp.gmail.com"),from, []string{email_addr}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Println("Sent recovery email")
}



type RecoveryData struct {
	Username string `json:"username"`
}


