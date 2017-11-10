package main

type ThreadsMessage struct {
	ID int `json:"id"`
	Name string `json:"name"`
	UserID int `json:"userid"`
	//gotta fix this, possible to not have it and use curtime() in insert statement
	Time DateTime `json:"time"`
	Han_ID int `json:"han_id"`
	Post_URL string `json:"post_url"`
}