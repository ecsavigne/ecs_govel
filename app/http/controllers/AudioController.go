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

func (c *Audio) SendAudioMessage(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-SendAudioMessage")
	w.Header().Set("Content-Type", "application/json")
	serviceDate := carbon.Now().Format("Y-m-d H:i:s")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending audio message in AudioController-ContactPhone " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[AudioController-SendAudioMessage] ", vars["CompanyPhone"], "", logMessage, true)
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
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	helpers.ValidateContactPhone(r.FormValue("RemoteJid"))
	source := helpers.ValidateSource(r.FormValue("Source"))
	userId := helpers.ValidateUserId(r.FormValue("UserId"))
	companyId, companyWhatsappId := helpers.ValidateCompanyIdAndCompanyWhatsappId(r.FormValue("CompanyId"), r.FormValue("CompanyWhatsappId"), vars["CompanyPhone"])
	file, handler, err := r.FormFile("Audio")
	if err != nil {
		panic("Missing Audio file in request")
	}

	// 3. request params to chat struct
	chat := migrations.Chats{
		CompanyPhone:        vars["CompanyPhone"],
		ContactPhone:        strings.Replace(r.FormValue("RemoteJid"), "@s.whatsapp.net", "", -1),
		ResponseMessageText: r.FormValue("PreviousMessage"),
		ResponseMessageID:   r.FormValue("PreviousMessageId"),
		ClientOriginalName:  handler.Filename,
		Source:              source,
		UserId:              userId,
		CompanyId:           companyId,
		CompanyWhatsappId:   companyWhatsappId,
	}

	// 4. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 5. send message
	response := repositories.AudioRepository.SendAudioMessage(chat, file, serviceDate)

	// 6. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
