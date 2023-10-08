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
	loadRoute.HandleFunc("/registrarEstudiante", new(controllers.EstudianteController).RegistrarEstudiante)
	loadRoute.HandleFunc("/registrarEstudianteEnCurso", new(controllers.EstudianteController).RegistrarEstudianteEnCurso)
	loadRoute.HandleFunc("/modificarEstudiante", new(controllers.EstudianteController).ModificarEstudiante)
	loadRoute.HandleFunc("/modificarEstudianteCurso", new(controllers.EstudianteController).ModificarDatosEstudianteCurso)
	loadRoute.HandleFunc("/eliminarEstudiante", new(controllers.EstudianteController).EliminarEstudiante)
	loadRoute.HandleFunc("/eliminarEstudianteCurso", new(controllers.EstudianteController).ElminarEstudianteCurso)
	loadRoute.HandleFunc("/mostrarEstudianteDeCurso", new(controllers.EstudianteController).MostrarEstudianteDeCurso)
	loadRoute.HandleFunc("/mostrarEstudianteJuridicoNoJuridico", new(controllers.EstudianteController).MostrarEstudianteJuridicoNoJuridico)

	//Instructor
	loadRoute.HandleFunc("/registrarInstructor", new(controllers.InstructorController).RegistrarInstructor)
	loadRoute.HandleFunc("/modificarInstructor", new(controllers.InstructorController).MostrarInstructor)
	loadRoute.HandleFunc("/eliminarInstructor", new(controllers.InstructorController).EliminarInstructor)
	loadRoute.HandleFunc("/mostrarInstructor", new(controllers.InstructorController).MostrarInstructor)
	loadRoute.HandleFunc("/agregarInstructorToCurso", new(controllers.InstructorController).AgregarCursoToInstructor)
	loadRoute.HandleFunc("/eliminarInstructorDeCurso", new(controllers.InstructorController).EliminarCursoOffInstructor)
	loadRoute.HandleFunc("/modificarInstructorCurso", new(controllers.InstructorController).ModificarCursoOffInstructor)

	//Gestion de Usuario
	loadRoute.HandleFunc("/registrarUsr", new(controllers.UsuarioController).RegistrarUsr)
	loadRoute.HandleFunc("/recuperarPass", new(controllers.UsuarioController).RecuperarPass)
	loadRoute.HandleFunc("/cambioPass", new(controllers.UsuarioController).CambiarPass)

	//matricula
	loadRoute.HandleFunc("/mostrarMatricula", new(controllers.MatriculaController).MostrarMatriculaDeCurso)

	//Contenido
	loadRoute.HandleFunc("/registrarContenido", new(controllers.ContenidoController).RegistrarContenido)
	loadRoute.HandleFunc("/mostrarContenido", new(controllers.ContenidoController).MostrarContenido)
	loadRoute.HandleFunc("/eliminarContenido", new(controllers.ContenidoController).EliminarContenido)
	loadRoute.HandleFunc("/modificarContenido", new(controllers.ContenidoController).ModificarContenido)
	loadRoute.HandleFunc("/asociarContenidoToEstructura", new(controllers.ContenidoController).AsociarContenidoToEstructura)
	loadRoute.HandleFunc("/modificarContenidoToEstructura", new(controllers.ContenidoController).ModificarContenidoToEstructura)
	loadRoute.HandleFunc("/mostrarAllContenido", new(controllers.ContenidoController).MostrarAllContenido)

	//Estructura
	loadRoute.HandleFunc("/registrarEstructura", new(controllers.EstructuraController).RegistrarEstructura)
	loadRoute.HandleFunc("/mostrarEstructura", new(controllers.EstructuraController).MostrarEstructura)
	loadRoute.HandleFunc("/eliminarEstructura", new(controllers.EstructuraController).EliminarEstructura)
	loadRoute.HandleFunc("/modificarEstructura", new(controllers.EstructuraController).ModificarEstructura)
}
