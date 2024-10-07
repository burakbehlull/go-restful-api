package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
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


func CreateUser(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Veriler okunmadı", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "JSON'a çevrildi: ", http.StatusBadRequest)
		return
	}

	fmt.Printf("Alınan Kullanıcı Verisi: %+v\n", user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Kullanıcı yaratıldı."})
}