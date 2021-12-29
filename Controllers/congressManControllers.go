package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/plouiserre/exposecongressman/models"
)

func CongressMans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	congressmans := models.AllCongressMans()

	json.NewEncoder(w).Encode(congressmans)
}
