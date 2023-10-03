package controllers

import (
	"encoding/json"
	"net/http"
)

type InstructorController struct{}

func (c *InstructorController) RegistrarInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
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
		"result": 1,
		"func":   "RegistrarInstructor",
	})
}

func (c *InstructorController) ModificarInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
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
		"result": 1,
		"func":   "RegistrarInstructor",
	})
}

func (c *InstructorController) EliminarInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
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
		"result": 1,
		"func":   "EliminarInstructor",
	})
}

func (c *InstructorController) MostrarInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
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
		"result": 1,
		"func":   "MostrarInstructor",
	})
}

func (c *InstructorController) AgregarCursoToInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
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
		"result": 1,
		"func":   "AgregarCursoToInstructor",
	})
}

func (c *InstructorController) ModificarCursoOffInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
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
		"result": 1,
		"func":   "ModificarCursoOffInstructor",
	})
}

func (c *InstructorController) EliminarCursoOffInstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
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
		"result": 1,
		"func":   "EliminarCursoOffInstructor",
	})
}
