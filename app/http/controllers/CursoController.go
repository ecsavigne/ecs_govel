package controllers

import (
	"ecs_govel/app/helpers"
	"ecs_govel/app/models"
	"ecs_govel/app/repositories"
	"encoding/json"
	"net/http"
)

type CursoController struct {
	cursoRepositoy repositories.CursoRepository
}

func (c *CursoController) RegistrarCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	// idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	idConteido, _ := helpers.ValidateInt(r.FormValue("idConteido"))
	dataIngreso, _ := helpers.ValidateFecha(r.FormValue("dataIngreso"))
	duracao, _ := helpers.ValidateInt(r.FormValue("duracao"))
	si_certificado, _ := helpers.ValidateBool(r.FormValue("si_certificado"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Err": err,
				},
			)
		}
	}()
	defer r.Body.Close()
	curso := models.Curso{
		FechaIngreso:  dataIngreso,
		DuracionHora:  duracao,
		SiCertificado: si_certificado,
	}
	contenido_curso := models.CursoContenido{
		ContenidoId: idConteido,
	}
	c.cursoRepositoy.RegistrarCurso(&curso, &contenido_curso)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idConteido":     idConteido,
			"dataIngreso":    dataIngreso,
			"duracao":        duracao,
			"si_certificado": si_certificado,
		},
		"func":   "Valor Test Registrar curso",
		"Status": http.StatusOK,
	})
}

func (c *CursoController) ModificarCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
	idConteido, _ := helpers.ValidateInt(r.FormValue("idConteido"))
	dataIngreso, _ := helpers.ValidateFecha(r.FormValue("dataIngreso"))
	duracao, _ := helpers.ValidateInt(r.FormValue("duracao"))
	si_certificado, _ := helpers.ValidateBool(r.FormValue("si_certificado"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Err": err,
				},
			)
		}
	}()
	defer r.Body.Close()
	curso := models.Curso{
		ID:            idCurso,
		FechaIngreso:  dataIngreso,
		DuracionHora:  duracao,
		SiCertificado: si_certificado,
	}

	contenido_curso := models.CursoContenido{
		ContenidoId: idConteido,
		CursoId:     idCurso,
	}

	c.cursoRepositoy.ModificarCurso(&curso, &contenido_curso)

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
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Err": err,
				},
			)
		}
	}()
	defer r.Body.Close()
	curso := models.Curso{
		ID: idCurso,
	}

	c.cursoRepositoy.EliminarCurso(&curso)

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
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Err": err,
				},
			)
		}
	}()
	defer r.Body.Close()
	curso := models.Curso{
		ID: idCurso,
	}

	res := c.cursoRepositoy.MostrarCurso(&curso)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso": idCurso,
		},
		"func":   "MostrarCurso",
		"cursos": res,
	})
}
