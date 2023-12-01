package controllers

import (
	"ecs_govel/app/helpers"
	"ecs_govel/app/models"
	"ecs_govel/app/repositories"
	"encoding/json"
	"net/http"
)

type InstructorController struct {
	instructorRepository repositories.InstructorRepository
}

func (c *InstructorController) RegistrarInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	nome := r.FormValue("nome")
	sobreNome := r.FormValue("sobreNome")
	enderecao := r.FormValue("enderecao")
	correio, _ := helpers.ValidateCorreio(r.FormValue("correio"))
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

	instructor := models.Persona{
		CI:        ci,
		Nombre:    nome,
		Apellidos: sobreNome,
		Dir:       enderecao,
		Mail:      correio,
	}
	c.instructorRepository.RegistrarInstructor(&instructor)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"nome":      nome,
			"sobreNome": sobreNome,
			"ci":        ci,
			"enderecao": enderecao,
			"correio":   correio,
		},
		"func": "RegistrarInstructor",
	})
}

func (c *InstructorController) ModificarInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	nome := r.FormValue("nome")
	sobreNome := r.FormValue("sobreNome")
	enderecao := r.FormValue("enderecao")
	correio, _ := helpers.ValidateCorreio(r.FormValue("idCurso"))
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

	inst := models.Persona{
		CI:        ci,
		Nombre:    nome,
		Apellidos: sobreNome,
		Dir:       enderecao,
		Mail:      correio,
	}
	c.instructorRepository.ModificarInstructor(&inst)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"nome":      nome,
			"sobreNome": sobreNome,
			"ci":        ci,
			"enderecao": enderecao,
			"correio":   correio,
		},
		"func": "RegistrarInstructor",
	})
}

func (c *InstructorController) EliminarInstructor(w http.ResponseWriter, r *http.Request) {
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

	inst := models.Instructore{
		CI: ci,
	}
	c.instructorRepository.EliminarInstructor(&inst)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci": ci,
		},
		"func": "EliminarInstructor",
	})
}

func (c *InstructorController) MostrarInstructor(w http.ResponseWriter, r *http.Request) {
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

	inst := models.Instructore{
		CI: ci,
	}
	res := c.instructorRepository.MostrarInstructor(&inst)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci": ci,
		},
		"func": "MostrarInstructor",
		"data": res,
	})
}

func (c *InstructorController) AgregarCursoToInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
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

	inst_cur := models.InstructorCurso{
		CursoId:      idCurso,
		InstructorCi: ci,
	}
	c.instructorRepository.AgregarInstructorToCurso(&inst_cur)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso": idCurso,
			"ci":      ci,
		},
		"func": "AgregarCursoToInstructor",
	})
}

func (c *InstructorController) ModificarCursoOffInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
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
	inst_cur := models.InstructorCurso{
		CursoId:      idCurso,
		InstructorCi: ci,
	}
	c.instructorRepository.ModificarInstructorCurso(&inst_cur)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso": idCurso,
			"ci":      ci,
		},
		"func": "ModificarCursoOffInstructor",
	})
}

func (c *InstructorController) EliminarCursoOffInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
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
	inst_cur := models.InstructorCurso{
		CursoId:      idCurso,
		InstructorCi: ci,
	}
	c.instructorRepository.EliminarCursoOffInstructor(&inst_cur)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci":      ci,
			"idCurso": idCurso,
		},
		"func": "EliminarCursoOffInstructor",
	})
}

func (c *InstructorController) MostrarInstructores(w http.ResponseWriter, r *http.Request) {
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

	res := c.instructorRepository.MostrarInstructores()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"Instructores": res,
		},
		"instructores": res,
		"func":         "MostrarInstructores",
	})
}

func (c *InstructorController) MostrarCursosOfInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	res := c.instructorRepository.MostrarCursosOfInstructor(ci)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"cursos": res,
		},
		"cursos": res,
		"func":   "MostrarCursosOfInstructor",
	})
}
