package routes

import (
	"ecs_govel/app/http/controllers"

	"github.com/gorilla/mux"
)

func services(loadRoute *mux.Router) {
	//Curso
	loadRoute.HandleFunc("/registrarCurso", new(controllers.CursoController).RegistrarCurso)
	loadRoute.HandleFunc("/modificarCurso", new(controllers.CursoController).ModificarCurso)
	loadRoute.HandleFunc("/eliminarCurso", new(controllers.CursoController).EliminarCurso)
	loadRoute.HandleFunc("/mostrarCurso", new(controllers.CursoController).MostrarCurso)

	//Estudiante

	//Instructor

	//Gestion de Usuario
}
