package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"github.com/streadway/amqp"
	"github.com/gorilla/mux"
)

func setheader (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin ,Accept, Content-Type, Content-Length, Accept-Encoding")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else { w.Header().Set("Access-Control-Allow-Origin", "*")}
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

	go func() {
		var request PostRequest

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &request); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

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

	go func() {
		var usr HNUser

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

	//vars := mux.Vars(r)
	//id := vars["storyid"]



	req := PostRequest{Username:"Hej",Post_type:"story",Post_parrent:-1,Post_text:"",Post_title:"HELLO!",Pwd_hash:"d421d",Hanesst_id:2003}
	/// Query data base and set req to correct data format.


	msgs, err := json.Marshal(req); if err != nil{ panic(err) }

	fmt.Fprint(w, msgs)
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


	vars := mux.Vars(r)
	id := vars["storyid"]
	fmt.Fprint(w, id)
}

