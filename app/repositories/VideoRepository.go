package repositories

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang-module/carbon"
	"google.golang.org/protobuf/proto"

	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

// Repositories Video
func (vr *VideoRepositories) SendVideoMessage(chat migrations.Chats, file multipart.File, serviceDate string) structs.Response {
	// 1. save file into storage
	path, mimetype, err := helpers.StoreMessageFile(file, chat.ClientOriginalName, chat.ContactPhone)
	if err != nil {
		panic("Imposible Store Message File")
	}

	// 2. upload file to whatsapp server
	data, _ := os.ReadFile(path)
	uploaded, err := structs.Connections[chat.CompanyPhone].Upload(context.Background(), data, whatsmeow.MediaVideo)
	if err != nil {
		panic("Failed uploading video file to Whatsapp Server: " + fmt.Sprintf("%v", path) + ": " + fmt.Sprintf("%v", err) + "")
	}

	// 3. prepare to send message
	if chat.Message == "" {
		chat.Message = "Vídeo"
	}
	wppmsg := &waProto.Message{VideoMessage: &waProto.VideoMessage{
		Caption:       proto.String(chat.Message),
		Url:           proto.String(uploaded.URL),
		DirectPath:    proto.String(uploaded.DirectPath),
		MediaKey:      uploaded.MediaKey,
		Mimetype:      proto.String(http.DetectContentType(data)),
		FileEncSha256: uploaded.FileEncSHA256,
		FileSha256:    uploaded.FileSHA256,
		FileLength:    proto.Uint64(uint64(len(data))),
	}}
	if chat.ResponseMessageID != "" {
		QuotedTarget := ""
		if chat.Source == 1 {
			QuotedTarget = chat.ContactPhone + "@s.whatsapp.net"
		} else {
			QuotedTarget = chat.CompanyPhone + "@s.whatsapp.net"
		}
		wppmsg.VideoMessage.ContextInfo = &waProto.ContextInfo{
			StanzaId:      proto.String(chat.ResponseMessageID),
			Participant:   proto.String(QuotedTarget),
			QuotedMessage: &waProto.Message{},
		}
		wppmsg.VideoMessage.Caption = proto.String(chat.ResponseMessageText)
	}
	ContactJID := types.JID{
		User:   chat.ContactPhone,
		Server: types.DefaultUserServer,
	}

	// 4. send message to Whatsapp
	timer, err := structs.Connections[chat.CompanyPhone].SendMessage(context.Background(), ContactJID, wppmsg)
	// msgid := whatsmeow.GenerateMessageID()
	msgid := timer.ID
	if err != nil {
		panic("Error sending video. Error was:" + err.Error())
	} else {
		logg.GeneralLogger.Println("Message ", msgid, " sended at ", timer)
	}

	// 5. save logs
	whatsappDate := carbon.Now().Format("Y-m-d H:i:s")
	logMessage := "Client " + chat.CompanyPhone + ": Video message sended successfully to ContactPhone " + chat.ContactPhone + ". Caption is: " + chat.Message
	logg.Log("[VideoRepository-SendVideoMessage]", chat.CompanyPhone, chat.ContactPhone, logMessage, false)

	Url := ""
	dns := structs.Dns[chat.CompanyPhone]
	filename := msgid + filepath.Ext(chat.ClientOriginalName)
	switch chat.Source {
	case 0:
		Url = "https://" + dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + chat.ContactPhone + "/" + filename
	case 2:
		Url = "https://" + dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + chat.ContactPhone + "/" + filename
	case 4:
		Url = "https://shippingnew.socialhub.pro/external_files/companies/" + strconv.FormatUint(chat.CompanyId, 10) + "/shippings_files/" + chat.ClientOriginalName
	case 5:
		Url = "https://" + dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + chat.ContactPhone + "/" + filename
	case 6:
		Url = "https://" + dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + chat.ContactPhone + "/" + filename
	}

	// 6. Save file in Google Drive
	// idFolderParent, fileUploadId := helpers.SaveInGoogleDrive(path, mimetype, chat.ContactPhone)
	idFolderParent, fileUploadId := "", ""+mimetype

	// 7. save message in chat table
	chat.StatusId = 2
	chat.Status = 1
	chat.TypeId = 4
	chat.MessageId = msgid
	chat.WhatsappDate = whatsappDate
	chat.ServiceDate = serviceDate
	chat.Path = "https://drive.google.com/uc?export=view&id=" + fileUploadId
	chat.IdFileDrive = fileUploadId
	chat.IdParentFolderDrive = idFolderParent
	chat.Url = Url
	models.ChatModel.Create(chat)
	// os.Remove(path)

	// 8. return response
	return structs.Response{
		Status:  true,
		Message: msgid,
	}
}
