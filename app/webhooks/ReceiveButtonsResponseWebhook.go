package webhooks

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
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

func ReceiveButtonsResponseWebhook(vInfo types.MessageInfo, vMessage *proto.Message, companyWhatsapp string, wac *whatsmeow.Client) {
	// 0. Load env vars
	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/reciveTextMessage"

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
			logMessage := "Client " + companyWhatsapp + ": Exception in ReceiveButtonsResponseWebhook for " + contactJid + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveButtonsResponseWebhook-ReceiveButtonsResponseWebhook] ", companyWhatsapp, contactJid, logMessage, true)
		}
	}()

	// 3. Verify is message_id already exist in chats table
	err := models.ChatModel.AlreadyExistChat(vInfo.ID)
	if err != nil {
		logg.ErrorLogger.Println("\033[31m", err, "\033[0m")
	}

	// 4. Get text message
	text := vMessage.ButtonsResponseMessage.GetSelectedDisplayText()

	// 6. Save message into chats table
	company, _ := models.CompanyWhatsappModel.GetCompanyWhatsapp(companyWhatsapp)
	chat := migrations.Chats{
		CompanyId:         company.CompanyId,
		CompanyWhatsappId: company.CompanyWhatsappId,
		Message:           text,
		StatusId:          1,
		Status:            1,
		TypeId:            structs.TEXT_MESSAGE,
		CompanyPhone:      companyWhatsapp,
		ContactPhone:      contactJid,
		Source:            Source,
		MessageId:         msgid,
		WhatsappDate:      whatsappDate,
		ServiceDate:       serviceDate,
	}
	models.ChatModel.SaveMessageIntoChats(chat)

	// 7. Send request to Principal Laravel App
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("company_id", fmt.Sprintf("%d", company.CompanyId))
	_ = writer.WriteField("company_whatsapp_id", fmt.Sprintf("%d", company.CompanyWhatsappId))
	_ = writer.WriteField("company_phone", companyWhatsapp)
	_ = writer.WriteField("jid", contactJid)
	_ = writer.WriteField("message", text)
	_ = writer.WriteField("source", fmt.Sprint(Source))
	_ = writer.WriteField("type", "text")
	_ = writer.WriteField("whatsapp_date", whatsappDate)
	_ = writer.WriteField("message_id", vInfo.ID)
	_ = writer.WriteField("service_date", serviceDate)
	_ = writer.WriteField("type_id", "1")
	_ = writer.WriteField("status", "5")

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
		logg.Log("[ReceiveButtonsResponseWebhook-ReceiveButtonsResponseWebhook] ", companyWhatsapp, contactJid, logMessage, true)
	}
	defer resp.Body.Close()
}
