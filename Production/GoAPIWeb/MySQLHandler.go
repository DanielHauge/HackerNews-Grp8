package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func FindLatest()string{

	hanid := "error"
	row := DB.QueryRow("SELECT IFNULL(MAX(Han_ID), 0) Han_ID FROM (SELECT Han_ID FROM HackerNewsDB.Comment UNION ALL SELECT Han_ID FROM HackerNewsDB.Thread) a")
	err := row.Scan(&hanid); if err != nil{
		fmt.Print(err.Error())
	}
	return hanid
}

func SqlStatus() bool{

	err := DB.Ping()
	alive := false
	if err != nil{
		log.Printf(err.Error())
	}else
	{
		alive = true
	}

	return alive
}

func VerifyUser(usr UserLogin)bool {
	result := false

	pwd := ""
	row := DB.QueryRow("select Password from User where Name = ?;", usr.Username)
	err := row.Scan(&pwd);
	if err != nil {
		fmt.Print(err.Error())
	}

	if usr.Password == pwd{
		result = true
	}

	return result
}

func QueryLatestStories(dex int, dexto int)LatestStories{
	results := LatestStories{}

	rows, err := DB.Query("SELECT Name, UserID, Time, Post_URL FROM HackerNewsDB.Thread ORDER BY ID DESC LIMIT ?, ?", dex, dexto)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var Name string
		var UserID int
		var Time DateType
		var Post_URL string

		if err := rows.Scan(&Name, &UserID, &Time, &Post_URL); err != nil {
			log.Fatal(err)
		}
		results.Stories = append(results.Stories, Story{Name,UserID, Time,Post_URL})
	}
	return results
}

func QueryAllComments(T_id int, Han_ID int)AllComments{
	results := AllComments{}
	var id int
	var where string
	if Han_ID == 0{
		id = T_id
		where = "ThreadID"
	}else { id = Han_ID; where = "ParentID" }

	rows, err := DB.Query("SELECT Name, UserID, CommentKarma, Time FROM HackerNewsDB.Comment WHERE ? LIKE ? ORDER BY ID DESC", where, id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var Name string
		var UserID int
		var CommentKarma int
		var Time DateType

		if err := rows.Scan(&Name, &UserID, &CommentKarma, &Time); err != nil {
			log.Fatal(err)
		}
		results.Comments = append(results.Comments, Comment{Name,UserID,CommentKarma,Time})
	}
	return results
}
