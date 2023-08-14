package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-module/carbon"
	"github.com/gorilla/mux"
	"go.mau.fi/whatsmeow"

	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/repositories"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"
)

func (c *Text) SendTextMessage(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-SendTextMessage")
	w.Header().Set("Content-Type", "application/json")
	serviceDate := carbon.Now().Format("Y-m-d H:i:s")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending text message in TextController-SendTextMessage " + r.FormValue("ContactPhone") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("31m[TextController-SendTextMessage] ", vars["CompanyPhone"], "", logMessage, true)
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

	// 3. request params to chat struct
	chat := migrations.Chats{
		Message:             r.FormValue("Message"),
		CompanyPhone:        vars["CompanyPhone"],
		ContactPhone:        strings.Replace(r.FormValue("RemoteJid"), "@s.whatsapp.net", "", -1),
		ResponseMessageText: r.FormValue("PreviousMessage"),
		ResponseMessageID:   r.FormValue("PreviousMessageId"),
		Source:              source,
		UserId:              userId,
		CompanyId:           companyId,
		CompanyWhatsappId:   companyWhatsappId,
	}

	// 4. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 5. send message
	response := repositories.TextRepository.SendTextMessage(chat, serviceDate)

	// 6. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *Text) SendCommentaryMessage(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-SendCommentaryMessage")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception sending commentary message in TextController-SendCommentaryMessage" + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[TextController-SendCommentaryMessage] ", vars["CompanyPhone"], "", logMessage, true)
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
	userId := helpers.ValidateUserId(r.FormValue("UserId"))
	companyId, companyWhatsappId := helpers.ValidateCompanyIdAndCompanyWhatsappId(r.FormValue("CompanyId"), r.FormValue("CompanyWhatsappId"), vars["CompanyPhone"])

	// 3. request params to chat struct
	chat := migrations.Chats{
		Message:             r.FormValue("Message"),
		CompanyPhone:        vars["CompanyPhone"],
		ContactPhone:        strings.Replace(r.FormValue("RemoteJid"), "@s.whatsapp.net", "", -1),
		ResponseMessageText: r.FormValue("PreviousMessage"),
		ResponseMessageID:   r.FormValue("PreviousMessageId"),
		MessageId:           whatsmeow.GenerateMessageID(),
		TypeId:              structs.COMMENTARY_MESSAGE,
	}
	chat.UserId = userId
	chat.CompanyId = companyId
	chat.CompanyWhatsappId = companyWhatsappId

	// 5. send message
	models.ChatModel.SaveMessageIntoChats(chat)

	// 6. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(structs.Response{
		Status: true,
	})
}

func (c *Text) GetFile(w http.ResponseWriter, r *http.Request) { //TODO arreglar respuesta
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	remoteJid := vars["remoteJid"]

	defer func() {
		if r := recover(); r != nil {
			logMessage := ": Recovered from ecxeption when downlaoding picImage for remoteJid " + remoteJid + ". Interface in defer is: " + fmt.Sprintf("%+v", r)
			logg.ErrorLogger.Println("\033[31m[MessageController-getFile] ", logMessage+"\033[0m")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				structs.Response{
					Status:     false,
					Message:    "",
					LogMessage: logMessage,
					ErrorCode:  1,
				},
			)
		}
	}()

	defer r.Body.Close()
	helpers.ValidatePath(vars["path"])

	logg.GeneralLogger.Println("MessajeFileS - Controllers - Contact: " + os.Getenv("MESSAJE_FILES"))
	fileBytes, err := os.ReadFile(os.Getenv("MESSAJE_FILES") + "/messageFiles/" + remoteJid + "/" + vars["path"])
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}
