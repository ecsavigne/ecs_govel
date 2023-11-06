package controllers

import (
	"ecs_govel/app/helpers"
	"ecs_govel/app/models"
	"ecs_govel/app/repositories"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type UsuarioController struct {
	usuarioRepository repositories.UsuarioRepository
}

func (c *UsuarioController) RegistrarUsr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pass := r.FormValue("pass")
	nome := r.FormValue("nome")
	correio := r.FormValue("correio")
	var err error = nil
	defer func() {
		if err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Error": err.Error(),
					"code":  false,
				},
			)
		}
	}()
	defer r.Body.Close()

	res := c.usuarioRepository.RegistrarUsr(nome, correio, pass)
	if res != true {
		err = errors.New("Ocurrio un error al registrar usuario : " + fmt.Sprintf("%v", res))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"code": true,
			"res":  res,
		},
		"func": "RegistrarUsr",
	})
}

func (c *UsuarioController) CambiarPass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newPass := r.FormValue("newPass")
	idUsuario, _ := helpers.ValidateInt(r.FormValue("idUsuario"))
	var err error = nil
	defer func() {
		if err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"Error": err.Error(),
					"code":  false,
				},
			)
		}
	}()
	defer r.Body.Close()

	res := c.usuarioRepository.CambiarPass(idUsuario, newPass)
	if res != true {
		err = errors.New("Ocurried error to changing password : " + fmt.Sprintf("%v", res))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"code": true,
			"res":  res,
		},
		"func": "CambiarPass",
	})
}

func (c *UsuarioController) RecuperarPass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	correio, err := helpers.ValidateCorreio(r.FormValue("correio"))
	defer func() {
		if err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err.Error(),
					"code":  false,
				},
			)
		}
	}()

	if err != nil {
		return
	}

	defer r.Body.Close()

	res := c.usuarioRepository.RecuperarPass(correio)
	if res != true {
		if res == false {
			err = errors.New("Ocurrio un error Recuperando password, no existe correo : " + correio)
		} else {
			err = errors.New("Ocurrio un error Recuperando password : " + fmt.Sprintf("%v", res))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"code": true,
		},
		"func": "RecuperarPass",
	})
}

func (c *UsuarioController) VerificarUsr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usuario := r.FormValue("usuario")
	pass := r.FormValue("passwd")
	var err error = nil
	defer func() {
		if err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error": err.Error(),
					"code":  false,
				},
			)
		}
	}()

	if err != nil {
		return
	}

	defer r.Body.Close()

	usr := models.Usuario{
		Nombre:   usuario,
		Password: pass,
	}

	res := c.usuarioRepository.VerificarUsr(&usr)
	if res != true {
		if res == false {
			err = errors.New("Ocurrio un error Verificando el usuario")
		} else {
			err = errors.New("Ocurrio un error Verificando el usuario : " + fmt.Sprintf("%v", res))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result1": map[string]interface{}{
			"code": true,
		},
		"func": "VerificarUsr",
	})
}
