package routes

import "github.com/gorilla/mux"

func test(Router *mux.Router) {
	// Ruta test
	// Limita a ecuchar peticiones POST y GET
	//Router.HandleFunc("url_sin_urlBase", controllers.NombreController.Insert).Methods("GET", "POST")
	
	// Escucha cualquiera
	//Router.HandleFunc("url_sin_urlBase", controllers.NombreController.Insert)
}
