package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Greetings = []Greeting{
	{Language: "French", Message: "Salut tout le monde!!!"},
	{Language: "English", Message: "Hello every body!!!!"},
}

type Greeting struct {
	Language string `json:"Langage"`
	Message  string `json:"Greeting"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!!!!!")
	fmt.Println("Endpoint Hit: homepage")
}

func sayGreetings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: greetings Page")

	json.NewEncoder(w).Encode(Greetings)
}

func greetingSpecificLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Language"]

	for _, greeting := range Greetings {
		if greeting.Language == key {
			json.NewEncoder(w).Encode(greeting)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/greetings", sayGreetings)
	myRouter.HandleFunc("/greeting/{Language}", greetingSpecificLanguage)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
