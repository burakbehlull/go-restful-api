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
	client, err := utils.Db()
	
	if err != nil {
		fmt.Println("MongoDB bağlantı hatası:", err)
		defer client.Disconnect(nil)
	}

	port := os.Getenv("PORT")

	fmt.Println("Sunucu başlatıldı!")

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Sunucu başlatılamadı: ", err)
	}
}
