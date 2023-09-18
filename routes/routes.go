package routes

import (
	"ecs_govel/app/http/controllers"

	"github.com/gorilla/mux"
)

var Router *mux.Router = new(mux.Router)

func init() {
	index(Router)

	// Ruta test
	Router.HandleFunc("url_sin_urlBase", controllers.NombreController.Insert).Methods("GET", "POST")
}
