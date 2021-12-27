package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Greeting string `json:"Greeting"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!!!!!")
	fmt.Println("Endpoint Hit: homepage")
}

func sayHelloWorldFrench(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: french Page")
	hiFrench := Message{
		Greeting: "Salut tout le monde!!!",
	}
	json.NewEncoder(w).Encode(hiFrench)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/french", sayHelloWorldFrench)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
