package controllers

import (
	"ecs_govel/app/helpers"
	"ecs_govel/app/models"
	"ecs_govel/app/repositories"
	"encoding/json"
	"net/http"
)

type EstudianteController struct {
	estudianteRepository repositories.EstudianteRepository
}

func (c *EstudianteController) RegistrarEstudianteEnCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	fecha_ingreso, err := helpers.ValidateFecha(r.FormValue("fecha_ingreso"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()
	if err != nil {
		panic("Error formateando fecha en [Controller-Estudiante.RegistrarEstudianteEnCurso] - Error: " + err.Error() + " Fecha in :" + r.FormValue("fecha_ingreso"))
	}

	matricula := models.Matricula{
		CursoId:      idCurso,
		EstudianteCi: ci,
		FechaIngreso: fecha_ingreso,
	}
	c.estudianteRepository.RegistrarEstudianteEnCurso(&matricula)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci":      ci,
			"idCurso": idCurso,
		},
		"func": "RegistrarEstudianteEnCurso",
	})
}

func (c *EstudianteController) RegistrarEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	nome := r.FormValue("nome")
	sobreNome := r.FormValue("sobreNome")
	enderecao := r.FormValue("enderecao")
	correio, _ := helpers.ValidateCorreio(r.FormValue("correio"))
	siJuridico, _ := helpers.ValidateBool(r.FormValue("siJuridico"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()

	estud := models.Persona{
		CI:        ci,
		Nombre:    nome,
		Apellidos: sobreNome,
		Dir:       enderecao,
		Mail:      correio,
	}
	c.estudianteRepository.RegistrarEstudiante(&estud, siJuridico)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci":         ci,
			"nome":       nome,
			"sobreNome":  sobreNome,
			"enderecao":  enderecao,
			"correio":    correio,
			"siJuridico": siJuridico,
		},
		"func": "RegistrarEstudiante",
	})
}

func (c *EstudianteController) ModificarEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	nome := r.FormValue("nome")
	sobreNome := r.FormValue("sobreNome")
	enderecao := r.FormValue("enderecao")
	correio, _ := helpers.ValidateCorreio(r.FormValue("idCurso"))
	siJuridico, _ := helpers.ValidateBool(r.FormValue("siJuridico"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()

	estud := models.Persona{
		CI:        ci,
		Nombre:    nome,
		Apellidos: sobreNome,
		Dir:       enderecao,
		Mail:      correio,
	}
	c.estudianteRepository.ModificarEstudiante(&estud, siJuridico)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci":         ci,
			"nome":       nome,
			"sobreNome":  sobreNome,
			"enderecao":  enderecao,
			"correio":    correio,
			"siJuridico": siJuridico,
		},
		"func": "ModificarEstudiante",
	})
}

func (c *EstudianteController) ModificarDatosEstudianteCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	idCursoA, _ := helpers.ValidateInt(r.FormValue("idCursoA"))
	idCursoN, _ := helpers.ValidateInt(r.FormValue("idCursoN"))
	fecha_ingreso, _ := helpers.ValidateFecha(r.FormValue("fechaIngreso"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()

	matric := models.Matricula{
		CursoId:      idCursoN,
		EstudianteCi: ci,
		FechaIngreso: fecha_ingreso,
	}
	if idCursoA != idCursoN {
		c.estudianteRepository.ModificarEstudianteCurso(&matric, idCursoA)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso": idCursoN,
			"ci":      ci,
		},
		"func": "ModificarDatosEstudianteCurso",
	})
}

func (c *EstudianteController) EliminarEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()
	estudiante := models.Persona{
		CI: ci,
	}
	c.estudianteRepository.EliminarEstudiante(&estudiante)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci": ci,
		},
		"func": "EliminarEstudiante",
	})
}

func (c *EstudianteController) ElminarEstudianteCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()
	matricula := models.Matricula{
		EstudianteCi: ci,
		CursoId:      idCurso,
	}
	c.estudianteRepository.EliminarEstudianteCurso(&matricula)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci":      ci,
			"idCurso": idCurso,
		},
		"func": "ElminarEstudianteCurso",
	})
}

func (c *EstudianteController) MostrarEstudianteDeCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()
	res := c.estudianteRepository.MostrarEstudianteDeCurso(idCurso)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso":     idCurso,
			"estudiantes": res,
		},
		"estudiantes": res,
		"func":        "MostrarEstudianteDeCurso1",
	})
}

func (c *EstudianteController) MostrarEstudianteJuridicoNoJuridico(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	siJuridico, _ := helpers.ValidateBool(r.FormValue("siJuridico"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()

	estudiante := models.Estudiante{
		SiJuridico: siJuridico,
	}
	res := c.estudianteRepository.MostrarEstudianteJuridicoNoJuridico(&estudiante)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"siJuridico": siJuridico,
		},
		"func":        "MostrarEstudianteJuridicoNoJuridico",
		"estudiantes": res,
	})
}

func (c *EstudianteController) MostrarEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err,
				},
			)
		}
	}()
	defer r.Body.Close()

	res := c.estudianteRepository.MostrarEstudiante()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"estudiantes": res,
		},
		"estudiantes": res,
		"func":        "MostrarEstudiante",
	})
}

// func (c *EstudianteController) AsocisarEstudianteCurso(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	//vars := mux.Vars(r)
// 	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
// 	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
// 	defer func() {
// 		if err := recover(); err != nil {
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(
// 				map[string]interface{}{
// 					"Test": "Validacion de Excepcion",
// 				},
// 			)
// 		}
// 	}()
// 	defer r.Body.Close()

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"R1": map[string]interface{}{
// 			"siJuridico": siJuridico,
// 		},
// 		"R2": "AsocisarEstudianteCurso",
// 	})
// }
