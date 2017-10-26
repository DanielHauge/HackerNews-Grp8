package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	//"github.com/microcosm-cc/bluemonday" // go get github.com/russross/blackfriday-tool
	"gopkg.in/russross/blackfriday.v2" // go get -u gopkg.in/russross/blackfriday.v2
	"log"
	"strconv"
)

func setheader (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin ,Accept, Content-Type, Content-Length, Accept-Encoding")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else { w.Header().Set("Access-Control-Allow-Origin", "*")}
}

func Index2(w http.ResponseWriter, r *http.Request){

	input := "# Welcome!\n" +
		"\n" +
		"This api will accepts the following calls:\n" +
		"\n\n" +
		"## PostStory\n" +
		"This call is to post stories or comments, it will run concurrently, but is not guarenteed to be succesfully insertet. But if it can be succesfull insertet, it will at some point. \n\n"+
		"- Route: /post\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: " +
			"\n\n>``` {\"username\": \"<string>\",\"post_type\": \"<string>\", \"pwd_hash\": \"<string>\",\"post_title\": \"<string>\",'post_url\": \"<string>\", \"post_parent\": <int>, \"hanesst_id\": <int>, \"post_text\": \"<string>\"}```\n\n" +
		"- Example: \n\n > ```{\"post_title\": \"NYC Developer Dilemma\", \"post_text\": \"\", \"hanesst_id\": 4, \"post_type\": \"story\", \"post_parent\": -1, \"username\": \"onebeerdave\", \"pwd_hash\": \"fwozXFe7g0\",  \"post_url\": \"http://avc.blogs.com/a_vc/2006/10/the_nyc_develop.html\"}```"+
		"\n\n" +
		"## GetLatest\n" +
		"This call will return the latest ingested story or comment which was sent by the simulator program. \n\n"+
		"- Route: /post\n" +
		"- Method type: GET\n" +
		"\n" +
		"## GetStatus\n" +
		"This call is only used by the core API to check if status is okay. \n\n"+
		"- Route: /status\n" +
		"- Method type: GET\n" +
		"\n" +
		"## CreateUser\n" +
		"This call is to create users, it will run concurrently, but is not guarenteed to be succesfully insertet. But if it can be succesfull insertet, it will at some point. \n\n"+
		"- Route: /create\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: \n" +
			"\n\n> ```{\"username\": \"<string>\",\"password\": \"<string>\",\"email_addr\": \"<string>\"}```\n" +
		"\n- Example: \n\n>```{\"username\": \"Retrospective\",\"password\": \"Th1sp4ssw0rdW1llN3verG3tH4ck3d\",\"email_addr\": \"Retrospective@icloud.com\"}```"+
		"\n\n" +
		"## Login\n" +
		"This call is to verify users, it will see if username and password is correct and respond accordingly. \n\n"+
		"- Route: /login\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: \n\n> ```{\"username\": \"<string>\",\"password\": \"<string>\"}```\n\n" +
		"- Example: \n\n> ```{\"username\": \"farmer\",\"password\": \"xMBVi4fAO5\"}```\n"+
		"\n\n" +
		"## GetLatestStories\n" +
		"This call will return an array of stories. dex:0 dex_to 100 will return last 100 threads, dex:100,dex_to:200 will return the 100 threads before the very newest 100 threads \n\n"+
		"- Route: /stories\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: " +
		"\n\n>``` {\"dex\": <int>,\"dex_to\": <int>} ```\n\n" +
		"- Example: \n\n > ```{\"dex\": 0,\"dex_to\": 100}```"+
		"\n\n" +
		"## GetStoryByID\n" +
		"This call is not implementet yet, but it works and will return a static story. \n\n"+
		"- Route: /stories/{storyid}\n" +
		"- Method type: GET\n" +
		"\n" +
		"## GetComments\n" +
		"This call is to create users, it will run concurrently, but is not guarenteed to be succesfully insertet. But if it can be succesfull insertet, it will at some point. \n\n"+
		"- Route: /comments/{storyid}\n" +
		"- Method type: GET\n" +
		"\n\n" +
		"## RecoverPassword\n" +
		"This call will send an email to the address which is linked with the username provided. \n\n"+
		"- Route: /recover\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: " +
		"\n\n>``` {\"username\": \"<string>\"} ```\n\n" +
		"- Example: \n\n > ```{\"username\": \"farmer\"}```"+
		"\n\n"+
		"## UpdatePassword\n" +
		"This call will change the password of the user. \n\n"+
		"- Route: /update\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: " +
		"\n\n>``` {\"username\": \"<string>\", \"password\": \"<string>\", \"new_password\": \"<string>\"} ```\n\n" +
		"- Example: \n\n > ```{\"username\": \"farmer\", \"password\": \"currentpassword\", \"new_password\": \"newpassword\"}```"+
		"\n\n"
	output := blackfriday.Run([]byte(input))
	w.Write(output)

}

func Index(w http.ResponseWriter, r *http.Request) {
	setheader(w, r)
	fmt.Fprintln(w, "Welcome!\n" +
					"\n" +
					"This api will accepts the following calls:\n" +
					"\n\n" +
					"PostStory\n" +
					"Route: /post\n" +
					"Method type: POST\n" +
					"Accepts Json & text/plain Syntax: {\"username\": \"<string>\",\"post_type\": \"<string>\", \"pwd_hash\": \"<string>\",\"post_title\": \"<string>\",'post_url\": \"<string>\", \"post_parent\": <int>, \"hanesst_id\": <int>, \"post_text\": \"<string>\"}\n" +
					"Example: {\"post_title\": \"NYC Developer Dilemma\", \"post_text\": \"\", \"hanesst_id\": 4, \"post_type\": \"story\", \"post_parent\": -1, \"username\": \"onebeerdave\", \"pwd_hash\": \"fwozXFe7g0\",  \"post_url\": \"http://avc.blogs.com/a_vc/2006/10/the_nyc_develop.html\"}"+
					"\n\n" +
					"GetLatest\n" +
					"Route: /post\n" +
					"Method type: GET\n" +
					"\n" +
					"GetStatus\n" +
					"Route: /status\n" +
					"Method type: GET\n" +
					"\n" +
					"CreateUser\n" +
					"Route: /create\n" +
					"Method type: POST\n" +
					"Accepts Json & text/plain Syntax: {\"username\": \"<string>\",\"password\": \"<string>\",\"email_addr\": \"<string>\"}\n" +
					"Example: {\"username\": \"Retrospective\",\"password\": \"Th1sp4ssw0rdW1llN3verG3tH4ck3d\",\"email_addr\": \"Retrospective@icloud.com\"}"+
					"\n\n" +
					"Login\n" +
					"Route: /login\n" +
					"Method type: POST\n" +
					"Accepts Json & text/plain Syntax: {\"username\": \"<string>\",\"password\": \"<string>\"}\n" +
					"Example: {\"username\": \"farmer\",\"password\": \"xMBVi4fAO5\"}"+
					"\n\n" +
					"GetLatestStories\n" +
					"Route: /stories\n" +
					"Method type: POST\n" +
					"Accepts text/plain Syntax: <int>\n" +
					"Example: 8"+
					"\n" +
					"GetStoryByID\n" +
					"Route: /stories/{storyid}\n" +
					"Method type: GET\n" +
					"\n" +
					"GetLatestStories\n" +
					"Route: /comments/{storyid}\n" +
					"Method type: GET\n" +
					"\n" +
						" ")
}

func GetLatest(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	setheader(w, r)

	input := FindLatest()
	w.Write([]byte(input))
}

func PostStory(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	setheader(w, r)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	go func() {
		/// Implement MySQL
		props := amqp.Publishing{
			ContentType: "application/json; charset=UTF-8",
			Body:        body,
		}

		SendToRabbit(props, Post_Q.Name)
	}()


	fmt.Fprint(w, "Publishing to RQ for DB Insertion")
}

func GetStatus(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Alive")
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	setheader(w, r)



		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}

	go func() {
		/// Implement MySQL
		props := amqp.Publishing{
			ContentType: "application/json; charset=UTF-8",
			Body:        body,
		}
		SendToRabbit(props, User_Q.Name)

	}()
	fmt.Fprint(w, "Publishing to RQ for DB Insertion")
}

func Login(w http.ResponseWriter, r *http.Request){

	setheader(w, r)

	var usr UserLogin
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &usr); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}


	if VerifyUser(usr) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Succesfull - User is logged in")
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, "Error, incorrect username and/or password")
	}


}

func GetStoryByID(w http.ResponseWriter, r *http.Request){
	setheader(w, r)

	vars := mux.Vars(r)
	id := vars["storyid"]

	realid, err := strconv.Atoi(id); if err != nil {log.Println("Error in parsing int"); fmt.Fprint(w, "Error in parsing int, you should only use ints has ID")}
	log.Print(realid)

	story := GetSingleStory(realid)
	log.Print(story.Username)
	/// Query data base and set req to correct data format.

	comments := QueryAllComments(realid)
	final := StoryWithComments{story, comments}
	msgs, err := json.Marshal(final); if err != nil{ panic(err) }

	fmt.Fprint(w, string(msgs))
}

func GetLatestStories(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	var AllStories LatestStories
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	var req StoryRequest
	if err := json.Unmarshal(body, &req); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	AllStories = QueryLatestStories(req.Dex,req.DexTo)

	msgs, err := json.Marshal(AllStories); if err != nil{ panic(err) }
	fmt.Fprint(w, string(msgs))
}

func GetComments(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["storyid"]

	threadid, err := strconv.Atoi(id); if err != nil {log.Println("Error in parsing int"); fmt.Fprint(w, "Error in parsing int, you should only use ints has ID")}
	comments := QueryAllComments(threadid)
	msgs, err := json.Marshal(comments); if err != nil{ panic(err) }

	fmt.Fprint(w, string(msgs))


}

func StartRecovery(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	var recinfo RecoveryData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &recinfo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	go func() {
	pwd, email := GetRecoveryInformation(recinfo.Username)
	log.Printf(recinfo.Username)
	log.Printf(email)

		SendEmail("Daniel.f.hauge@icloud.com", pwd)

	}()
	fmt.Fprint(w, "A Email has been sent to the specified email linked with the username provided")


}



func UpdatePassword(w http.ResponseWriter, r *http.Request){
	setheader(w, r)

	var change PasswordChangeData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &change); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if VerifyUser(UserLogin{change.Username,change.Password}){
		err := ChangePassword(change.NewPassword, GetUserID(change.Username))
		if err != nil{
			fmt.Fprint(w, "There was an error which might have prevented the password to change, password"+err.Error())
		}
		fmt.Fprint(w, "Succesfully Changed Password")
	}else {
		fmt.Fprint(w, "Error in credentials, wrong password or username")
	}


}