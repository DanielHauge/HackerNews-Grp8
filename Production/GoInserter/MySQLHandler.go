package main

import _ "github.com/go-sql-driver/mysql"
import "log"

func SaveData(request JsonMessage){

	if (request.Post_type == "story"){
		SaveStory(request)
	}
	if (request.Post_type == "comment"){
		if (request.Hanesst_id < 0){
			SaveCommentWebSite(request)
		}else {




		}
	}



}

func SaveStory(request JsonMessage){

	stmt, err := DB.Prepare("INSERT INTO HackerNewsDB.ThreadDummy (Name, UserID, Time, Han_ID, Post_URL, Karma) VALUES (?, (SELECT HackerNewsDB.User.ID FROM HackerNewsDB.User WHERE HackerNewsDB.User.Name = ?), curtime(), ?, ?, ?);")
	if err != nil{
		log.Println(err.Error())
	}

	_, err = stmt.Exec(request.Post_title, request.Username, request.Hanesst_id, request.Post_url, 0)
	if err != nil{
		log.Println(err.Error())
	}

}

func SaveCommentWebSite(request JsonMessage){

	stmt, err := DB.Prepare("Insert into HackerNewsDB.CommentDummy (ThreadID, Name, UserID, Karma, Time, Han_ID, PostParrent) values (?, ?, (SELECT HackerNewsDB.User.ID FROM HackerNewsDB.User WHERE HackerNewsDB.User.Name = ?), ?, curtime(), ?, ?)")
	if err != nil{
		log.Println(err.Error())

	}

	_, err = stmt.Exec(request.Post_parent, request.Post_text, request.Username, 0, request.Hanesst_id, request.Post_parent)
	if err != nil{
		log.Println(err.Error())
	}

}

func SaveComment(request JsonMessage, tid int){
	stmt, err := DB.Prepare("Insert into HackerNewsDB.CommentDummy (ThreadID, Name, UserID, Karma, Time, Han_ID, PostParrent) values (?, ?, (SELECT HackerNewsDB.User.ID FROM HackerNewsDB.User WHERE HackerNewsDB.User.Name = ?), ?, curtime(), ?, ?)")
	if err != nil{
		log.Println(err.Error())
	}

	_, err = stmt.Exec(tid, request.Post_text, request.Username, 0, request.Hanesst_id, request.Post_parent)
	if err != nil{
		log.Println(err.Error())
	}
}

func SaveCommentOfComment(request JsonMessage){

	stmt, err := DB.Prepare("Insert into HackerNewsDB.CommentDummy (ThreadID, Name, UserID, Karma, Time, Han_ID, PostParrent) values ((SELECT HackerNewsDB.CommentDummy.ThreadID FROM HackerNewsDB.CommentDummy WHERE HackerNewsDB.CommentDummy.Han_ID = ?), ?, (SELECT HackerNewsDB.User.ID FROM HackerNewsDB.User WHERE HackerNewsDB.User.Name = ?), ?, curtime(), ?, ?)")
	if err != nil{
		log.Println(err.Error())
	}

	_, err = stmt.Exec(request.Post_parent, request.Post_text, request.Username, 0, request.Hanesst_id, request.Post_parent)
	if err != nil{
		log.Println(err.Error())
	}
}

func WorkAroundCheck(request JsonMessage){

	threadid := 0
	row := DB.QueryRow("Select ID from HackerNewsDB.ThreadDummy where Han_ID LIKE ?;", request.Post_parent)
	err := row.Scan(&threadid); if err != nil{
		log.Println(err.Error())
	}

	if threadid>0{
		SaveComment(request, threadid)
	}else {
		SaveCommentOfComment(request)
	}


}