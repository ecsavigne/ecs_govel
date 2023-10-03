package controllers

import (
	"encoding/json"
	"net/http"
)

type EstudianteController struct{}

func (c *EstudianteController) RegistrarEstudianteEnCurso(w http.ResponseWriter, r *http.Request) {
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
		"func":   "RegistrarEstudianteEnCurso",
	})
}

func (c *EstudianteController) RegistrarEstudiante(w http.ResponseWriter, r *http.Request) {
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
		"func":   "RegistrarEstudiante",
	})
}

func (c *EstudianteController) ModificarEstudiante(w http.ResponseWriter, r *http.Request) {
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
		"func":   "ModificarEstudiante",
	})
}

func (c *EstudianteController) ModificarDatosEstudianteCurso(w http.ResponseWriter, r *http.Request) {
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
		"func":   "ModificarDatosEstudianteCurso",
	})
}

func (c *EstudianteController) EliminarEstudiante(w http.ResponseWriter, r *http.Request) {
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
		"func":   "EliminarEstudiante",
	})
}

func (c *EstudianteController) AsocisarEstudianteCurso(w http.ResponseWriter, r *http.Request) {
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
		"R1": 1,
		"R2": "AsocisarEstudianteCurso",
	})
}

func (c *EstudianteController) ElminarEstudianteCurso(w http.ResponseWriter, r *http.Request) {
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
		"func":   "ElminarEstudianteCurso",
	})
}

func (c *EstudianteController) MostrarEstudianteDeCurso(w http.ResponseWriter, r *http.Request) {
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
		"func":   "MostrarEstudianteDeCurso",
	})
}

func (c *EstudianteController) MostrarEstudianteJuridicoNoJuridico(w http.ResponseWriter, r *http.Request) {
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
		"func":   "MostrarEstudianteJuridicoNoJuridico",
	})
}
