package routes

import (
	"github.com/gorilla/mux"
)

func index(loadRoute *mux.Router) {
	// loadRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./page/index.html") // Sirve el archivo HTML
	// })

	// Carga pagina Web Front las rutas se configuran en frontend.go
	// laod(loadRoute)
}
