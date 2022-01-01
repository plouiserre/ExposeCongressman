package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/plouiserre/exposecongressman/Models"
)

//TODO factoriser les parties communes entre chaque méthode
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

func CreateCongressMan(w http.ResponseWriter, r *http.Request) {
	//TODO factoriser cette partie
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	//TODO implemnter error 400
	if err != nil {
		log.Fatal(err)
	}

	var congressMan models.CongressMan

	errJson := json.Unmarshal(body, &congressMan)

	//TODO implémenter erreur 500
	if errJson != nil {
		log.Fatal(err)
	}

	lid := models.InsertCongressMan(&congressMan)
	congressMan.Id = lid
	json.NewEncoder(w).Encode(congressMan)
}

func UpdateCongressMan(w http.ResponseWriter, r *http.Request) {
	//TODO implement error 404 if no existing congressman with this ID
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	//TODO implement error 400
	if err != nil {
		log.Fatal(err)
	}

	body, errBody := ioutil.ReadAll(r.Body)
	//TODO implemnter error 400
	if errBody != nil {
		log.Fatal(errBody)
	}

	congressman := models.GetCongressMan(id)
	errJson := json.Unmarshal(body, &congressman)
	//TODO implemnter error 400
	if errBody != nil {
		log.Fatal(errJson)
	}

	models.UpdateCongressMan(congressman)

	json.NewEncoder(w).Encode(congressman)
}

func DeleteCongressMan(w http.ResponseWriter, r *http.Request) {
	//TODO implement error 404 if no existing congressman with this ID
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	models.DeleteCongressMan(id)
}
