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

func GetUsername(userid int)string{
	un := "0"
	row := DB.QueryRow("select Name from User where ID = ?;", userid)
	err := row.Scan(&un); if err != nil{
		fmt.Print(err.Error())
	}
	return un
}

func GetSingleStory(threadid int)Story{
	var st Story
	var userid int
	var date DateType
	row := DB.QueryRow("SELECT ID, Name, UserID, Time, Post_URL FROM HackerNewsDB.Thread WHERE ID LIKE ?;", threadid)
	err := row.Scan(&st.Id, &st.Title, &userid, &date, &st.Url); if err != nil{
		fmt.Print(err.Error())
	}
	st.Username = GetUsername(userid)
	st.Time = date.String()
	log.Print(st.Username)
	return st
}

func GetUserID(username string)int {

	uid := 0
	row := DB.QueryRow("select ID from User where Name = ?;", username)
	err := row.Scan(&uid); if err != nil{
		fmt.Print(err.Error())
	}
	return uid
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

func GetRecoveryInformation(username string)(string,string){

	pwd := ""
	email := ""

	row := DB.QueryRow("select Password, Email  from User where Name = ?;", username)
	err := row.Scan(&pwd, &email);
	if err != nil {
		fmt.Print(err.Error())
	}
	return pwd, email

}

func QueryLatestStories(dex int, dexto int)LatestStories{
	results := LatestStories{}

	rows, err := DB.Query("SELECT ID, Name, UserID, Time, Post_URL FROM HackerNewsDB.Thread ORDER BY ID DESC LIMIT ?, ?", dex, dexto)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var Name string
		var UserID int
		var Time DateType

		var Post_URL string
		var ID int

		if err := rows.Scan(&ID ,&Name, &UserID, &Time, &Post_URL); err != nil {
			log.Fatal(err)
		}
		results.Stories = append(results.Stories, Story{ID,Name,GetUsername(UserID), Time.String(),Post_URL})
	}
	return results
}

func QueryAllComments(T_id int)[]Comment{
	results := []Comment{}
	var id int

	rows, err := DB.Query("SELECT Name, UserID, CommentKarma, Time FROM HackerNewsDB.Comment WHERE ID LIKE ? ORDER BY ID DESC", id)
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
		results = append(results, Comment{Name,GetUsername(UserID),CommentKarma,Time})
	}
	return results
}

func ChangePassword(newpwd string, id int)error{

	stmt, err := DB.Prepare("UPDATE HackerNewsDB.User SET Password = ? WHERE ID = ?;")
	if err != nil{
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(newpwd, id)
	if err != nil{
		fmt.Print(err.Error())
	}

	return err
}