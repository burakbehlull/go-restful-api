package routers

import (
	"api/controllers"
	"net/http"
)

func HomeRoutes() {
	http.HandleFunc("/", controllers.Home)
}
