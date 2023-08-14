package webhooks

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"

	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"

	"github.com/golang-module/carbon"
	"github.com/joho/godotenv"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func ReceiveContactWebhook(vInfo types.MessageInfo, vMessage *proto.Message, companyWhatsapp string, wac *whatsmeow.Client) {
	// 0. Load env vars
	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/reciveFileMessage"

	// 1. Get chat datas (part 1)
	contactJid := ""
	Source := 1
	msgid := vInfo.ID
	whatsappDate := vInfo.Timestamp.Format("2006-01-02 15:04:05")
	serviceDate := carbon.Now().Format("Y-m-d H:i:s")
	if vInfo.MessageSource.Sender.User != "" {
		if vInfo.IsFromMe {
			Source = 3
			contactJid = vInfo.MessageSource.Chat.User
		} else {
			contactJid = vInfo.MessageSource.Sender.User
		}
	} else {
		contactJid = strings.Replace(vInfo.Chat.User, "s.whatsapp.net", "", 1)
	}

	// 2. Capture exception
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Exception in ReceiveContactWebhook for " + contactJid + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveContactWebhook-ReceiveContactWebhook] ", companyWhatsapp, contactJid, logMessage, true)
			return
		}
	}()

	// 3. Verify is message_id already exist in chats table
	err := models.ChatModel.AlreadyExistChat(vInfo.ID)
	if err != nil {
		logg.ErrorLogger.Println("\033[31m", err, "\033[0m")
	}

	// 4. Get contact message
	cleanWaid := strings.Split(*vMessage.ContactMessage.Vcard, "\n")[4]
	cleanWaid = strings.Split(cleanWaid, ":")[1]
	rjid := regexp.MustCompile(`[^0-9]+`).ReplaceAllString(cleanWaid, "")
	contactinfo := fmt.Sprintf("Nome: %s, Número: %s", *vMessage.ContactMessage.DisplayName, rjid)

	// 5. Get chat datas (part 2)
	/*
		responseText := ""
		responsePath := ""
		responseType := "1"
		responseDuration := "0"
		responseOriginalName := "0"
		responseJpegThumbnail := ""
		responseSource := ""
		if message.ContextInfo.QuotedMessageID != "" {
			responseSource = message.ContextInfo.Participant
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

	// 6. Save message into chats table
	company, _ := models.CompanyWhatsappModel.GetCompanyWhatsapp(companyWhatsapp)
	chat := migrations.Chats{
		CompanyId:         company.CompanyId,
		CompanyWhatsappId: company.CompanyWhatsappId,
		StatusId:          1,
		Status:            1,
		TypeId:            structs.CONTACT_MESSAGE,
		CompanyPhone:      companyWhatsapp,
		ContactPhone:      contactJid,
		Source:            Source,
		MessageId:         msgid,
		WhatsappDate:      whatsappDate,
		ServiceDate:       serviceDate,
		ContactJson:       "{\"name\": \"" + *vMessage.ContactMessage.DisplayName + "\", \"whatsapp\": \"" + rjid + "\"}",
	}
	// if previousMessageId != "" {
	// 	chat.ResponseMessageID = previousMessageId
	// }
	// if previousMessage != "" {
	// 	chat.ResponseMessageText = previousMessage
	// }
	models.ChatModel.SaveMessageIntoChats(chat)

	// 7. Send request to Principal Laravel App
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("company_id", fmt.Sprintf("%d", company.CompanyId))
	_ = writer.WriteField("company_whatsapp_id", fmt.Sprintf("%d", company.CompanyWhatsappId))
	_ = writer.WriteField("company_phone", companyWhatsapp)
	_ = writer.WriteField("jid", contactJid)
	_ = writer.WriteField("message", contactinfo)
	_ = writer.WriteField("source", fmt.Sprint(Source))
	_ = writer.WriteField("type", "contact")
	_ = writer.WriteField("whatsapp_date", whatsappDate)
	_ = writer.WriteField("message_id", vInfo.ID)
	_ = writer.WriteField("service_date", serviceDate)
	_ = writer.WriteField("type_id", "8")
	_ = writer.WriteField("status", "5")
	_ = writer.WriteField("contact_json", "{\"name\": \""+*vMessage.ContactMessage.DisplayName+"\", \"whatsapp\": \""+rjid+"\"}")
	// _ = writer.WriteField("response_message_id", message.ContextInfo.QuotedMessageID)
	// _ = writer.WriteField("QuotedMessage", responseText)
	// _ = writer.WriteField("response_message_path", responsePath)
	// _ = writer.WriteField("response_message_type", responseType)
	// _ = writer.WriteField("response_message_duration", responseDuration)
	// _ = writer.WriteField("response_message_original_name", responseOriginalName)
	// _ = writer.WriteField("response_message_jpeg_thumbnail", responseJpegThumbnail)

	err = writer.Close()
	if err != nil {
		logMessage := "Client " + companyWhatsapp + ": Is receiving an Text message from contactJid: " + contactJid + " but an error occurr creating form-data to send through the Webhook to Laravel app. Msgid is: " + msgid + ". Error was: " + err.Error()
		panic(logMessage)
	}
	req, _ := http.NewRequest("POST", webhook, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Close = true
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		defer resp.Body.Close()
		logMessage := "Client " + companyWhatsapp + ": Error sending TexttMessage of contactJid " + contactJid + " from Webhook to LaravelApp. Msgid is: " + msgid + ". Error was: " + err.Error()
		panic(logMessage)
	} else {
		boddy, _ := io.ReadAll(resp.Body)
		serverresponse := string(boddy)
		logMessage := "Client " + companyWhatsapp + ": Sended TextMessage of contactJid " + contactJid + " from Webhook to LaravelApp successfully.  Msgid is: " + msgid + ". Serverresponse is " + serverresponse
		logg.Log("[ReceiveTextWebhook-ReceiveTextWebhook] ", companyWhatsapp, contactJid, logMessage, true)
	}
	defer resp.Body.Close()
}
