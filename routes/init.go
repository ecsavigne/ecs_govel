/*
Fichero de inicialiciacion del paquete routes
*/

package routes

import (
	"github.com/gorilla/mux"
)

var Router *mux.Router = new(mux.Router)

func init() {
	// Inicializ el grupo de routas que se desean
	//index(Router)
	services(Router)
	frontend(Router)
}
