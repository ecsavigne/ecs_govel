package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func index(t *mux.Router) {
	t.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./page/index.html") // Sirve el archivo HTML
	})
}
