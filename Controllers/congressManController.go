package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error cast " + err.Error())
	} else {
		congressman := models.GetCongressMan(id)

		if congressman != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(congressman)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func CreateCongressMan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error Body " + err.Error())
	} else {
		var congressMan models.CongressMan

		errJson := json.Unmarshal(body, &congressMan)

		if errJson != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err.Error())
		}

		lid := models.InsertCongressMan(&congressMan)
		congressMan.Id = lid
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(congressMan)
	}
}

func UpdateCongressMan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error Body " + err.Error())
	} else {
		congressman := models.GetCongressMan(id)
		if congressman == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Println("No congressman find with this id " + vars["id"])
		} else {

			body, errBody := ioutil.ReadAll(r.Body)
			if errBody != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println(err.Error())
			} else {
				errJson := json.Unmarshal(body, &congressman)
				if errJson != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Println(err.Error())
				} else {
					models.UpdateCongressMan(congressman, id)
					w.WriteHeader(http.StatusOK)

					json.NewEncoder(w).Encode(congressman)
				}
			}
		}
	}
}

func DeleteCongressMan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error Body " + err.Error())
	} else {
		nbDelete := models.DeleteCongressMan(id)

		if nbDelete > 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
