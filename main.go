package main

import (
	"api/routers"
	"fmt"
	"net/http"
)

func main() {
	// routes
	routers.UserRoutes()

	fmt.Println("Sunucu başlatıldı!")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Sunucu başlatılamadı: ", err)
	}
}
