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

func FillInBlanks(request JsonMessage)JsonMessage{

	if (len(request.Username)==0){
		request.Username = "UnknownUser"
	}
	if len(request.Post_text)==0{
		request.Post_text = "This comment has been deleted"
	}
	if len(request.Post_url)==0{
		request.Post_url = "http://165.227.151.217:8080/threads"
	}
	if len(request.Post_title)==0{
		request.Post_title = "This link title has been deleted, will link to home"
	}

	return request
}