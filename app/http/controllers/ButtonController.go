package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-module/carbon"
	"github.com/gorilla/mux"

	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/repositories"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"
)

func (c *Button) SendButtonMessage(w http.ResponseWriter, r *http.Request) {
	debugMessage := "1"

	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-SendButtonMessage")
	w.Header().Set("Content-Type", "application/json")
	serviceDate := carbon.Now().Format("Y-m-d H:i:s")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending button message in ButtonController-SendButtonMessage " + r.FormValue("CompanyPhone") + ". Interface in defer is: " + fmt.Sprintf("%+v", err) + ". DebugMessage is: " + debugMessage
			logg.Log("[ButtonController-SendButtonMessage] ", vars["CompanyPhone"], "", logMessage, true)
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
	debugMessage = "2"

	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	helpers.ValidateContactPhone(r.FormValue("RemoteJid"))
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	source := helpers.ValidateSource(r.FormValue("Source"))
	userId := helpers.ValidateUserId(r.FormValue("UserId"))
	buttons := r.FormValue("Buttons")
	message := r.FormValue("Message")
	companyId, companyWhatsappId := helpers.ValidateCompanyIdAndCompanyWhatsappId(r.FormValue("CompanyId"), r.FormValue("RemoteJid"), vars["CompanyPhone"])

	debugMessage = "3"

	// 3. request params to chat struct
	chat := migrations.Chats{
		CompanyPhone:        vars["CompanyPhone"],
		ContactPhone:        strings.Replace(r.FormValue("RemoteJid"), "@s.whatsapp.net", "", -1),
		ResponseMessageText: r.FormValue("PreviousMessage"),
		ResponseMessageID:   r.FormValue("PreviousMessageId"),
		Message:             message,
		ButtonsJson:         buttons,
		Source:              source,
		UserId:              userId,
		CompanyId:           companyId,
		CompanyWhatsappId:   companyWhatsappId,
	}
	debugMessage = "4"

	// 4. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])
	debugMessage = "5"

	// 5. send message
	response := repositories.ButtonRepository.SendButtonMessage(chat, serviceDate)
	debugMessage = "6"

	// 6. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
