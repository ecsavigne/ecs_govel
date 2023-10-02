package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Se configuran las routas que serviran los archivos del fronted
func frontend(loadRoute *mux.Router) {
	// cargar index
	loadRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/index.html") // Sirve el archivo HTML
	})

	loadRoute.HandleFunc("/main/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/principal.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/update/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/update/update.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/create/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/create/curso.ecs") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/delete/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/delete/delete.ecs") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/show/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/show/show.ecs") // Sirve el archivo HTML
	})
}
