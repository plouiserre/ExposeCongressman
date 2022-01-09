package Controllers

import (
	"fmt"
	"net/http"
)

func Deputies(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Deputies method called")

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func Deputy(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Deputy method called")

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func CreateDeputy(w http.ResponseWriter, r *http.Request) {

	fmt.Println("CreateDeputy method called")

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateDeputy(w http.ResponseWriter, r *http.Request) {

	fmt.Println("UpdateDeputy method called")

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteDeputy(w http.ResponseWriter, r *http.Request) {

	fmt.Println("DeleteDeputy method called")

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
