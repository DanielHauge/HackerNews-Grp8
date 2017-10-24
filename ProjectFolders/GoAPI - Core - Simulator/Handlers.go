package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"github.com/streadway/amqp"
	"log"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!\n" +
					"\n" +
					"This api will accepts the following calls:\n" +
					"\n" +
					"PostStory\n" +
					"Route: /post\n" +
					"Method type: POST\n" +
					"Accepts Json Syntax: {'username': '<string>','post_type': '<string>',	'pwd_hash': '<string>','post_title': '<string>','post_url': '<string>", "post_parent': <int>, 'hanesst_id': <int>, 'post_text': '<string>'}\n" +
					"\n" +
					"GetLatest\n" +
					"Route: /post\n" +
					"Method type: GET\n" +
					"\n" +
					"GetStatust\n" +
					"Route: /status\n" +
					"Method type: GET\n" +
					"\n" +
						" ")
}

func GetLatest(w http.ResponseWriter, r *http.Request){

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin ,Accept, Content-Type, Content-Length, Accept-Encoding")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else { w.Header().Set("Access-Control-Allow-Origin", "*")}

	input := FindLatest()
	w.Write([]byte(input))
}

func PostStory(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin ,Accept, Content-Type, Content-Length, Accept-Encoding")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else { w.Header().Set("Access-Control-Allow-Origin", "*")}



		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
			log.Printf(err.Error())
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

	/// Get status of website, and other API and more

	if SqlStatus(){
		w.Write([]byte("Alive"))
	} else
	{
		w.Write([]byte("Down"))
	}


	// do some status things here


}