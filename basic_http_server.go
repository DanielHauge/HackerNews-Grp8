package main

import (
	"fmt"
	"log"
	"net/http"
	"gopkg.in/mgo.v2" 
	//"gopkg.in/mgo.v2/bson" // will use for specific try
)

type Person struct {
	Name string
	Phone string
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	
	//Gets Session
	session, err := mgo.Dial("192.168.20.3:27017")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
			   &Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
			log.Fatal(err)
	}

	/// Queres persons of with these parameters.
	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 		log.Fatal(err)
	// }

	// fmt.Println("Phone:", result.Phone)

	
	var results []Person
	
	error := c.Find(nil).All(&results)
	if error != nil {
			// TODO: Do something about the error
	} else {
		fmt.Fprintln(w, results) 
	}

	






}

func main() {
	port := 8866

	http.HandleFunc("/", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
