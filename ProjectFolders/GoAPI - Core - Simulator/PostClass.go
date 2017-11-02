package main


type PostRequest struct {
	Username string `json:"username"`
	Post_type string `json:"post_type"`
	Pwd_hash string `json:"pwd_hash"`
	Post_title string `json:"post_title"`
	Post_parrent int `json:"post_parrent"`
	Hanesst_id int `json:"hanesst_id"`
	Post_text string `json:"post_text"`
}

