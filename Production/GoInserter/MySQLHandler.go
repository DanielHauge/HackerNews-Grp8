package main


import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

func SaveData(request JsonMessage){

	if (request.Post_type == "story"){
		NewSaveStory(request)

		//SaveStory(request) //// Old way of inserting
	} else if (request.Post_type == "comment"){
		if (request.Hanesst_id < 0){


			SaveCommentWebSite(request) // Old way of inserting
		}else {

			NewSaveComment(request)
			// WorkAroundCheck(request) // Old way of inserting
		}
	} else {
		b, _ := json.Marshal(request)
		go RabbitMessage(Error_q.Name, string(b), "Prepare SaveStory","Was not story or comment!?")
	}


}


func NewSaveStory(request JsonMessage){
	stmt, err := DB.Prepare("INSERT INTO HackerNewsDB.Thread (Name, UserID, Time, Han_ID, Post_URL, Karma) VALUES (?, ?, curtime(), ?, ?, ?);")
	if err != nil{
		log.Println(err.Error())
		log.Println(stmt.Close())
		b, _ := json.Marshal(request)
		go RabbitMessage(Error_q.Name, string(b), "Prepare SaveStory",err.Error())
	}else {

		lid, err := stmt.Exec(request.Post_title, Users[request.Username], request.Hanesst_id, request.Post_url, 0)
		if err != nil{
			log.Println(err.Error())
			log.Println(stmt.Close())
			b, _ := json.Marshal(request)
			go RabbitMessage(Error_q.Name, string(b), "Execute SaveStory",err.Error())
		}else {
			storyid, err := lid.LastInsertId()
			if err != nil{
				log.Print("Error: ", err.Error())
			}
			ThreadIDs[request.Hanesst_id] = storyid
		}
	}
}



func SaveCommentWebSite(request JsonMessage){

	stmt, err := DB.Prepare("Insert into HackerNewsDB.Comment (ThreadID, Name, UserID, Karma, Time, Han_ID, PostParrent) values (?, ?, ?, ?, curtime(), ?, ?)")
	if err != nil{
		log.Println(err.Error())
		b, _ := json.Marshal(request)
		go RabbitMessage(Error_q.Name, string(b), "Prepare CommentWebSite",err.Error())
	}else {
		_, err = stmt.Exec(request.Post_parent, request.Post_text, Users[request.Username], 0, request.Hanesst_id, request.Post_parent)
		if err != nil{
			log.Println(err.Error())
			b, _ := json.Marshal(request)
			go RabbitMessage(Error_q.Name, string(b), "Execute CommentWebSite",err.Error())
		}
	}


}

func NewSaveComment(request JsonMessage){

	threadid := ThreadIDs[request.Post_parent]

	stmt, err := DB.Prepare("Insert into HackerNewsDB.Comment (ThreadID, Name, UserID, Karma, Time, Han_ID, PostParrent) values (?, ?, ?, ?, curtime(), ?, ?)")
	if err != nil{
		log.Println(err.Error())
		log.Println(stmt.Close())
		b, _ := json.Marshal(request)
		go RabbitMessage(Error_q.Name, string(b), "Prepare SaveComment",err.Error())
	}else {

		_, err = stmt.Exec(threadid,request.Post_text, Users[request.Username], 0, request.Hanesst_id, request.Post_parent)
		if err != nil{
			log.Println(err.Error())
			b, _ := json.Marshal(request)
			go RabbitMessage(Error_q.Name, string(b), "Execute SaveComment",err.Error())
		}
	}

	ThreadIDs[request.Hanesst_id] = threadid

}


//// Old way of inserting
/*


func SaveStory(request JsonMessage){

	stmt, err := DB.Prepare("INSERT INTO HackerNewsDB.Thread (Name, UserID, Time, Han_ID, Post_URL, Karma) VALUES (?, (SELECT HackerNewsDB.User.ID FROM HackerNewsDB.User WHERE HackerNewsDB.User.Name = ? LIMIT 1), curtime(), ?, ?, ?);")
	if err != nil{
		log.Println(err.Error())
		b, _ := json.Marshal(request)
		go RabbitMessage(Error_q.Name, string(b), "Prepare SaveStory",err.Error())
	}else {

		_, err = stmt.Exec(request.Post_title, request.Username, request.Hanesst_id, request.Post_url, 0)
		if err != nil{
			log.Println(err.Error())
			b, _ := json.Marshal(request)
			go RabbitMessage(Error_q.Name, string(b), "Execute SaveStory",err.Error())
		}
	}

}



func SaveComment(request JsonMessage){

	stmt, err := DB.Prepare("Insert into HackerNewsDB.Comment (ThreadID, Name, UserID, Karma, Time, Han_ID, PostParrent) values ((Select ID from HackerNewsDB.Thread where Han_ID LIKE ?), ?, (SELECT HackerNewsDB.User.ID FROM HackerNewsDB.User WHERE HackerNewsDB.User.Name = ? LIMIT 1), ?, curtime(), ?, ?)")
	if err != nil{
		log.Println(err.Error())
		b, _ := json.Marshal(request)
		go RabbitMessage(Error_q.Name, string(b), "Prepare SaveComment",err.Error())
	}else {

		_, err = stmt.Exec(request.Post_parent ,request.Post_text, request.Username, 0, request.Hanesst_id, request.Post_parent)
		if err != nil{
			log.Println(err.Error())
			b, _ := json.Marshal(request)
			go RabbitMessage(Error_q.Name, string(b), "Execute SaveComment",err.Error())
		}
	}


}

func SaveCommentOfComment(request JsonMessage, tid int){

	stmt, err := DB.Prepare("Insert into HackerNewsDB.Comment (ThreadID, Name, UserID, Karma, Time, Han_ID, PostParrent) values (?, ?, (SELECT HackerNewsDB.User.ID FROM HackerNewsDB.User WHERE HackerNewsDB.User.Name = ? LIMIT 1), ?, curtime(), ?, ?)")
	if err != nil{
		log.Println(err.Error())
		b, _ := json.Marshal(request)
		go RabbitMessage(Error_q.Name, string(b), "Prepare CommentofComment",err.Error())
	} else {
		_, err = stmt.Exec(tid, request.Post_text, request.Username, 0, request.Hanesst_id, request.Post_parent)
		if err != nil{
			log.Println(err.Error())
			b, _ := json.Marshal(request)
			go RabbitMessage(Error_q.Name, string(b), "Execute CommentofComment",err.Error())
		}
	}


}

func WorkAroundCheck(request JsonMessage){

	threadid := 0
	row := DB.QueryRow("Select ThreadID from HackerNewsDB.Comment where Han_ID LIKE ?;", request.Post_parent)
	err := row.Scan(&threadid); if err != nil{
		log.Println(err.Error()+" : Didn't find a Comment with that Han_ID, this means it is not a comment of comment!")
		SaveComment(request)
	}else {
		SaveCommentOfComment(request, threadid)
	}
}
*/


func SetupUserDictionairy()map[string]int{
	result := map[string]int{}

	rows, err := DB.Query("SELECT ID, Name FROM HackerNewsDB.User;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var Name string
		var id int
		if err := rows.Scan(&id ,&Name); err != nil {
			log.Fatal(err)
		}
		result[Name] = id
	}

	return result

}

func SetupThreadIDdictionairy()map[int]int64{
	result := map[int]int64{}

	rows, err := DB.Query("SELECT ID, Han_ID FROM HackerNewsDB.Thread;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var ThreadID int64
		var Han_ID int
		if err := rows.Scan(&ThreadID ,&Han_ID); err != nil {
			log.Fatal(err)
		}
		result[Han_ID] = ThreadID
	}

	rows1, err := DB.Query("SELECT ThreadID, Han_ID FROM HackerNewsDB.Comment;")
	if err != nil {
		log.Fatal(err)
	}

	for rows1.Next() {
		var ThreadID int64
		var Han_ID int
		if err := rows1.Scan(&ThreadID ,&Han_ID); err != nil {
			log.Fatal(err)
		}
		result[Han_ID] = ThreadID
	}

	return result
}