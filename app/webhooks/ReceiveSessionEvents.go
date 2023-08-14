package webhooks

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/structs"

	"github.com/joho/godotenv"
	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow/types/events"
)

func ReciveHistorySync(CompanyPhone string, v *events.HistorySync) {
	/*
		fmt.Println("------------------ NEW events.HistorySync -------------------------")
		Conversations := v.Data.Conversations
		if Conversations != nil {
			if v.Data.SyncType != proto.HistorySync_FULL.Enum() {
				for _, o := range v.Data.Conversations {
					if strings.Contains(*o.Id, "@g.us") {
						fmt.Println("\t Ignore messages from group: ", *o.Id)
						return
					}
					ContactPhone := strings.Replace(*o.Id, "@s.whatsapp.net", "", -1)
					CompanyWhatsapp, _ := database.Database.GetCompanyWhatsapp(CompanyPhone)
					for _, m := range o.Messages {
						js, _ := json.Marshal(m)
						fmt.Printf("\n\t\t\t ************ A %s messages: %s\n", v.Data.SyncType, js)
						chat := util.GetChatFromHistoryMessage(m, ContactPhone, CompanyPhone, CompanyWhatsapp)
						database.Database.SaveMessageIntoChats(chat)
					}
				}
			}
		} else {
			fmt.Printf("\t\t\tEvent message was: %+v\n", v)
		}
	*/
}

func ReciveQRScannedWithoutMultidevice(CompanyPhone string, v *events.QRScannedWithoutMultidevice) {
	defer func() {
		companyWhatsapp := CompanyPhone
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Recovered from Exception when handle QRScannedWithoutMultidevice event. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveSessionEvents-ReciveQRScannedWithoutMultidevice] ", CompanyPhone, CompanyPhone, logMessage, true)
		}
	}()

	logMessage := "Client " + CompanyPhone + ": Try scan qrcode with a non Multidevice phone app. Event QRScannedWithoutMultidevice."
	logg.Log("[ReceiveSessionEvents-ReciveQRScannedWithoutMultidevice] ", CompanyPhone, CompanyPhone, logMessage, false)

	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/qrScannedWithoutMultidevice"
	hookmesssage := structs.Response{
		Status:     true,
		Message:    CompanyPhone,
		LogMessage: logMessage,
	}
	jsonPayload, _ := json.Marshal(hookmesssage)

	req, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()
}

func RecivePairSuccess(CompanyPhone string, v *events.PairSuccess) {
	defer func() {
		companyWhatsapp := CompanyPhone
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Recovered from Exception when handle PairSuccess event. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveSessionEvents-RecivePairSuccess] ", CompanyPhone, CompanyPhone, logMessage, true)
		}
	}()

	logMessage := "Client " + CompanyPhone + ": Has scanned successfully the qrcode with the number: " + CompanyPhone
	logg.Log("[ReceiveSessionEvents-RecivePairSuccess] ", CompanyPhone, CompanyPhone, logMessage, false)

	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/pairSuccess"
	hookmesssage := structs.Response{
		Status:     true,
		Message:    CompanyPhone,
		LogMessage: logMessage,
	}
	jsonPayload, _ := json.Marshal(hookmesssage)

	req, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()
}

func RecivePairError(CompanyPhone string, v *events.PairError) {
	defer func() {
		companyWhatsapp := CompanyPhone
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Recovered from Exception when handle PairError event. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveSessionEvents-RecivePairError] ", CompanyPhone, CompanyPhone, logMessage, true)
		}
	}()

	logMessage := "Client " + CompanyPhone + ": An error occurr when scanning qrcode with the number: " + CompanyPhone + "Error was: " + v.Error.Error()
	logg.Log("[ReceiveSessionEvents-RecivePairError] ", CompanyPhone, CompanyPhone, logMessage, true)

	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/pairError"
	hookmesssage := structs.Response{
		Status:     true,
		Message:    CompanyPhone,
		LogMessage: logMessage,
	}
	jsonPayload, _ := json.Marshal(hookmesssage)

	req, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()
}

func ReciveTemporaryBan(CompanyPhone string, v *events.TemporaryBan) {
	defer func() {
		companyWhatsapp := CompanyPhone
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Recovered from Exception when handle TemporaryBan event. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveSessionEvents-ReciveTemporaryBan] ", CompanyPhone, CompanyPhone, logMessage, true)
		}
	}()

	logMessage := "Client " + CompanyPhone + ": A temporary BAN occurr with number: " + CompanyPhone + "Temporary BAN event was: " + fmt.Sprintf("%+v", v)
	logg.Log("[ReceiveSessionEvents-ReciveTemporaryBan] ", CompanyPhone, CompanyPhone, logMessage, true)

	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/temporaryBan"
	hookmesssage := structs.Response{
		Status:     true,
		Message:    CompanyPhone,
		LogMessage: logMessage,
	}
	jsonPayload, _ := json.Marshal(hookmesssage)

	req, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()
}

func ReciveTempBanReason(CompanyPhone string, v *events.TempBanReason) {
	defer func() {
		companyWhatsapp := CompanyPhone
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Recovered from Exception when handle TempBanReason event. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveSessionEvents-ReciveTempBanReason] ", CompanyPhone, CompanyPhone, logMessage, true)
		}
	}()

	logMessage := "Client " + CompanyPhone + ": The temporary reason with number: " + CompanyPhone + "was: " + v.String() + ". Temporary BAN event was: " + fmt.Sprintf("%+v", v)
	logg.Log("[ReceiveSessionEvents-ReciveTempBanReason] ", CompanyPhone, CompanyPhone, logMessage, true)

	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/tempBanReason"
	hookmesssage := structs.Response{
		Status:     true,
		Message:    CompanyPhone,
		LogMessage: logMessage,
	}
	jsonPayload, _ := json.Marshal(hookmesssage)

	req, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()
}

func ReciveQR(CompanyPhone string, v *events.QR) {
	defer func() {
		companyWhatsapp := CompanyPhone
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Recovered from Exception when handle QR event. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveSessionEvents-ReciveTempBanReason] ", CompanyPhone, CompanyPhone, logMessage, true)
		}
	}()

	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/qr"
	qrcodes := []string{}

	for i, evt := range v.Codes {
		png, _ := qrcode.Encode(evt, qrcode.Medium, 256)
		qrcode := base64.StdEncoding.EncodeToString(png)
		qrcodebase64 := "data:image/png;base64," + qrcode
		qrcodes[i] = qrcodebase64

		logMessage := "Client " + CompanyPhone + ": Has received a QR event. Code " + strconv.Itoa(i) + " is: " + evt
		logg.Log("[ReceiveSessionEvents-ReciveTempBanReason] ", CompanyPhone, CompanyPhone, logMessage, false)
	}

	hookmesssage := structs.Response{
		Status:     true,
		Message:    CompanyPhone,
		LogMessage: strings.Join(qrcodes, " "),
	}
	jsonPayload, _ := json.Marshal(hookmesssage)

	req, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()
}

func LoggedOut(CompanyPhone string) {
	defer func() {
		companyWhatsapp := CompanyPhone
		if err := recover(); err != nil {
			logMessage := "Client " + companyWhatsapp + ": Recovered from Exception when handle LoggedOut event. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[ReceiveSessionEvents-LoggedOut] ", CompanyPhone, CompanyPhone, logMessage, true)
		}
	}()

	godotenv.Load(".env")
	webhook := os.Getenv("BASE_URL") + "/api/RPI/loggedOut"

	hookmesssage := structs.Response{
		Status:     true,
		Message:    CompanyPhone,
		LogMessage: "LoggedOut event",
	}
	jsonPayload, _ := json.Marshal(hookmesssage)

	req, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()
}
