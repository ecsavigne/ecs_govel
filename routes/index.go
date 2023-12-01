package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func index(loadRoute *mux.Router) {
	// Cargando desde un html
	// loadRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./page/index.html") // Sirve el archivo HTML
	// })
	loadRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Welcome to ecs_govel</h1>"))
	})
}
