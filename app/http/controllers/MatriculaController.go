package controllers

import (
	"ecs_govel/app/helpers"
	"ecs_govel/app/repositories"
	"encoding/json"
	"net/http"
)

type MatriculaController struct {
	matriculaRepository repositories.MatriculaRepository
}

func (c *MatriculaController) MostrarMatriculaDeCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	res := c.matriculaRepository.MostrarMatricula(idCurso)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso":    idCurso,
			"matriculas": res,
		},
		"funcion": "MostrarMatriculaDeCurso",
	})
}
