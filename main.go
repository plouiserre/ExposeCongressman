package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Greeting struct {
	Language string `json:"Langage"`
	Message  string `json:"Greeting"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!!!!!")
	fmt.Println("Endpoint Hit: homepage")
}

func sayHelloWorldFrench(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: greetings Page")
	Greetings := []Greeting{
		{Language: "French", Message: "Salut tout le monde!!!"},
		{Language: "Hello", Message: "Hello every body!!!!"},
	}
	json.NewEncoder(w).Encode(Greetings)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/greetings", sayHelloWorldFrench)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
