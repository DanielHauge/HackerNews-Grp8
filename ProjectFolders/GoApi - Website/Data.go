package main

import "time"

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

type PasswordChangeData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NewPassword string `json:"new_password"`
}

type StoryWithComments struct {
	Thread Story `json:"thread"`
	Acomments []Comment `json:"comments"`
}

type LatestStories struct {
	Stories []Story `json:"stories"`
	CommentAmount int `json:"comment_amount"`
}

type Story struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Username string `json:"username"`
	Time string `json:"time"`
	Url string `json:"url"`
}

type Comment struct {
	Comment string `json:"comment"`
	Username string `json:"username"`
	Points int `json:"points"`
	Time DateType `json:"time"`
}

type StoryRequest struct {
	Dex int `json:"dex"`
	DexTo int `json:"dex_to"`
}

type CommentsRequest struct {
	ThreadID int `json:"thread_id"`
	Han_id int `json:"han_id"`
}

type DateType time.Time

func (t DateType) String() string {
	return time.Time(t).String()
}