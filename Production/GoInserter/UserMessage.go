package main

type PostRequest struct {
	//limit to 20
	Username string `json:"username"`
	//limit to 20
	Password string `json:"password"`
	//limit to 80
	Email_addr string `json:"email_addr"`
}