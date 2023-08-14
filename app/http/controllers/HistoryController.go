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

func (c *History) RevokeMessage(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-RevokeMessage")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception deleting the msgid " + vars["message_id"] + "Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[Text-SendTextMessage] ", vars["CompanyPhone"], "", logMessage, true)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				structs.Response{
					Status:     false,
					Message:    "",
					LogMessage: logMessage,
				},
			)
		}
	}()
	defer r.Body.Close()

	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	helpers.ValidateMessageId(r.FormValue("MessageId"))
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])

	// 4. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 5. send message
	response := repositories.HistoriesRepository.RevokeMessage(vars["CompanyPhone"], r.FormValue("RemoteJid"), r.FormValue("MessageId"), r.FormValue("FileUrl"))

	// 6. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *History) GetHistoryPage(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-GetHistoryPage")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception in HistoryController-GetHistoryPage " + r.FormValue("RemoteJid") + " Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[HistoryController-GetHistoryPage] ", vars["CompanyPhone"], "", logMessage, true)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				structs.Response{
					Status:     false,
					Message:    "",
					LogMessage: logMessage,
				},
			)
		}
	}()
	defer r.Body.Close()
	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	helpers.ValidateContactPhone(r.FormValue("RemoteJid"))
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	helpers.ValidatePerPage(r.FormValue("PerPage"))
	helpers.ValidatePage(r.FormValue("Page"))
	helpers.ValidateCompanyId(r.FormValue("CompanyId"))
	helpers.ValidateCompanyWhatsappId(r.FormValue("CompanyWhatsappId"))

	// 3. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 4. send message
	response := repositories.HistoriesRepository.GetHistoryPage(vars["CompanyPhone"], r.FormValue("RemoteJid"), r.FormValue("CompanyId"), r.FormValue("CompanyWhatsappId"), r.FormValue("PerPage"), r.FormValue("Page"))

	// 5. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *History) GetHistoryPageApi(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-GetHistoryPageApi")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception in HistoryController-GetHistoryPageApi " + r.FormValue("RemoteJid") + " Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[HistoryController-GetHistoryPage] ", vars["CompanyPhone"], "", logMessage, true)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				structs.Response{
					Status:     false,
					Message:    "",
					LogMessage: logMessage,
				},
			)
		}
	}()
	defer r.Body.Close()

	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	helpers.ValidateContactPhone(r.FormValue("RemoteJid"))
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	helpers.ValidatePerPage(r.FormValue("PerPage"))
	helpers.ValidatePage(r.FormValue("Page"))
	helpers.ValidateCompanyId(r.FormValue("CompanyId"))
	helpers.ValidateFirstMessageId(r.FormValue("FirstMessageId"))

	// 3. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 4. send message
	response := repositories.HistoriesRepository.GetHistoryPageApi(vars["CompanyPhone"], r.FormValue("RemoteJid"), r.FormValue("CompanyId"), r.FormValue("PerPage"), r.FormValue("Page"), r.FormValue("FirstMessageId"), r.Host)

	// 5. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
