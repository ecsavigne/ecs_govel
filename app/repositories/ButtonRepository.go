package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang-module/carbon"
	"google.golang.org/protobuf/proto"

	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"

	waProto "go.mau.fi/whatsmeow/binary/proto"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

// Repositories Button
func (b *ButtonRepositories) SendButtonMessage(chat migrations.Chats, serviceDate string) structs.Response {
	// 1. prepare to send message
	buttonsText := make([]string, 0)
	json.Unmarshal([]byte(chat.ButtonsJson), &buttonsText)
	//fmt.Println("buttons are: ", buttonsText)

	wppmsg := &waProto.Message{
		ViewOnceMessage: &waProto.FutureProofMessage{
			Message: &waProto.Message{
				ButtonsMessage: &waProto.ButtonsMessage{
					HeaderType: waProto.ButtonsMessage_EMPTY.Enum(),
					Header:     nil,
					// HeaderType: waProto.ButtonsMessage_TEXT.Enum(),
					// Header: &waProto.ButtonsMessage_Text{
					// 	Text: "Header",
					// },
					// FooterText:  proto.String("_Selecione apenas uma opção_"),
					ContentText: proto.String(chat.Message),
				},
				MessageContextInfo: &waProto.MessageContextInfo{
					DeviceListMetadataVersion: proto.Int32(2),
					DeviceListMetadata:        &waProto.DeviceListMetadata{},
				},
			},
		},
	}

	Buttons := make([]*waProto.ButtonsMessage_Button, 0)
	for i, text := range buttonsText {
		var btn = waProto.ButtonsMessage_Button{
			ButtonId: proto.String(fmt.Sprintf("id%d", i)),
			ButtonText: &waProto.ButtonsMessage_Button_ButtonText{
				DisplayText: proto.String(text),
			},
			Type: waProto.ButtonsMessage_Button_RESPONSE.Enum(),
		}
		Buttons = append(Buttons, &btn)
	}
	wppmsg.ViewOnceMessage.Message.ButtonsMessage.Buttons = Buttons

	ContactJID := types.JID{
		User:   chat.ContactPhone,
		Server: types.DefaultUserServer,
	}

	msgid := whatsmeow.GenerateMessageID()

	//fmt.Println("here 2")
	// 2. send message to Whatsapp
	timer, err := structs.Connections[chat.CompanyPhone].SendMessage(context.Background(), ContactJID, wppmsg)
	if err != nil {
		panic("Error sending button. Error was:" + err.Error())
	} else {
		logg.GeneralLogger.Println("Message ", msgid, " sended at ", timer)
	}

	// 3. save logs
	whatsappDate := carbon.Now().Format("Y-m-d H:i:s")
	logMessage := "Client " + chat.CompanyPhone + ": Button message sended successfully to ContactPhone " + chat.ContactPhone
	logg.Log("[ButtonMessageRepository-SendButtonMessage]", chat.CompanyPhone, chat.ContactPhone, logMessage, false)

	// 4. save message in chat table
	chat.StatusId = 2
	chat.Status = 1
	chat.TypeId = 9
	chat.MessageId = msgid
	chat.WhatsappDate = whatsappDate
	chat.ServiceDate = serviceDate
	models.ChatModel.Create(chat)

	// 5. return response
	return structs.Response{
		Status:     true,
		Message:    msgid,
		LogMessage: logMessage,
	}
}
