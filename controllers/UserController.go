package controllers

import (
	"api/models"
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	data := models.User{
		Id:       1,
		Name:     "burak",
		Username: "burakbehlül",
		Age:      24,
		IsAdmin:  true,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
