package controllers

import (
	"ecs_govel/app/helpers"
	"encoding/json"
	"net/http"
)

type InstructorController struct {
	//isnstructorRepository repositories.InstructorRepository
}

func (c *InstructorController) RegistrarInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	nome := r.FormValue("nome")
	sobreNome := r.FormValue("sobreNome")
	enderecao := r.FormValue("enderecao")
	correio, _ := helpers.ValidateCorreio(r.FormValue("idCurso"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Test": "Validacion de Excepcion",
				},
			)
		}
	}()
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso":   idCurso,
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
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	nome := r.FormValue("nome")
	sobreNome := r.FormValue("sobreNome")
	enderecao := r.FormValue("enderecao")
	correio, _ := helpers.ValidateCorreio(r.FormValue("idCurso"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Test": "Validacion de Excepcion",
				},
			)
		}
	}()
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso":   idCurso,
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
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Test": "Validacion de Excepcion",
				},
			)
		}
	}()
	defer r.Body.Close()

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
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Test": "Validacion de Excepcion",
				},
			)
		}
	}()
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci": ci,
		},
		"func": "MostrarInstructor",
	})
}

func (c *InstructorController) AgregarCursoToInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Test": "Validacion de Excepcion",
				},
			)
		}
	}()
	defer r.Body.Close()

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
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Test": "Validacion de Excepcion",
				},
			)
		}
	}()
	defer r.Body.Close()

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
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Test": "Validacion de Excepcion",
				},
			)
		}
	}()
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"ci":      ci,
			"idCurso": idCurso,
		},
		"func": "EliminarCursoOffInstructor",
	})
}
