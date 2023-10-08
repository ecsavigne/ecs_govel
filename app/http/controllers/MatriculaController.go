package controllers

import (
	"ecs_govel/app/helpers"
	"encoding/json"
	"net/http"
)

type MatriculaController struct {
	//matriculaRepository repositories.MatriculaRepository
}

func (c *MatriculaController) MostrarMatriculaDeCurso(w http.ResponseWriter, r *http.Request) {
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
		"funcion": "MostrarMatriculaDeCurso",
	})
}
