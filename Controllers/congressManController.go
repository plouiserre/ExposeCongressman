package Controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/plouiserre/exposecongressman/Models"
)

func CongressMans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	congressmans := models.AllCongressMans()

	json.NewEncoder(w).Encode(congressmans)
}

func CongressMan(w http.ResponseWriter, r *http.Request) {
	//TODO implement error 404 if no existing congressman with this ID
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	//TODO implement error 400
	if err != nil {
		log.Fatal(err)
	}

	congressman := models.GetCongressMan(id)

	json.NewEncoder(w).Encode(congressman)
}
