package controllers

import (
	"ecs_govel/app/helpers"
	"encoding/json"
	"net/http"
)

type EstudianteController struct {
	//estudianteRepository repositories.EstudianteRepository
}

func (c *EstudianteController) RegistrarEstudianteEnCurso(w http.ResponseWriter, r *http.Request) {
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
		"func": "RegistrarEstudianteEnCurso",
	})
}

func (c *EstudianteController) RegistrarEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	nome := r.FormValue("nome")
	sobreNome := r.FormValue("sobreNome")
	enderecao := r.FormValue("enderecao")
	correio, _ := helpers.ValidateCorreio(r.FormValue("idCurso"))
	siJuridico := r.FormValue("siJuridico")
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
			"ci":         ci,
			"nome":       nome,
			"sobreNome":  sobreNome,
			"enderecao":  enderecao,
			"correio":    correio,
			"siJuridico": siJuridico,
		},
		"func": "RegistrarEstudiante",
	})
}

func (c *EstudianteController) ModificarEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
	nome := r.FormValue("nome")
	sobreNome := r.FormValue("sobreNome")
	enderecao := r.FormValue("enderecao")
	correio, _ := helpers.ValidateCorreio(r.FormValue("idCurso"))
	siJuridico := r.FormValue("siJuridico")
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
			"ci":         ci,
			"nome":       nome,
			"sobreNome":  sobreNome,
			"enderecao":  enderecao,
			"correio":    correio,
			"siJuridico": siJuridico,
		},
		"func": "ModificarEstudiante",
	})
}

func (c *EstudianteController) ModificarDatosEstudianteCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
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
			"ci":      ci,
		},
		"func": "ModificarDatosEstudianteCurso",
	})
}

func (c *EstudianteController) EliminarEstudiante(w http.ResponseWriter, r *http.Request) {
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
		"func": "EliminarEstudiante",
	})
}

func (c *EstudianteController) ElminarEstudianteCurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
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
			"ci":      ci,
			"idCurso": idCurso,
		},
		"func": "ElminarEstudianteCurso",
	})
}

func (c *EstudianteController) MostrarEstudianteDeCurso(w http.ResponseWriter, r *http.Request) {
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
		"func": "MostrarEstudianteDeCurso",
	})
}

func (c *EstudianteController) MostrarEstudianteJuridicoNoJuridico(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	siJuridico := r.FormValue("siJuridico")
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
			"siJuridico": siJuridico,
		},
		"func": "MostrarEstudianteJuridicoNoJuridico",
	})
}

// func (c *EstudianteController) AsocisarEstudianteCurso(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	//vars := mux.Vars(r)
// 	ci, _ := helpers.ValidateCi(r.FormValue("ci"))
// 	idCurso, _ := helpers.ValidateInt(r.FormValue("idCurso"))
// 	defer func() {
// 		if err := recover(); err != nil {
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(
// 				map[string]interface{}{
// 					"Test": "Validacion de Excepcion",
// 				},
// 			)
// 		}
// 	}()
// 	defer r.Body.Close()

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"R1": map[string]interface{}{
// 			"siJuridico": siJuridico,
// 		},
// 		"R2": "AsocisarEstudianteCurso",
// 	})
// }
