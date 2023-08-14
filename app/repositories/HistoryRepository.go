package repositories

import (
	"encoding/json"
	"fmt"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"strconv"

	"go.mau.fi/whatsmeow/types"
)

// Repositories History
func (hr *HistoriesRepositories) RevokeMessage(CompanyPhone string, ContactPhone string, MessageId string, FileUrl string) structs.Response {
	ContactJID := types.JID{
		User:   ContactPhone,
		Server: types.DefaultUserServer,
	}

	_, err := structs.Connections[CompanyPhone].RevokeMessage(ContactJID, MessageId)

	if err != nil {
		panic("Client " + CompanyPhone + ": Error when deleting a message. The error was: " + err.Error())
	}

	models.ChatModel.DeteleChat(MessageId, ContactPhone)

	logMessage := "Client " + CompanyPhone + ": Message " + MessageId + " with contact " + ContactPhone + " deteted successfully."
	logg.Log("[TextMessageRepository-RevokeMessage] ", CompanyPhone, "", logMessage, false)

	// 5. return response
	return structs.Response{
		Status:     true,
		Message:    "",
		LogMessage: logMessage,
	}
}

// Repositories History
func (hr *HistoriesRepositories) GetHistoryPage(CompanyPhone string, ContactPhone string, CompanyId string, CompanyWhatsappId string, PerPage string, Page string) structs.Response {
	perPage, _ := strconv.ParseInt(PerPage, 10, 64)
	page, _ := strconv.ParseInt(Page, 10, 64)

	messages := models.ChatModel.LoadFromChats(CompanyPhone, ContactPhone, CompanyId, CompanyWhatsappId, perPage, page)
	// resp, _ := json.MarshalIndent(messages, "", "")
	resp, _ := json.Marshal(messages)
	fmt.Printf("%v", string(resp))
	// 8. return response
	return structs.Response{
		Status:  true,
		Message: string(resp),
	}
}

// Repositories History
func (hr *HistoriesRepositories) GetHistoryPageApi(CompanyPhone string, ContactPhone string, CompanyId string, PerPage string, Page string, MessageId string, Host string) structs.Response {
	perPage, _ := strconv.ParseInt(PerPage, 10, 64)
	page, _ := strconv.ParseInt(Page, 10, 64)

	messages := models.ChatModel.LoadFromChatsApi(CompanyPhone, ContactPhone, CompanyId, perPage, page, MessageId)

	var messagesResp []structs.ChatsApi
	for _, message := range messages {
		if message.Path != "" {
			message.Path = "https://" + Host + "/S4h_EPRZm-b46kyoUbUJ/getFile/" + ContactPhone + "/" + message.Path
		}
		messagesResp = append(messagesResp, message)
	}

	resp, _ := json.Marshal(messagesResp)

	// 8. return response
	return structs.Response{
		Status:  true,
		Message: string(resp),
	}
}
