package repositories

import (
	"context"

	"github.com/golang-module/carbon"
	"google.golang.org/protobuf/proto"

	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

// Repositories Text
func (tr *TextRepositories) SendTextMessage(chat migrations.Chats, serviceDate string) structs.Response {
	wppmsg := &waProto.Message{}
	if chat.ResponseMessageID != "" {
		QuotedTarget := ""
		if chat.Source == 1 {
			QuotedTarget = chat.ContactPhone + "@s.whatsapp.net"
		} else {
			QuotedTarget = chat.CompanyPhone + "@s.whatsapp.net"
		}
		msg := &waProto.ExtendedTextMessage{
			Text: proto.String(chat.Message),
			ContextInfo: &waProto.ContextInfo{
				StanzaId:      proto.String(chat.ResponseMessageID),
				Participant:   proto.String(QuotedTarget),
				QuotedMessage: &waProto.Message{}, //Conversation: proto.String(""),
			},
		}
		wppmsg.ExtendedTextMessage = msg
	} else {
		wppmsg = &waProto.Message{
			Conversation: proto.String(chat.Message),
		}
	}

	ContactJID := types.JID{
		User:   chat.ContactPhone,
		Server: types.DefaultUserServer,
	}

	// 2. send message to Whatsapp
	// msgid := whatsmeow.GenerateMessageID()
	timer, err := structs.Connections[chat.CompanyPhone].SendMessage(context.Background(), ContactJID, wppmsg)
	msgid := timer.ID
	if err != nil {
		panic("Error sending text. Error was:" + err.Error())
	} else {
		logg.GeneralLogger.Println("Message ", msgid, " sended at ", timer)
	}

	// 3. save logs
	whatsappDate := carbon.Now().Format("Y-m-d H:i:s")
	logMessage := "Client " + chat.CompanyPhone + ": Text message sended successfully in TextMessageRepository-SendTextMessage " + chat.ContactPhone + ". Message is: " + chat.Message
	logg.Log("[TextMessageRepository-SendTextMessage]", chat.CompanyPhone, chat.ContactPhone, logMessage, false)

	// 4. save message in chat table
	chat.StatusId = 2
	chat.Status = 1
	chat.TypeId = 1
	chat.MessageId = msgid
	chat.WhatsappDate = whatsappDate
	chat.ServiceDate = serviceDate
	models.ChatModel.Create(chat)

	// 5. return response
	return structs.Response{
		Status:  true,
		Message: msgid,
	}
}
