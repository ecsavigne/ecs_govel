package routes

import (
	"ecs_govel/app/http/controllers"

	"github.com/gorilla/mux"
)

func services(loadRoute *mux.Router) {
	//Curso
	loadRoute.HandleFunc("/registrarCurso", new(controllers.CursoController).RegistrarCurso)
	loadRoute.HandleFunc("/eliminarCurso", new(controllers.CursoController).EliminarCurso)
	loadRoute.HandleFunc("/modificarCurso", new(controllers.CursoController).ModificarCurso)
	loadRoute.HandleFunc("/MostrarCursoById", new(controllers.CursoController).MostrarCursoById)
	loadRoute.HandleFunc("/mostrarCursos", new(controllers.CursoController).MostrarCursos)
	loadRoute.HandleFunc("/MostrarCursosContenido", new(controllers.CursoController).MostrarCursosContenido)

	//Estudiante
	loadRoute.HandleFunc("/registrarEstudiante", new(controllers.EstudianteController).RegistrarEstudiante)
	loadRoute.HandleFunc("/registrarEstudianteEnCurso", new(controllers.EstudianteController).RegistrarEstudianteEnCurso)
	loadRoute.HandleFunc("/eliminarEstudiante", new(controllers.EstudianteController).EliminarEstudiante)
	loadRoute.HandleFunc("/eliminarEstudianteCurso", new(controllers.EstudianteController).ElminarEstudianteCurso)
	loadRoute.HandleFunc("/modificarEstudiante", new(controllers.EstudianteController).ModificarEstudiante)
	loadRoute.HandleFunc("/modificarEstudianteCurso", new(controllers.EstudianteController).ModificarDatosEstudianteCurso)
	loadRoute.HandleFunc("/mostrarEstudianteDeCurso", new(controllers.EstudianteController).MostrarEstudianteDeCurso)
	loadRoute.HandleFunc("/mostrarEstudiante", new(controllers.EstudianteController).MostrarEstudiante)
	loadRoute.HandleFunc("/mostrarEstudianteJuridicoNoJuridico", new(controllers.EstudianteController).MostrarEstudianteJuridicoNoJuridico)

	//Instructor
	loadRoute.HandleFunc("/registrarInstructor", new(controllers.InstructorController).RegistrarInstructor)
	loadRoute.HandleFunc("/agregarInstructorToCurso", new(controllers.InstructorController).AgregarCursoToInstructor)
	loadRoute.HandleFunc("/eliminarInstructor", new(controllers.InstructorController).EliminarInstructor)
	loadRoute.HandleFunc("/eliminarInstructorDeCurso", new(controllers.InstructorController).EliminarCursoOffInstructor)
	loadRoute.HandleFunc("/modificarInstructor", new(controllers.InstructorController).MostrarInstructor)
	loadRoute.HandleFunc("/modificarInstructorCurso", new(controllers.InstructorController).ModificarCursoOffInstructor)
	loadRoute.HandleFunc("/mostrarInstructor", new(controllers.InstructorController).MostrarInstructor)
	loadRoute.HandleFunc("/mostrarInstructores", new(controllers.InstructorController).MostrarInstructores)
	loadRoute.HandleFunc("/mostrarCursosOfInstructor", new(controllers.InstructorController).MostrarCursosOfInstructor)

	//Gestion de Usuario
	loadRoute.HandleFunc("/registrarUsr", new(controllers.UsuarioController).RegistrarUsr)
	loadRoute.HandleFunc("/cambioPass", new(controllers.UsuarioController).CambiarPass)
	loadRoute.HandleFunc("/recuperarPass", new(controllers.UsuarioController).RecuperarPass)
	loadRoute.HandleFunc("/verificarUsr", new(controllers.UsuarioController).VerificarUsr)

	//matricula
	loadRoute.HandleFunc("/mostrarMatricula", new(controllers.MatriculaController).MostrarMatriculaDeCurso)

	//Contenido
	loadRoute.HandleFunc("/registrarContenido", new(controllers.ContenidoController).RegistrarContenido)
	loadRoute.HandleFunc("/asociarContenidoToEstructura", new(controllers.ContenidoController).AsociarContenidoToEstructura)
	loadRoute.HandleFunc("/eliminarContenido", new(controllers.ContenidoController).EliminarContenido)
	loadRoute.HandleFunc("/modificarContenidoToEstructura", new(controllers.ContenidoController).ModificarContenidoToEstructura)
	loadRoute.HandleFunc("/modificarContenido", new(controllers.ContenidoController).ModificarContenido)
	loadRoute.HandleFunc("/mostrarContenido", new(controllers.ContenidoController).MostrarContenido)
	loadRoute.HandleFunc("/mostrarAllContenido", new(controllers.ContenidoController).MostrarAllContenido)

	//Estructura
	loadRoute.HandleFunc("/registrarEstructura", new(controllers.EstructuraController).RegistrarEstructura)
	loadRoute.HandleFunc("/eliminarEstructura", new(controllers.EstructuraController).EliminarEstructura)
	loadRoute.HandleFunc("/mostrarEstructura", new(controllers.EstructuraController).MostrarEstructura)
	loadRoute.HandleFunc("/modificarEstructura", new(controllers.EstructuraController).ModificarEstructura)
}
