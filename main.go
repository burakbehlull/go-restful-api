package main

import (
	"api/routers"
	"api/utils"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// routes
	routers.HomeRoutes()
	routers.UserRoutes()
	
	utils.LoadDotenv()
	utils.Db()

	port := os.Getenv("PORT")

	fmt.Println("Sunucu başlatıldı!")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Sunucu başlatılamadı: ", err)
	}
}
