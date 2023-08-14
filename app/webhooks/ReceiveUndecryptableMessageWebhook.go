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
	"go.mau.fi/whatsmeow/types"
)

func ReceiveUndecryptableMessageWebhook(vInfo types.MessageInfo, companyWhatsapp string, wac *whatsmeow.Client) {
	// 0. Load env vars
	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/reciveUndecryptableMessage"

	// 1. Get chat datas
	Source := 1
	msgid := vInfo.ID
	whatsappDate := vInfo.Timestamp.Format("2006-01-02 15:04:05")
	serviceDate := carbon.Now().Format("Y-m-d H:i:s")
	contactJid := strings.Replace(vInfo.Chat.User, "s.whatsapp.net", "", 1)

	// 2. Capture exception
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Exception in ReceiveUndecryptableMessageWebhook for " + contactJid + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveUndecryptableMessageWebhook-ReceiveUndecryptableMessageWebhook] ", companyWhatsapp, contactJid, logMessage, true)
			return
		}
	}()

	if contactJid == "status" {
		return
	}

	phones := []string{
		"+" + contactJid,
	}
	data, err := wac.IsOnWhatsApp(phones)
	if err != nil || !data[0].IsIn {
		panic("Number no registred in Whatsapp - Error posible: " + err.Error())
	}

	text := "Você recebeu uma mensagem deste contato que está disponível apenas no seu aplicativo do telefone.\n\nIsso acontece porque seu contato pode ter uma versão anterior da sua e algumas mensagens não são interpretaras pela nova versão."

	// 6. Save message into chats table
	company, _ := models.CompanyWhatsappModel.GetCompanyWhatsapp(companyWhatsapp)
	chat := migrations.Chats{
		CompanyId:         company.CompanyId,
		CompanyWhatsappId: company.CompanyWhatsappId,
		Message:           text,
		StatusId:          1,
		Status:            1,
		TypeId:            structs.UNDECRYPTABLE_MESSAGE,
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
	_ = writer.WriteField("type_id", "10")
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
		logg.Log("[ReceiveUndecryptableMessageWebhook-ReceiveUndecryptableMessageWebhook] ", companyWhatsapp, contactJid, logMessage, true)
	}
	defer resp.Body.Close()
}
