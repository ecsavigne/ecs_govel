package controllers

import (
	"ecs_govel/app/helpers"
	"ecs_govel/app/repositories"
	"encoding/json"
	"net/http"
)

type UsuarioController struct {
	usuarioRepository repositories.UsuarioRepository
}

func (c *UsuarioController) RegistrarUsr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	pass := r.FormValue("pass")
	nome := r.FormValue("nome")
	correio := r.FormValue("correio")

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

	//c.usuarioRepository.RegistrarUsr(nome, correio, pass)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"pass":    pass,
			"correio": correio,
			"nome":    nome,
		},
		"func": "RegistrarUsr",
	})
}

func (c *UsuarioController) CambiarPass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	newPass := r.FormValue("newPass")
	idUsuario, _ := helpers.ValidateInt(r.FormValue("idUsuario"))
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
			"newPass":   newPass,
			"idUsuario": idUsuario,
		},
		"func": "CambiarPass",
	})
}

func (c *UsuarioController) RecuperarPass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	correio, _ := helpers.ValidateCorreio(r.FormValue("correio"))
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
			"correio": correio,
		},
		"func": "RecuperarPass",
	})
}
