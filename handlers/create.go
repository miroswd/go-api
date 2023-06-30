package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/miroswd/go-api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Json decode error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]string

	if err != nil {
		resp = map[string]string{
			"Error":   "true",
			"Message": fmt.Sprintf("Failed to insert a new todo\n%v", err),
		}
	} else {
		resp = map[string]string{
			"Error":   "false",
			"Message": fmt.Sprintf("Todo inserted successfully, id: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
