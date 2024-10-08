package main

import (
	"api/routers"
	"api/utils"
	"fmt"
	"net/http"
)

func main() {
	// routes
	routers.HomeRoutes()
	routers.UserRoutes()
	
	utils.Db()

	fmt.Println("Sunucu başlatıldı!")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Sunucu başlatılamadı: ", err)
	}
}
