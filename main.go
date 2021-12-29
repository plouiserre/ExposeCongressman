package main

import (
	"log"
	"net/http"
)

func main() {
	myRouter := InitializeRouter()

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
