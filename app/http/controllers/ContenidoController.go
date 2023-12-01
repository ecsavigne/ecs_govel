package controllers

import (
	"ecs_govel/app/helpers"
	"ecs_govel/app/models"
	"ecs_govel/app/repositories"
	"encoding/json"
	"net/http"
)

type ContenidoController struct {
	contenidoRepositoy repositories.ContenidoRepository
}

func (c *ContenidoController) RegistrarContenido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	tema := r.FormValue("tema")
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error":       err,
					"descripcion": "Validacion de Excepcion",
					"func":        "RegistrarContenido",
				},
			)
		}
	}()
	defer r.Body.Close()
	if tema == "" {
		panic("Tema no puede ser vacio")
	}

	contenido := models.Contenido{
		Tema: tema,
	}
	c.contenidoRepositoy.RegistrarContenido(&contenido)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"tema": tema,
		},
		"func": "RegistrarContenido",
	})
}

func (c *ContenidoController) MostrarContenido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idContenido, _ := helpers.ValidateInt(r.FormValue("idContenido"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error":       err,
					"descripcion": "Validacion de Excepcion",
					"func":        "MostrarContenido",
				},
			)
		}
	}()
	defer r.Body.Close()

	contenido := models.Contenido{
		ID: idContenido,
	}
	res := c.contenidoRepositoy.MostrarContenido(&contenido)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idContenido": idContenido,
		},
		"func":      "MostrarContenido",
		"contenido": res,
	})
}

func (c *ContenidoController) MostrarAllContenido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error":       err,
					"descripcion": "Validacion de Excepcion",
					"func":        "MostrarAllContenido",
				},
			)
		}
	}()
	defer r.Body.Close()

	res := c.contenidoRepositoy.MostrarAllContenido()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"contenidos": res,
		},
		"func":       "MostrarAllContenido",
		"contenidos": res,
	})
}

func (c *ContenidoController) EliminarContenido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idContenido, _ := helpers.ValidateInt(r.FormValue("idContenido"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error":       err,
					"descripcion": "Validacion de Excepcion",
					"func":        "EliminarContenido",
				},
			)
		}
	}()
	defer r.Body.Close()

	contenido := models.Contenido{
		ID: idContenido,
	}
	c.contenidoRepositoy.EliminarContenido(&contenido)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idContenido": idContenido,
		},
		"func": "EliminarContenido",
	})
}

func (c *ContenidoController) ModificarContenido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idContenido, _ := helpers.ValidateInt(r.FormValue("idContenido"))
	tema := r.FormValue("tema")
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error":       err,
					"descripcion": "Validacion de Excepcion",
					"func":        "ModificarContenido",
				},
			)
		}
	}()
	defer r.Body.Close()

	contenido := models.Contenido{
		ID:   idContenido,
		Tema: tema,
	}
	c.contenidoRepositoy.ModificarContenido(&contenido)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idContenido": idContenido,
			"tema":        tema,
		},
		"func": "ModificarContenido",
	})
}

func (c *ContenidoController) AsociarContenidoToEstructura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idContenido, _ := helpers.ValidateInt(r.FormValue("idContenido"))
	idEstructura, _ := helpers.ValidateInt(r.FormValue("idEstructura"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error":       err,
					"descripcion": "Validacion de Excepcion",
					"func":        "asociarContenidoToEstructura",
				},
			)
		}
	}()
	defer r.Body.Close()

	est_cont := models.EstructuraContenido{
		ContenidoId:  idContenido,
		EstructuraId: idEstructura,
	}
	c.contenidoRepositoy.AsociarContenidoToEstructura(&est_cont)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idContenido":  idContenido,
			"idEstructura": idEstructura,
		},
		"func": "asociarContenidoToEstructura",
	})
}

func (c *ContenidoController) ModificarContenidoToEstructura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	idContenido, _ := helpers.ValidateInt(r.FormValue("idContenido"))
	idEstructura, _ := helpers.ValidateInt(r.FormValue("idEstructura"))
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(509)
			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"error":       err,
					"descripcion": "Validacion de Excepcion",
					"func":        "modificarContenidoToEstructura",
				},
			)
		}
	}()
	defer r.Body.Close()

	est_cont := models.EstructuraContenido{
		ContenidoId:  idContenido,
		EstructuraId: idEstructura,
	}
	c.contenidoRepositoy.ModificarContenidoToEstructura(&est_cont)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": map[string]interface{}{
			"idContenido":  idContenido,
			"idEstructura": idEstructura,
		},
		"func": "modificarContenidoToEstructura",
	})
}
