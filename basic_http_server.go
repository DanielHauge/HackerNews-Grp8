package main

import (
	"fmt"
	"log"
	"net/http"
	//"gopkg.in/mgo.v2" 
	//"gopkg.in/mgo.v2/bson" // will use for specific try
)

type Person struct {
	Name string
	Phone string
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	
		fmt.Fprintln(w, "Hello Static website - Og lidt mere - og endnu mere, og nu med ekstra mere") 
	

	






}

func main() {
	port := 8080

	http.HandleFunc("/", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
