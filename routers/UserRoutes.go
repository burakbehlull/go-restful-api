package routers

import (
	"api/controllers"
	"net/http"
)

func UserRoutes() {
	http.HandleFunc("/getUser", controllers.GetUser)
}
