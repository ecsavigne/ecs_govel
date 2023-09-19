package routes

import (
	"github.com/gorilla/mux"
)

// Se configuran las routas que serviran los archivos del fronted
func frontend(loadRoute *mux.Router) {
	// loadRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../frontend/index.html") // Sirve el archivo HTML
	// })

	// loadRoute.HandleFunc("/main/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../frontend/main/principal.html") // Sirve el archivo HTML
	// })
	// loadRoute.HandleFunc("/main/update/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../frontend/main/principal.html") // Sirve el archivo HTML
	// })
	// loadRoute.HandleFunc("/main/create/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../frontend/main/create/create.html") // Sirve el archivo HTML
	// })
	// loadRoute.HandleFunc("/main/delete/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../frontend/main/delete/delete.ecs") // Sirve el archivo HTML
	// })
}
