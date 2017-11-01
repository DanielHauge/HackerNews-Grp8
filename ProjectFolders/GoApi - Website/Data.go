package main

import (
	"time"
	"strconv"
)

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

type LoggedInUser struct {
	Username string `json:"username"`
	EmailAddr string `json:"email_addr"`
	Karma int `json:"karma"`
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

}

type UpvoteData struct {
	ThreadID int `json:"thread_id"`
	CommentID int `json:"comment_id"`
	Username string `json:"username"`
}

type Story struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Username string `json:"username"`
	Points int `json:"points"`
	Time string `json:"time"`
	Url string `json:"url"`
	CommentAmount int `json:"commentamount"`
}

type Comment struct {
	Id int `json:"id"`
	Comment string `json:"comment"`
	Username string `json:"username"`
	Points int `json:"points"`
	Time string `json:"time"`
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
	/*
	then, err := time.Parse("2006-01-02 15:4:5", time.Time(t).String())
	if err != nil{
		fmt.Println(err)
	}
	*/
	then := time.Time(t)
	duration := 0
	describer := ""
	if int(time.Since(then).Minutes()) < 60{
		duration = int(time.Since(then).Minutes())
		describer = " Minutes Ago"
	}else{
		duration = int(time.Since(then).Hours())
		describer = " Hours Ago"
	}
	if int(time.Since(then).Hours())>24{
		duration = duration/24
		describer = " Days Ago"
	}
	msg := strconv.Itoa(duration) + describer
	return string(msg)
}