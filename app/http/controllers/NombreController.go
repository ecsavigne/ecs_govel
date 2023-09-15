package controllers

// Ejemplo de funciones miembro de las modelos
// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/golang-module/carbon"
// 	"github.com/gorilla/mux"

//	"app/helpers"
//	"app/helpers/logg"
//	"app/repositories"
//	"app/structs"
//	"database/migrations"
//
// )
// func (c *Nombre) Insert(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "application/json")
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

// func (c *Nombre) Delete(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "application/json")
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

// func (c *Nombre) Update(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "application/json")
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

// func (c *Nombre) Show(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "application/json")
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
