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

func (c *Image) SendImageMessage(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-SendImageMessage")
	debugMessage := "1"
	w.Header().Set("Content-Type", "application/json")
	serviceDate := carbon.Now().Format("Y-m-d H:i:s")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending image message in ImageController-SendImageMessage " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err) + ". DebugMessage is: " + debugMessage
			logg.Log("[ImageController-SendImageMessage] ", vars["CompanyPhone"], "", logMessage, true)
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
	source := helpers.ValidateSource(r.FormValue("Source"))
	userId := helpers.ValidateUserId(r.FormValue("UserId"))
	companyId, companyWhatsappId := helpers.ValidateCompanyIdAndCompanyWhatsappId(r.FormValue("CompanyId"), r.FormValue("CompanyWhatsappId"), vars["CompanyPhone"])
	debugMessage = "2"
	file, handler, err := r.FormFile("Image")
	if err != nil {
		panic("\033[31mMissing Image file in request\033[0m")
	}

	debugMessage = "3"

	// 3. request params to chat struct
	chat := migrations.Chats{
		Message:             r.FormValue("Message"),
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
	debugMessage = "4"
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 5. send message
	debugMessage = "5"
	response := repositories.ImageRepository.SendImageMessage(chat, file, serviceDate)

	// 6. http response
	debugMessage = "6"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
