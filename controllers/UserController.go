package controllers

import (
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	data := nil
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}