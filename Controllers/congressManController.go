package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
)

func CongressMans(w http.ResponseWriter, r *http.Request) {
	repo := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	congressmans, noError := repo.AllCongressMans()

	if noError {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(congressmans)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func CongressMan(w http.ResponseWriter, r *http.Request) {
	repo := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error cast " + err.Error())
	} else {
		congressman, noError := repo.GetCongressMan(id)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if congressman != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(congressman)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func CreateCongressMan(w http.ResponseWriter, r *http.Request) {
	repo := InitCongressmanRepository()
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

		lid, noError := repo.InsertCongressMan(&congressMan)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			congressMan.Id = lid
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(congressMan)
		}
	}
}

func UpdateCongressMan(w http.ResponseWriter, r *http.Request) {
	repo := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error Body " + err.Error())
	} else {
		congressman, noError := repo.GetCongressMan(id)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if congressman == nil {
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
					noError := repo.UpdateCongressMan(congressman, id)
					if !noError {
						w.WriteHeader(http.StatusInternalServerError)
					} else {
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(congressman)
					}
				}
			}
		}
	}
}

func DeleteCongressMan(w http.ResponseWriter, r *http.Request) {
	repo := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error Body " + err.Error())
	} else {
		nbDelete, noError := repo.DeleteCongressMan(id)

		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if nbDelete > 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func InitCongressmanRepository() repository.CongressmanRepository {
	congressManRepository := repository.CongressmanRepository{}
	return congressManRepository
}
