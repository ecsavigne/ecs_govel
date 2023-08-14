package repositories

import (
	"context"
	"encoding/json"

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

// Repositories List
func (lr *ListRepositories) SendListMessage(chat migrations.Chats, serviceDate string) structs.Response {
	// 1. prepare to send message
	ListsText := make([]string, 0)
	json.Unmarshal([]byte(chat.ListJson), &ListsText)

	wppmsg := &waProto.Message{
		ListMessage: &waProto.ListMessage{
			Title:       proto.String("ListMessage title"),
			Description: proto.String("ListMessage Description"),
			FooterText:  proto.String("ListMessage footer"),
			ButtonText:  proto.String("ListMessage ListText"),
			ListType:    waProto.ListMessage_SINGLE_SELECT.Enum(),
			Sections: []*waProto.ListMessage_Section{
				{
					Title: proto.String("Section1 title"),
					Rows: []*waProto.ListMessage_Row{
						{
							RowId:       proto.String("id1"),
							Title:       proto.String("ListMessage section row title"),
							Description: proto.String("ListMessage section row desc"),
						},
						{
							RowId:       proto.String("id2"),
							Title:       proto.String("title 2"),
							Description: proto.String("desc 2"),
						},
					},
				},
				{
					Title: proto.String("Section2 title"),
					Rows: []*waProto.ListMessage_Row{
						{
							RowId:       proto.String("id1"),
							Title:       proto.String("ListMessage section row title"),
							Description: proto.String("ListMessage section row desc"),
						},
						{
							RowId:       proto.String("id2"),
							Title:       proto.String("title 2"),
							Description: proto.String("desc 2"),
						},
					},
				},
			},
		},
	}

	ContactJID := types.JID{
		User:   chat.ContactPhone,
		Server: types.DefaultUserServer,
	}

	msgid := whatsmeow.GenerateMessageID()

	// 2. send message to Whatsapp
	timer, err := structs.Connections[chat.CompanyPhone].SendMessage(context.Background(), ContactJID, wppmsg)
	if err != nil {
		panic("Error sending List. Error was:" + err.Error())
	} else {
		logg.GeneralLogger.Println("Message ", msgid, " sended at ", timer)
	}

	// 3. save logs
	whatsappDate := carbon.Now().Format("Y-m-d H:i:s")
	logMessage := "Client " + chat.CompanyPhone + ": List message sended successfully to ContactPhone " + chat.ContactPhone
	logg.Log("[ListRepository-SendListMessage]", chat.CompanyPhone, chat.ContactPhone, logMessage, false)

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
