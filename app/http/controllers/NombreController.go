package controllers

import (
	"encoding/json"
	"net/http"
)

// Ejemplo de funciones miembro de las modelos
// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/golang-module/carbon"
// 	"github.com/gorilla/mux"

//	"ecs_govel/app/helpers"
//	"ecs_govel/app/helpers/logg"
//	"ecs_govel/app/repositories"
//	"ecs_govel/app/structs"
//	"database/migrations"
//
// )
type NombreController struct{}

func (c *NombreController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "ecs_govel/application/json")
	//vars := mux.Vars(r)

	//si sucede algun error
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending audio message in AudioController-ContactPhone " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
	// 		logg.Log("[AudioController-SendAudioMessage] ", vars["CompanyPhone"], "", logMessage, true)
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(
	// 			structs.Response{
	// 				Status:     false,
	// 				Message:    "",
	// 				LogMessage: logMessage,
	// 			},
	// 		)
	// 	}
	// }()
	defer r.Body.Close()

	// if err != nil {
	// 	panic("Missing Audio file in request")
	// }

	// Si todo sale bien
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(make(map[string]interface{}, 0))
}

// func (c *NombreController) Delete(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "ecs_govel/application/json")
// vars := mux.Vars(r)

// si sucede algun error
// 	defer func() {
// 		if err := recover(); err != nil {
// 			logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending audio message in AudioController-ContactPhone " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
// 			logg.Log("[AudioController-SendAudioMessage] ", vars["CompanyPhone"], "", logMessage, true)
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(
// 				structs.Response{
// 					Status:     false,
// 					Message:    "",
// 					LogMessage: logMessage,
// 				},
// 			)
// 		}
// 	}()
// 	defer r.Body.Close()

//	if err != nil {
//			panic("Missing Audio file in request")
//		}
//
// Si todo sale bien
// w.WriteHeader(http.StatusOK)
// json.NewEncoder(w).Encode(response)
//}

// func (c *NombreController) Update(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "ecs_govel/application/json")
// vars := mux.Vars(r)

// si sucede algun error
// 	defer func() {
// 		if err := recover(); err != nil {
// 			logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending audio message in AudioController-ContactPhone " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
// 			logg.Log("[AudioController-SendAudioMessage] ", vars["CompanyPhone"], "", logMessage, true)
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(
// 				structs.Response{
// 					Status:     false,
// 					Message:    "",
// 					LogMessage: logMessage,
// 				},
// 			)
// 		}
// 	}()
// 	defer r.Body.Close()

//	if err != nil {
//			panic("Missing Audio file in request")
//		}
//
// Si todo sale bien
// w.WriteHeader(http.StatusOK)
// json.NewEncoder(w).Encode(response)
//}

// func (c *NombreController) Show(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "ecs_govel/application/json")
// vars := mux.Vars(r)

// si sucede algun error
// 	defer func() {
// 		if err := recover(); err != nil {
// 			logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending audio message in AudioController-ContactPhone " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
// 			logg.Log("[AudioController-SendAudioMessage] ", vars["CompanyPhone"], "", logMessage, true)
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(
// 				structs.Response{
// 					Status:     false,
// 					Message:    "",
// 					LogMessage: logMessage,
// 				},
// 			)
// 		}
// 	}()
// 	defer r.Body.Close()

//	if err != nil {
//			panic("Missing Audio file in request")
//		}
//
// Si todo sale bien
// w.WriteHeader(http.StatusOK)
// json.NewEncoder(w).Encode(response)
//}
