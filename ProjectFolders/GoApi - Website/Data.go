package main


type PostRequest struct {
	Username string `json:"username"`
	Post_type string `json:"post_type"`
	Pwd_hash string `json:"pwd_hash"`
	Post_title string `json:"post_title"`
	Post_parrent int `json:"post_parrent"`
	Hanesst_id int `json:"hanesst_id"`
	Post_text string `json:"post_text"`
	Post_url string `json:"post_url"`
}

type HNUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	EmailAddr string `json:"email_addr"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LatestStories struct {
	Stories []PostRequest `json:"stories"`
}