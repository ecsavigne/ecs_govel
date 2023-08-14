package webhooks

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"

	"github.com/golang-module/carbon"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func ReceiveVideoWebhook(vInfo types.MessageInfo, vMessage *proto.Message, companyWhatsapp string, wac *whatsmeow.Client, Dns string) {
	// 0. Load env vars
	webhook := os.Getenv("BASE_URL") + "/api/RPI/reciveFileMessage"

	// 1. Get chat datas (part 1)
	contactJid := ""
	Source := 1
	msgid := vInfo.ID
	whatsappDate := vInfo.Timestamp.Format("2006-01-02 15:04:05")
	serviceDate := carbon.Now().Format("Y-m-d H:i:s")
	text := vMessage.VideoMessage.GetCaption()
	ResponseMessageId := ""
	if vMessage.VideoMessage.ContextInfo != nil {
		ResponseMessageId = vMessage.VideoMessage.ContextInfo.GetStanzaId()
	}
	if vInfo.IsFromMe {
		Source = 3
		contactJid = vInfo.MessageSource.Chat.User
		if text != "" {
			text = text + "\n"
		}
	} else {
		contactJid = vInfo.MessageSource.Sender.User
	}

	// 2. Capture exception
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Exception in ReceiveVideoWebhook for " + contactJid + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveVideoWebhook-ReceiveVideoWebhook] ", companyWhatsapp, contactJid, logMessage, true)
			return
		}
	}()

	// 3. Verify is message_id already exist in chats table
	err := models.ChatModel.AlreadyExistChat(vInfo.ID)
	if err != nil {
		logg.ErrorLogger.Println(err)
	}

	// 4. Download audio file from Whatsapp
	localPath, filename, clientOriginalName, idParentDrive, idFileDrive := "", "", "", "", ""
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	textoMessage := ""
	if Source != 3 {
		localPath, filename, clientOriginalName, idParentDrive, idFileDrive, _ = DownloadVideoFromWPP(vInfo, vMessage, companyWhatsapp, wac, contactJid)
		file, err := os.Open(localPath)
		if err != nil {
			return
		}
		defer file.Close()

		part, err := writer.CreateFormFile("File", filepath.Base(localPath))
		if err != nil {
			return
		}
		_, _ = io.Copy(part, file)
	} else {
		text = ""
		textoMessage = "Uma mensagem de video foi enviada por você desde outro dispositivo. Pode visualizá-la no aplicativo ou no whatsapp web."
	}

	// 5. Save file in Google Drive
	// srv, _ := services.GDrive.GetService()
	// gdFile, _ := services.GDrive.CreateFile(srv, filename, mimetype, file, "")

	// 6. Get chat datas (part 2)
	/*
		responseText := ""
		responsePath := ""
		responseType := "1"
		responseDuration := "0"
		responseOriginalName := "0"
		// responseSource := ""
		responseJpegThumbnail := ""
		if message.ContextInfo.QuotedMessageID != "" {
			if message.ContextInfo.QuotedMessage.Conversation != nil {
				responseText = *message.ContextInfo.QuotedMessage.Conversation
			}
			if message.ContextInfo.QuotedMessage.ImageMessage != nil {
				mimetype := *message.ContextInfo.QuotedMessage.ImageMessage.Mimetype
				ext := DocHandlerExts[mimetype]
				responseType = "2"
				responsePath = fmt.Sprintf("%v.%v", message.ContextInfo.QuotedMessageID, ext)
			}
			if message.ContextInfo.QuotedMessage.AudioMessage != nil {
				mimetype := *message.ContextInfo.QuotedMessage.AudioMessage.Mimetype
				ext := DocHandlerExts[mimetype]
				responseType = "3"
				responsePath = fmt.Sprintf("%v.%v", message.ContextInfo.QuotedMessageID, ext)
				responseDuration = fmt.Sprintf("%v", int(*message.ContextInfo.QuotedMessage.AudioMessage.Seconds))
			}
			if message.ContextInfo.QuotedMessage.VideoMessage != nil {
				mimetype := *message.ContextInfo.QuotedMessage.VideoMessage.Mimetype
				ext := DocHandlerExts[mimetype]
				responseType = "4"
				responsePath = fmt.Sprintf("%v.%v", message.ContextInfo.QuotedMessageID, ext)
				responseDuration = fmt.Sprintf("%v", int(*message.ContextInfo.QuotedMessage.VideoMessage.Seconds))
			}
			if message.ContextInfo.QuotedMessage.DocumentMessage != nil {
				mimetype := *message.ContextInfo.QuotedMessage.DocumentMessage.Mimetype
				ext := DocHandlerExts[mimetype]
				responseType = "5"
				responsePath = fmt.Sprintf("%v.%v", message.ContextInfo.QuotedMessageID, ext)
				responseOriginalName = *message.ContextInfo.QuotedMessage.DocumentMessage.FileName
				responseJpegThumbnail = string(message.ContextInfo.QuotedMessage.DocumentMessage.JpegThumbnail)
			}
		}
	*/

	Url := "https://" + Dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + contactJid + "/" + vInfo.ID + filepath.Ext(clientOriginalName)

	// 7. Save message into chats table
	company, _ := models.CompanyWhatsappModel.GetCompanyWhatsapp(companyWhatsapp)
	chat := migrations.Chats{
		CompanyId:           company.CompanyId,
		CompanyWhatsappId:   company.CompanyWhatsappId,
		StatusId:            1,
		Status:              1,
		TypeId:              structs.VIDEO_MESSAGE,
		CompanyPhone:        companyWhatsapp,
		ContactPhone:        contactJid,
		Source:              Source,
		MessageId:           msgid,
		WhatsappDate:        whatsappDate,
		ServiceDate:         serviceDate,
		PathDrive:           os.Getenv("GOOGLE_DRIVE_URL") + idFileDrive,
		Path:                filename,
		ClientOriginalName:  clientOriginalName,
		IdParentFolderDrive: idParentDrive,
		IdFileDrive:         idFileDrive,
		Message:             text + textoMessage,
		ResponseMessageID:   ResponseMessageId,
		Url:                 Url,
	}
	// if previousMessageId != "" {
	// 	chat.ResponseMessageID = previousMessageId
	// }
	// if previousMessage != "" {
	// 	chat.ResponseMessageText = previousMessage
	// }
	models.ChatModel.SaveMessageIntoChats(chat)

	// 8. Send request to Principal Laravel App

	_ = writer.WriteField("company_id", fmt.Sprintf("%d", company.CompanyId))
	_ = writer.WriteField("company_whatsapp_id", fmt.Sprintf("%d", company.CompanyWhatsappId))
	_ = writer.WriteField("company_phone", companyWhatsapp)
	_ = writer.WriteField("jid", contactJid)
	_ = writer.WriteField("message", "")
	_ = writer.WriteField("source", fmt.Sprint(Source))
	_ = writer.WriteField("type", "video")
	_ = writer.WriteField("whatsapp_date", whatsappDate)
	_ = writer.WriteField("message_id", vInfo.ID)
	_ = writer.WriteField("service_date", serviceDate)
	_ = writer.WriteField("type_id", "4")
	_ = writer.WriteField("status", "5")
	_ = writer.WriteField("path", os.Getenv("GOOGLE_DRIVE_URL")+idFileDrive)
	_ = writer.WriteField("client_original_name", clientOriginalName)
	_ = writer.WriteField("url", Url)
	// _ = writer.WriteField("response_message_id", message.ContextInfo.QuotedMessageID)
	// _ = writer.WriteField("QuotedMessage", responseText)
	// _ = writer.WriteField("response_message_path", responsePath)
	// _ = writer.WriteField("response_message_type", responseType)
	// _ = writer.WriteField("response_message_duration", responseDuration)
	// _ = writer.WriteField("response_message_original_name", responseOriginalName)
	// _ = writer.WriteField("response_message_jpeg_thumbnail", responseJpegThumbnail)

	err = writer.Close()
	if err != nil {
		logMessage := "Client " + companyWhatsapp + ": Is receiving an Video message from contactJid: " + contactJid + " but an error occurr creating form-data to send through the Webhook to Laravel app. Msgid is: " + msgid + ". Error was: " + err.Error()
		logg.Log("[ReceiveVideoWebhook-ReceiveVideoWebhook] ", companyWhatsapp, contactJid, logMessage, true)
		return
	}
	req, _ := http.NewRequest("POST", webhook, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Close = true
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		logMessage := "Client " + companyWhatsapp + ": Error sending VideoMessage of contactJid " + contactJid + " from Webhook to LaravelApp. Msgid is: " + msgid + ". Error was: " + err.Error()
		logg.Log("[ReceiveVideoWebhook-ReceiveVideoWebhook] ", companyWhatsapp, contactJid, logMessage, true)
		return
	} else {
		boddy, _ := io.ReadAll(resp.Body)
		serverresponse := string(boddy)
		logMessage := "Client " + companyWhatsapp + ": Sended VideoMessage of contactJid " + contactJid + " from Webhook to LaravelApp successfully.  Msgid is: " + msgid + ". Path is " + localPath + ". Serverresponse is " + serverresponse
		logg.Log("[ReceiveVideoWebhook-ReceiveVideoWebhook] ", companyWhatsapp, contactJid, logMessage, false)
	}
	defer resp.Body.Close()
}

// DownloadVideoFromWPP : saves audio into storage and retrieves filepath & error
func DownloadVideoFromWPP(vInfo types.MessageInfo, vMessage *proto.Message, companyWhatsapp string, wac *whatsmeow.Client, contactJid string) (string, string, string, string, string, error) {
	//  0. Get datas
	mimetype := *vMessage.VideoMessage.Mimetype
	ext := helpers.DocHandlerExts[mimetype]

	// 1. Capture exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Recovered from Exception when downloading an Video file from Whsatsapp. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveVideoWebhook-DownloadVideoFromWPP] ", companyWhatsapp, contactJid, logMessage, true)
		}
	}()

	// 2. Verify valid client
	if wac == nil {
		logMessage := "Client " + companyWhatsapp + ": Wac is nil when downloading an Video file from Whsatsapp."
		panic(logMessage)
	}

	// 3. Downlaod audio file from whatsapp
	data, err := wac.Download(vMessage.VideoMessage)
	if err != nil {
		logMessage := "Client " + companyWhatsapp + ": An error ocurr when downloading an Video file from Whsatsapp." + " Error was: " + err.Error()
		panic(logMessage)
	}

	// 4. Save file temporarily in local storage
	helpers.CreateDirIfNotExist(os.Getenv("MESSAJE_FILES") + "/video/" + contactJid)
	filepath := fmt.Sprintf(os.Getenv("MESSAJE_FILES")+"/video/%s/%s.%s", contactJid, vInfo.ID, ext)
	filename := fmt.Sprintf("%s.%s", vInfo.ID, ext)
	file, err := os.Create(filepath)
	if err != nil {
		defer file.Close()
		panic(err)
	}
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//Guardar en drive
	// idFolderParent, fileUploadId := helpers.SaveInGoogleDrive(filepath, mimetype, contactJid)
	idFolderParent, fileUploadId := "", ""
	// 5. print logs ans return
	logMessage := "Client " + companyWhatsapp + ": Has downloaded the Video file from Whsatsapp fom contactJid: " + contactJid + ". Filepath :" + filepath
	logg.Log("[ReceiveVideoWebhook-DownloadVideoFromWPP] ", companyWhatsapp, contactJid, logMessage, false)
	clientOriginalName := ""
	return filepath, filename, clientOriginalName, idFolderParent, fileUploadId, nil
}
