package routes

import (
	"github.com/gorilla/mux"

	"ecs_govel/app/http/controllers"
)

var Router *mux.Router = new(mux.Router)

func init() {
	// Ruta test
	Router.HandleFunc("url_sin_urlBase", controllers.NombreController.Insert).Methods("GET", "POST")
}
