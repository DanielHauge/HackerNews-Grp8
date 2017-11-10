package main


type JsonMessage struct {
	Username string `json:"username"`
	Post_type string `json:"post_type"`
	Pwd_hash string `json:"pwd_hash"`
	Post_title string `json:"post_title"`
	Post_parent int `json:"post_parent"`
	Hanesst_id int `json:"hanesst_id"`
	Post_text string `json:"post_text"`
	Post_url string `json:"post_url"`
}

type PostRequest struct {
	//limit to 20
	Username string `json:"username"`
	//limit to 20
	Password string `json:"password"`
	//limit to 80
	Email_addr string `json:"email_addr"`
}