package main

import (
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)


func GetUserID(username string)string {

	uid := "0"
	row := DB.QueryRow("select ID from User where Name = ?;", username)
	err := row.Scan(&uid); if err != nil{
		fmt.Print(err.Error())
	}



	return uid
}

func FindLatest()string{

	hanid := "error"
	row := DB.QueryRow("SELECT IFNULL(MAX(Han_ID), 0) Han_ID FROM (SELECT Han_ID FROM HackerNewsDB.Comment UNION ALL SELECT Han_ID FROM HackerNewsDB.Thread) a")
	err := row.Scan(&hanid); if err != nil{
		fmt.Print(err.Error())
	}
	return hanid
}


func SaveData(request PostRequest) {



	if request.Post_type=="story"{

		var Name = request.Post_title
		var UserID = len(request.Username)      //GetUserID(request.Username)
		var Time = time.Now()
		var Han_ID = request.Hanesst_id

		stmt, err := DB.Prepare("insert into Thread (Name, UserID, Time, Han_ID) VALUES(?,?,?,?);")
		if err != nil{
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(Name, UserID, Time, Han_ID)
		if err != nil{
			fmt.Print(err.Error())
		}


	}
	if request.Post_type=="comment"{

		var ThreadID = request.Post_parrent
		var Comment = request.Post_text
		var UserID = len(request.Username)
		var Time = time.Now()
		var CommentKarma = request.Post_parrent
		var Han_ID = request.Hanesst_id

		stmt, err := DB.Prepare("insert into Comment (ThreadID, Comment, UserID, CommentKarma, Time, Han_ID) VALUES(?,?,?,?,?,?);")
		if err != nil{
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(ThreadID, Comment, UserID, CommentKarma, Time, Han_ID)
		if err != nil{
			fmt.Print(err.Error())
		}


	}


}