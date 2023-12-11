package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Se configuran las routas que serviran los archivos del fronted
func frontend(loadRoute *mux.Router) {
	//cargar styles
	_style(loadRoute)
	//carregar js
	_js(loadRoute)
	// carregar pages
	_pages(loadRoute)
}

func _pages(loadRoute *mux.Router) {
	// cargar index
	loadRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/index.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/login/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/login/login.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/recoveryPass/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/recoveryPass/recoveryPass.html") // Sirve el archivo HTML
	})

	loadRoute.HandleFunc("/main/", func(w http.ResponseWriter, r *http.Request) {
		login := true
		if login == true {
			http.ServeFile(w, r, "../frontend/main/principal.html") // Sirve el archivo HTML
		} else {

		}
	})
	// Curso
	loadRoute.HandleFunc("/main/curso/registrar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/curso/registrar.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/curso/mostrar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/curso/mostrar.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/curso/eliminar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/curso/eliminar.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/curso/modificar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/curso/modificar.html") // Sirve el archivo HTML
	})
	// Matricula
	loadRoute.HandleFunc("/main/matricula/mostrarOfCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/matricula/mostrarOfCurso.html") // Sirve el archivo HTML
	})

	// estudiante
	loadRoute.HandleFunc("/main/estudiante/asociarToCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/estudiante/asociarToCurso.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/estudiante/eliminar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/estudiante/eliminar.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/estudiante/eliminarOfCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/estudiante/eliminarOfCurso.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/estudiante/modificar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/estudiante/modificar.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/estudiante/modificarOfCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/estudiante/modificarOfCurso.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/estudiante/mostrarAllCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/estudiante/mostrarAllCurso.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/estudiante/mostrarJuridico/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/estudiante/mostrarJuridico.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/estudiante/registrar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/estudiante/registrar.html") // Sirve el archivo HTML
	})

	// Instructor
	loadRoute.HandleFunc("/main/instructor/agregarToCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/instructor/agregarToCurso.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/instructor/eliminar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/instructor/eliminar.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/instructor/eliminarOfCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/instructor/eliminarOfCurso.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/instructor/modificar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/instructor/modificar.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/instructor/modificarCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/instructor/modificarCurso.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/instructor/modificarInCurso/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/instructor/modificarInCurso.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/instructor/registrar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/instructor/registro.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/instructor/mostrar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/instructor/mostrar.html") // Sirve el archivo HTML
	})

	// Usuario
	loadRoute.HandleFunc("/main/usuario/cambiarSenha/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/usuario/cambiarSenha.html") // Sirve el archivo HTML
	})
	loadRoute.HandleFunc("/main/usuario/recuperarSenha/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/main/usuario/recuperarSenha.html") // Sirve el archivo HTML
	})
}

func _style(loadRoute *mux.Router) {
	// Style login
	loadRoute.PathPrefix("/main/login/estilos/").Handler(
		http.StripPrefix("/main/login/estilos/", http.FileServer(http.Dir("../frontend/main/login/estilos/"))))
	loadRoute.PathPrefix("/main/login/images/").Handler(
		http.StripPrefix("/main/login/images/", http.FileServer(http.Dir("../frontend/main/login/images/"))))
	// Cargar imagenes carpeat /main/img/
	loadRoute.PathPrefix("/main/img/").Handler(
		http.StripPrefix("/main/img/", http.FileServer(http.Dir("../frontend/main/img/"))))
	//Cargar stylos de carpeta /main/css/
	loadRoute.PathPrefix("/main/css/").Handler(
		http.StripPrefix("/main/css/", http.FileServer(http.Dir("../frontend/main/css/"))))
	// //Cargar stylos de carpeta /main/
	// loadRoute.PathPrefix("/main/styles.css").Handler(
	// 	http.StripPrefix("/main/styles.css", http.FileServer(http.Dir("../frontend/main/styles.css"))))
}

func _js(loadRoute *mux.Router) {
	// Js login
	loadRoute.PathPrefix("/main/login/js/").Handler(
		http.StripPrefix("/main/login/js/", http.FileServer(http.Dir("../frontend/main/login/js/"))))
	// Js libs
	loadRoute.PathPrefix("/main/js/").Handler(
		http.StripPrefix("/main/js/", http.FileServer(http.Dir("../frontend/main/js/"))))
}
