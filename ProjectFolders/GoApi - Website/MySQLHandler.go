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


