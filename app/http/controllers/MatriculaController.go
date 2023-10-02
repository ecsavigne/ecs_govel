package controllers

import (
	"encoding/json"
	"net/http"
)

type MatriculaController struct{}

func (c *MatriculaController) MostrarMatriculaDeCurso(w http.ResponseWriter, r *http.Request) {
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
		"result": "Joson resultado"
		"funcion": "MostrarMatriculaDeCurso",
	})
}
