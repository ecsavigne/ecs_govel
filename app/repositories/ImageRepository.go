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

// Repositories Image
func (ir *ImageRepositories) SendImageMessage(chat migrations.Chats, file multipart.File, serviceDate string) structs.Response {
	// 1. save file into storage
	fmt.Println("Debug 5.1")
	path, mimetype, err := helpers.StoreMessageFile(file, chat.ClientOriginalName, chat.ContactPhone)
	fmt.Println("Debug 5.2")
	if err != nil {
		panic("Imposible Store Message File")
	}
	// 2. upload file to whatsapp server
	fmt.Println("Debug 5.3")
	data, _ := os.ReadFile(path)
	fmt.Println("Debug 5.4")
	uploaded, err := structs.Connections[chat.CompanyPhone].Upload(context.Background(), data, whatsmeow.MediaImage)
	fmt.Println("Debug 5.5")
	if err != nil {
		panic("Failed uploading image file to Whatsapp Server: " + fmt.Sprintf("%v", path) + ": " + fmt.Sprintf("%v", err))
	}

	// 3. prepare to send message
	if chat.Message == "" {
		chat.Message = "Foto"
	}
	wppmsg := &waProto.Message{ImageMessage: &waProto.ImageMessage{
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
		wppmsg.ImageMessage.ContextInfo = &waProto.ContextInfo{
			StanzaId:      proto.String(chat.ResponseMessageID),
			Participant:   proto.String(QuotedTarget),
			QuotedMessage: &waProto.Message{},
		}
	}
	ContactJID := types.JID{
		User:   chat.ContactPhone,
		Server: types.DefaultUserServer,
	}
	fmt.Println("Debug 5.6")
	// 4. send message to Whatsapp
	fmt.Println("Debug 5.7")
	timer, err := structs.Connections[chat.CompanyPhone].SendMessage(context.Background(), ContactJID, wppmsg)
	// msgid := whatsmeow.GenerateMessageID()
	msgid := timer.ID
	if err != nil {
		panic("Error sending image. Error was:" + err.Error())
	} else {
		logg.GeneralLogger.Println("Message ", msgid, " sended at ", timer)
	}
	// 5. save logs
	whatsappDate := carbon.Now().Format("Y-m-d H:i:s")
	logMessage := "Client " + chat.CompanyPhone + ": Image message sended successfully to ContactPhone " + chat.ContactPhone + ". Caption is: " + chat.Message
	logg.Log("[ImageMessageRepository-SendImageMessage]", chat.CompanyPhone, chat.ContactPhone, logMessage, false)

	Url := ""
	filename := msgid + filepath.Ext(chat.ClientOriginalName)
	dns := structs.Dns[chat.CompanyPhone]
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
	fmt.Println("Debug 5.8")
	// idFolderParent, fileUploadId := helpers.SaveInGoogleDrive(path, mimetype, chat.ContactPhone)
	idFolderParent, fileUploadId := "", ""+mimetype
	fmt.Println("Debug 5.9")
	// 7. save message in chat table
	chat.StatusId = 2
	chat.Status = 1
	chat.TypeId = 2
	chat.MessageId = msgid
	chat.WhatsappDate = whatsappDate
	chat.ServiceDate = serviceDate
	// chat.Path = "https://drive.google.com/uc?export=view&id=" + fileUploadId
	chat.IdFileDrive = fileUploadId
	chat.IdParentFolderDrive = idFolderParent
	chat.Url = Url
	models.ChatModel.Create(chat)
	// os.Remove(path)

	//os.Remove(path)
	// 8. return response
	return structs.Response{
		Status:  true,
		Message: msgid,
	}
}
