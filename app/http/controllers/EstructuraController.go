package controllers

import (
	"ecs_govel/app/helpers"
	"encoding/json"
	"net/http"
)

type EstructuraController struct {
	//estructuraRepository repositories.EstructuraRepository
}

func (c *EstructuraController) RegistrarEstructura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	tipoEstructura := r.FormValue("tipoEstructura")
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"erro":        "",
					"descripcion": "Validacion de Excepcion",
					"func":        "RegistrarEstructura",
				},
			)
		}
	}()
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"tipoEstructura": tipoEstructura,
		},
		"func": "RegistrarEstructura",
	})
}

func (c *EstructuraController) MostrarEstructura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idEstructura, _ := helpers.ValidateInt(r.FormValue("idEstructura"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"erro":        "",
					"descripcion": "Validacion de Excepcion",
					"func":        "MostrarEstructura",
				},
			)
		}
	}()
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idEstructura": idEstructura,
		},
		"func": "MostrarEstructura",
	})
}

func (c *EstructuraController) EliminarEstructura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idEstructura, _ := helpers.ValidateInt(r.FormValue("idEstructura"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"erro":        "",
					"descripcion": "Validacion de Excepcion",
					"func":        "EliminarEstructura",
				},
			)
		}
	}()
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idEstructura": idEstructura,
		},
		"func": "EliminarEstructura",
	})
}

func (c *EstructuraController) ModificarEstructura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idEstructura, _ := helpers.ValidateInt(r.FormValue("idEstructura"))
	tipoEstructura := r.FormValue("tipoEstructura")
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"erro":        "",
					"descripcion": "Validacion de Excepcion",
					"func":        "ModificarEstructura",
				},
			)
		}
	}()
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idEstructura":   idEstructura,
			"tipoEstructura": tipoEstructura,
		},
		"func": "EliminarEstructura",
	})
}
