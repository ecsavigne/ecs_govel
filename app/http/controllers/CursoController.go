package controllers

import (
	"ecs_govel/app/helpers"
	"encoding/json"
	"net/http"
)

type CursoController struct {
	//cursoRepositoy repositories.CursoRepository
}

func (c *CursoController) RegistrarCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	idConteido, _ := helpers.ValidateCi(r.FormValue("idConteido"))
	dataIngreso := r.FormValue("dataIngreso")
	duracao, _ := helpers.ValidateCi(r.FormValue("duracao"))
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
			"idCurso":     idCurso,
			"idConteido":  idConteido,
			"dataIngreso": dataIngreso,
			"duracao":     duracao,
		},
		"func": "Valor Test Registrar curso",
	})
}

func (c *CursoController) ModificarCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	idConteido, _ := helpers.ValidateCi(r.FormValue("idConteido"))
	dataIngreso := r.FormValue("dataIngreso")
	duracao, _ := helpers.ValidateCi(r.FormValue("duracao"))
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
			"idCurso":     idCurso,
			"idConteido":  idConteido,
			"dataIngreso": dataIngreso,
			"duracao":     duracao,
		},
		"func": "ModificarCurso",
	})
}

func (c *CursoController) EliminarCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
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
		},
		"func": "eliminarCurso",
	})
}

func (c *CursoController) MostrarCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
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
		},
		"func": "MostrarCurso",
	})
}
