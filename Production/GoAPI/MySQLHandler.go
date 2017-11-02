package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func FindLatest(){


	row := DB.QueryRow("SELECT IFNULL(MAX(Han_ID), 0) Han_ID FROM (SELECT Han_ID FROM HackerNewsDB.Comment UNION ALL SELECT Han_ID FROM HackerNewsDB.Thread) a")
	err := row.Scan(&Hannest_id); if err != nil{
		fmt.Print(err.Error())
	}

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
