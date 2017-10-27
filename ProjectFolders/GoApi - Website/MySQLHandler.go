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

func CountComments(threadid int)int{
	amount := 0
	row := DB.QueryRow("SELECT COUNT(ID) AS amount FROM HackerNewsDB.Comment WHERE ThreadID LIKE ?;", threadid)
	err := row.Scan(&amount); if err != nil{
		fmt.Print(err.Error())
	}
	log.Print(amount)
	return amount
} /// Not used currently

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

	rows, err := DB.Query("SELECT HackerNewsDB.Thread.ID, HackerNewsDB.Thread.Name, HackerNewsDB.Thread.Time, HackerNewsDB.Thread.Post_URL, HackerNewsDB.Thread.Karma, HackerNewsDB.User.Name, COUNT(HackerNewsDB.Comment.ID) as commentamount FROM HackerNewsDB.Thread LEFT OUTER JOIN HackerNewsDB.User ON HackerNewsDB.Thread.UserID = HackerNewsDB.User.ID LEFT OUTER JOIN HackerNewsDB.Comment ON HackerNewsDB.Thread.ID = HackerNewsDB.Comment.ThreadID GROUP BY HackerNewsDB.Thread.ID ORDER BY ID DESC LIMIT ?, ?", dex, dexto)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var Name string
		var Username string
		var Time DateType
		var Karma int
		var Post_URL string
		var ID int
		var comamount int

		if err := rows.Scan(&ID ,&Name, &Time, &Post_URL, &Karma, &Username, &comamount); err != nil {
			log.Fatal(err)
		}
		results.Stories = append(results.Stories, Story{ID,Name,Username, Karma,Time.String(), Post_URL, comamount})
	}

	return results
}

func QueryAllComments(T_id int)[]Comment{
	results := []Comment{}

	rows, err := DB.Query("SELECT HackerNewsDB.Comment.ID, HackerNewsDB.Comment.Name, HackerNewsDB.User.Name, HackerNewsDB.Comment.Karma, HackerNewsDB.Comment.Time FROM HackerNewsDB.Comment LEFT OUTER JOIN HackerNewsDB.User ON HackerNewsDB.Comment.UserID = HackerNewsDB.User.ID WHERE HackerNewsDB.Comment.ThreadID LIKE ? ORDER BY HackerNewsDB.Comment.ID DESC", T_id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var Name string
		var username string
		var CommentKarma int
		var Time DateType
		var id int
		if err := rows.Scan(&id ,&Name, &username, &CommentKarma, &Time); err != nil {
			log.Fatal(err)
		}
		results = append(results, Comment{id,Name,username,CommentKarma,Time.String()})
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