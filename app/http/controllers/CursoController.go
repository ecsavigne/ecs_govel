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
	idConteido, _ := helpers.ValidateInt(r.FormValue("idConteido"))
	dataCreacion, err := helpers.ValidateFecha(r.FormValue("dataCreacion"))
	duracao, _ := helpers.ValidateInt(r.FormValue("duracao"))
	si_certificado, _ := helpers.ValidateBool(r.FormValue("si_certificado"))
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

	if err != nil {
		panic("Error formateando fecha en [Controller-Curso.RegistrarCurso] - Error: " + err.Error() + " Fecha in :" + r.FormValue("dataCreacion"))
	}

	curso := models.Curso{
		FechaCreacion: dataCreacion,
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
			"FechaCreacion":  dataCreacion,
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
	idContenidoAsoc, _ := helpers.ValidateInt(r.FormValue("idContenidoAsoc"))
	dataCreacion, _ := helpers.ValidateFecha(r.FormValue("dataCreacion"))
	duracao, _ := helpers.ValidateInt(r.FormValue("duracao"))
	si_certificado, _ := helpers.ValidateBool(r.FormValue("si_certificado"))
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
	curso := models.Curso{
		ID:            idCurso,
		FechaCreacion: dataCreacion,
		DuracionHora:  duracao,
		SiCertificado: si_certificado,
	}

	contenido_curso := models.CursoContenido{
		ContenidoId: idConteido,
		CursoId:     idCurso,
	}

	c.cursoRepositoy.ModificarCurso(&curso, &contenido_curso, idContenidoAsoc)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso":       idCurso,
			"idConteido":    idConteido,
			"FechaCreacion": dataCreacion,
			"duracao":       duracao,
		},
		"func": "ModificarCurso",
	})
}

func (c *CursoController) EliminarCurso(w http.ResponseWriter, r *http.Request) {
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
	curso := models.Curso{
		ID: idCurso,
	}

	c.cursoRepositoy.EliminarCurso(&curso)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso": idCurso,
		},
		"idCurso": idCurso,
		"func":    "eliminarCurso",
	})
}

func (c *CursoController) MostrarCursoById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
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
	curso := models.Curso{
		ID: idCurso,
	}

	res := c.cursoRepositoy.MostrarCurso(&curso)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idCurso": idCurso,
		},
		"func":   "MostrarCursoById",
		"cursos": res,
	})
}

func (c *CursoController) MostrarCursos(w http.ResponseWriter, r *http.Request) {
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

	res := c.cursoRepositoy.MostrarCursos()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"cursos": res,
		},
		"func": "MostrarCursos",
	})
}

func (c *CursoController) MostrarCursosContenido(w http.ResponseWriter, r *http.Request) {
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

	res := c.cursoRepositoy.MostrarCursosContenido()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"cursos": res,
		},
		"cursos": res,
		"func":   "MostrarCursos",
	})
}
