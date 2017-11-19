package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"gopkg.in/russross/blackfriday.v2"
	"log"
	"strconv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		"\n\n- Return of Example: \n\n > ```Publishing to RQ for DB Insertion```"+
		"\n\n" +
		"## GetLatest\n" +
		"This call will return the latest ingested story or comment which was sent by the simulator program. \n\n"+
		"- Route: /latest\n" +
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
		"\n\n- Return of Example: \n\n > ```Publishing to RQ for DB Insertion``` with 200 OK"+
		"\n\n > ```Username or email has been taken``` With 406"+
		"\n\n" +
		"## Login\n" +
		"This call is to verify users, it will see if username and password is correct and respond accordingly. \n\n"+
		"- Route: /login\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: \n\n> ```{\"username\": \"<string>\",\"password\": \"<string>\"}```\n\n" +
		"- Example: \n\n> ```{\"username\": \"farmer\",\"password\": \"xMBVi4fAO5\"}```\n"+
		"\n\n- Return of Example: \n\n > ```{\"username\":\"farmer\",\"email_addr\":\"\"none\"\",\"karma\":9}```\nwith 200 OK, 406 if bad"+
		"\n\n- Return of another exampme: \n\n > ```{\"username\":\"Daniel\",\"email_addr\":\"Animcuil@gmail.com\",\"karma\":0}```\nwith 200 OK, 406 if bad"+
		"\n\n" +
		"\n\n" +
		"## GetLatestStories\n" +
		"This call will return an array of stories. dex:0 dex_to 100 will return last 100 threads, dex:100,dex_to:200 will return the 100 threads before the very newest 100 threads \n\n"+
		"- Route: /stories\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: " +
		"\n\n>``` {\"dex\": <int>,\"dex_to\": <int>} ```\n\n" +
		"- Example: \n\n > ```{\"dex\": 0,\"dex_to\": 100}```"+
		"\n\n- Return for dex:0-dex_to:2 = \n\n > ``` {\"stories\":[{\"id\":446,\"title\":\"How Important is the .com TLD?\",\"username\":\"python_kiss\",\"points\":0,\"time\":\"15 Minutes Ago\",\"url\":\"http://www.netbusinessblog.com/2007/02/19/how-important-is-the-dot-com/\",\"commentamount\":0}]} ```"+
		"\n\n" +
		"## GetStoryByID\n" +
		"This call is not implementet yet, but it works and will return a static story. \n\n"+
		"- Route: /stories/{storyid}\n" +
		"- Method type: GET\n" +
		"\n\n- Return for 174 = \n\n > ``` {\"thread\":{\"id\":173,\"title\":\"_\",\"username\":\"akkartik\",\"points\":0,\"time\":\"24 Minutes Ago\",\"url\":\"http://alwayson.goingon.com/permalink/post/9894\",\"commentamount\":0},\"comments\":[{\"id\":76,\"comment\":\"I've gotten used to using submit to find the reddit discussion for a page. Turns out there's no duplication-detection here on news.yc yet. Apologies.\\n\\nNice that they allow editing the title.\",\"username\":\"akkartik\",\"points\":0,\"time\":\"24 Minutes Ago\"}]} ```"+
		"\n\n" +
		"## GetComments\n" +
		"This call is to create users, it will run concurrently, but is not guarenteed to be succesfully insertet. But if it can be succesfull insertet, it will at some point. \n\n"+
		"- Route: /comments/{storyid}\n" +
		"- Method type: GET\n" +
		"\n\n- Return of Example: \n\n > ```Same as above, but without the story aswell, just an array of comments```"+
		"\n\n" +
		"## RecoverPassword\n" +
		"This call will send an email to the address which is linked with the username provided. \n\n"+
		"- Route: /recover\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: " +
		"\n\n>``` {\"username\": \"<string>\"} ```\n\n" +
		"- Example: \n\n > ```{\"username\": \"farmer\"}```"+
		"\n\n- Return of Example: - Note: For security reasons: it will allways send this regardless if it succeds or not \n\n > ```Am Email has been sent to the specified email linked with the username provided```"+
		"\n\n"+
		"## UpdatePassword\n" +
		"This call will change the password of the user. \n\n"+
		"- Route: /update\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: " +
		"\n\n>``` {\"username\": \"<string>\", \"password\": \"<string>\", \"new_password\": \"<string>\"} ```\n\n" +
		"- Example: \n\n > ```{\"username\": \"farmer\", \"password\": \"currentpassword\", \"new_password\": \"newpassword\"}```"+
		"\n\n- Return of Example: \n\n > ```Succesfully Changed Password with 200 OK and 406 if not correct credentials.```"+
		"\n\n"+
		"## Upvote\n" +
		"This call will register the upvote. if comment set threadid to -1 \n\n"+
		"- Route: /upvote\n" +
		"- Method type: POST\n" +
		"- Accepts Json & text/plain Syntax: " +
		"\n\n>``` {\"thread_id\": -1,\"comment_id\": 4, \"username\": \"farmer\"} ```\n\n" +
		"- Example: \n\n > ```{\"thread_id\": -1,\"comment_id\": 4, \"username\": \"farmer\"}```"+
		"\n\n- Return of Example: \n\n > ```200 OK```"+
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
	promRequests.Inc()
	input := FindLatest()
	w.Write([]byte(input))
} /// Not used by Website

func PostStory(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	setheader(w, r)
	promRequests.Inc()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		Error.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		Error.Println(err)
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
} /// Used by website

func GetStatus(w http.ResponseWriter, r *http.Request){
	promRequests.Inc()
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Alive")
} /// Not used by website

func CreateUser(w http.ResponseWriter, r *http.Request){
	promRequests.Inc()
	setheader(w, r)


	var usr HNUser
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
			Error.Println(err)
			}
	if err := r.Body.Close(); err != nil {
			Error.Println(err)
	}
	if err := json.Unmarshal(body, &usr); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			Error.Println(err)
		}
	}
	if CheckIfTaken(usr){
		go func() {
			/// Implement MySQL
			props := amqp.Publishing{
				ContentType: "application/json; charset=UTF-8",
				Body:        body,
			}
			SendToRabbit(props, User_Q.Name)

		}()
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Publishing to RQ for DB Insertion")
	}else {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, "Username or email has been taken")
	}


} /// Used by website

func Login(w http.ResponseWriter, r *http.Request){
	promRequests.Inc()
	setheader(w, r)

	var usr UserLogin
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		Error.Println(err)
	}
	if err := r.Body.Close(); err != nil {
		Error.Println(err)
	}
	if err := json.Unmarshal(body, &usr); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			Error.Println(err)
		}
	}
	Correct, user := VerifyUser(usr)

	if  Correct{
		w.WriteHeader(http.StatusOK)

		msgs, err := json.Marshal(user); if err != nil{ Error.Println(err) }

		fmt.Fprint(w, string(msgs))
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, "Error, incorrect username and/or password")
	}


} /// Used by website

func GetStoryByID(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	promRequests.Inc()
	vars := mux.Vars(r)
	id := vars["storyid"]

	realid, err := strconv.Atoi(id); if err != nil {log.Println("Error in parsing int"); fmt.Fprint(w, "Error in parsing int, you should only use ints has ID")}
	log.Print(realid)

	story := GetSingleStory(realid)
	log.Print(story.Username)
	/// Query data base and set req to correct data format.

	comments := QueryAllComments(realid)
	final := StoryWithComments{story, comments}
	msgs, err := json.Marshal(final); if err != nil{  Error.Println(err)
	}

	fmt.Fprint(w, string(msgs))
} /// Used by website

func GetLatestStories(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	promRequests.Inc()
	var AllStories LatestStories
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		Error.Println(err)
	}
	if err := r.Body.Close(); err != nil {
		Error.Println(err)
	}
	var req StoryRequest
	if err := json.Unmarshal(body, &req); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(err); err != nil {
			Error.Println(err)
		}
	}
	AllStories = QueryLatestStories(req.Dex,req.DexTo)

	msgs, err := json.Marshal(AllStories); if err != nil{ Error.Println(err) }
	fmt.Fprint(w, string(msgs))
} /// Used by website

func GetComments(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	promRequests.Inc()
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["storyid"]

	threadid, err := strconv.Atoi(id); if err != nil {log.Println("Error in parsing int"); fmt.Fprint(w, "Error in parsing int, you should only use ints has ID")}
	comments := QueryAllComments(threadid)
	msgs, err := json.Marshal(comments); if err != nil{ Error.Println(err) }

	fmt.Fprint(w, string(msgs))


} /// Not used currently by website

func StartRecovery(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	promRequests.Inc()
	var recinfo RecoveryData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		Error.Println(err)
	}
	if err := r.Body.Close(); err != nil {
		Error.Println(err)
	}
	if err := json.Unmarshal(body, &recinfo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			Error.Println(err)
		}
	}
	go func() {
	pwd, email := GetRecoveryInformation(recinfo.Username)
	log.Printf(recinfo.Username)
	log.Printf(email)

		SendEmail(email, pwd)

	}()
	fmt.Fprint(w, "An Email has been sent to the specified email linked with the username provided")


} /// Used by website

func UpdatePassword(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	promRequests.Inc()
	var change PasswordChangeData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		Error.Println(err)
	}
	if err := r.Body.Close(); err != nil {
		Error.Println(err)
	}
	if err := json.Unmarshal(body, &change); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			Error.Println(err)
		}
	}
	Correct, _ := VerifyUser(UserLogin{change.Username,change.Password})
	if Correct{
		err := ChangePassword(change.NewPassword, GetUserID(change.Username))
		if err != nil{
			fmt.Fprint(w, "There was an error which might have prevented the password to change, password"+err.Error())
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Succesfully Changed Password")
	}else {
		fmt.Fprint(w, "Error in credentials, wrong password or username")
		w.WriteHeader(http.StatusNotAcceptable)
	}


} /// Used by website

func Upvote(w http.ResponseWriter, r *http.Request){
	setheader(w, r)
	promRequests.Inc()
	w.WriteHeader(http.StatusOK)

	var op UpvoteData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		RabbitMessage(Log_Q.Name, "Error in Opening the bodyreader in upvote")
		Error.Println(err)
	}
	if err := r.Body.Close(); err != nil {
		RabbitMessage(Log_Q.Name, "Error in Closing the bodyreader in upvote")
		Error.Println(err)

	}
	if err := json.Unmarshal(body, &op); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		RabbitMessage(Log_Q.Name, "Error in parsing Upvote Message to json")
		if err := json.NewEncoder(w).Encode(err); err != nil {

			Error.Println(err)
		}
	}

	go UpdateUpvote(op)


}

func GetMetrics (w http.ResponseWriter, r *http.Request){
	promhttp.Handler().ServeHTTP(w, r)
}