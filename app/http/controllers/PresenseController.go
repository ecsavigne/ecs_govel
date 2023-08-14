package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/repositories"
	"new_whatsmeow/app/structs"

	"github.com/gorilla/mux"
)

func (c *Presense) SendChatPresence(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-SendChatPresence")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Recovered from ecxeption when SendChatPresence in PresenseController-SendChatPresence " + vars["RemoteJid"] + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[PresenseController-SendChatPresence] ", vars["CompanyPhone"], vars["RemoteJid"], logMessage, true)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				structs.Response{
					Status:     true,
					Message:    "",
					LogMessage: logMessage,
				},
			)
		}
	}()
	defer r.Body.Close()

	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	helpers.ValidateContactPhone(vars["RemoteJid"])
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	helpers.ValidateAction(r.FormValue("Action"))

	// 4. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 5. send message
	response := repositories.PresenseRepository.SendChatPresence(vars["CompanyPhone"], vars["RemoteJid"], r.FormValue("Action"))

	// 6. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
