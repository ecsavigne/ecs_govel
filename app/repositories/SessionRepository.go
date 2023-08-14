package repositories

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/app/webhooks"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

// Repositories Session
func (sr *SessionRepositories) Qrcode(CompanyPhone string, Apihash string, Webhook string, Host string) structs.Response {
	// 1. load env vars
	// godotenv.Load()
	// WHATSAPP_DEBUG, _ := strconv.ParseBool(os.Getenv("WHATSAPP_DEBUG"))
	WHATSAPP_DEBUG := true
	fmt.Println("Here 4.1 ", CompanyPhone)

	// 2. create and connect to sqlite3 database for this CompanyPhone
	var dbLog waLog.Logger
	if WHATSAPP_DEBUG {
		dbLog = waLog.Stdout("Database", "DEBUG", true)
	} else {
		dbLog = nil
	}
	container, err := sqlstore.New("sqlite3", "file:"+os.Getenv("SESSIONS")+"/"+CompanyPhone+".db?_foreign_keys=on", dbLog)
	if err != nil {
		panic("Error in sqlite3 database conection. Error is: " + err.Error())
	}

	// 3. get first device
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic("Error in GetFirstDevice" + err.Error())
	}
	fmt.Println("Here 4.3")

	// 4. open new wss connection to Whatsapp Server
	var clientLog waLog.Logger
	if WHATSAPP_DEBUG {
		clientLog = waLog.Stdout("Client", "DEBUG", true)
	} else {
		clientLog = nil
	}
	fmt.Println("Here 4.4")
	deviceClient := whatsmeow.NewClient(deviceStore, clientLog)
	fmt.Println("Here 4.5")

	// 5. add handler events
	classClient := new(webhooks.Client)
	classClient.WAClient = deviceClient
	classClient.CompanyPhone = CompanyPhone
	deviceClient.AddEventHandler(classClient.EventHandler)

	fmt.Println("Here 4.6")
	structs.Connections[CompanyPhone] = deviceClient
	structs.Dns[CompanyPhone] = strings.Split(Host, ".")[0]
	app, _ := models.ApplicationModel.GetApp(Apihash)
	_ = models.WPAccountModel.CreateNewWPAccount(Host, CompanyPhone, app, Webhook)

	// 6. verify is already logged in and try to connect
	fmt.Println("Here 4.7")
	if deviceClient.Store.ID != nil {
		err = deviceClient.Connect()
		deviceClient.EnableAutoReconnect = true
		if err != nil {
			panic(err.Error())
		}

		logMessage := "Already connected after try get new qrcode in RegisterNewSession"
		logg.Log("[SessionController-Qrcode] ", CompanyPhone, "", logMessage, false)

		fmt.Println("Here 4.8")
		return structs.Response{
			Status:     true,
			Message:    "Already connected",
			LogMessage: logMessage,
		}
	}

	// 7. generate new qrcode
	fmt.Println("Here 4.9")
	qrChan, _ := deviceClient.GetQRChannel(context.Background())
	err = deviceClient.Connect()
	fmt.Println("Here 4.9 ", err)
	if err != nil {
		panic("Websocket connection cannot be established")
	}
	fmt.Println("Here 4.10")
	for evt := range qrChan {
		fmt.Println("Here 4.10---")
		if evt.Event == "code" {
			png, _ := qrcode.Encode(evt.Code, qrcode.Medium, 256)
			qrcode := base64.StdEncoding.EncodeToString(png)
			qrcodebase64 := "data:image/png;base64," + qrcode

			fmt.Println("Here 4.11")
			logMessage := "Client " + CompanyPhone + ": has generated Qrcode successfully from host " + Host
			logg.Log("[SessionController-Qrcode] ", CompanyPhone, "", logMessage, false)

			fmt.Println("Here 4.12")
			return structs.Response{
				Status:       true,
				Message:      qrcodebase64,
				Qrcodebase64: qrcodebase64,
				LogMessage:   logMessage,
			}
		}
	}

	// 8. default return
	fmt.Println("Here 4.12")
	return structs.Response{
		Status:  false,
		Message: "Unable to generate qrcode",
	}
}

// Repositories Session
func (sr *SessionRepositories) RestoreAssignedClients() {
	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Exception restoring all number in the main13XXX process. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[WhatsappSessionRepository-RestoreClient] ", "", "", logMessage, true)
		}
	}()

	// 2. get the CompanyPhone in this mainX
	accounts, _ := models.WPAccountModel.GetAllWPAccounts()

	// 3. restore every CompanyPhone
	for i := range accounts {
		err := sr.RestoreClient(accounts[i].PhoneNumber, strings.Split(accounts[i].WhatsappDns, ".")[0])
		if err != nil {
			logg.Log("[WhatsappSessionRepository-RestoreClient]", accounts[i].PhoneNumber, "", "Error restoring the session of CompanyPhone ", true)
		}
		time.Sleep(120 * time.Millisecond)
	}
}

// Repositories Session
func (sr *SessionRepositories) RestoreClient(CompanyPhone string, dns string) error {
	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + CompanyPhone + ": Recovered from ecxeption in WhatsappSessionRepository-RestoreClient. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[WhatsappSessionRepository-RestoreClient] ", "", "", logMessage, true)
		}
	}()

	// 2. set debug vars
	WHATSAPP_DEBUG, _ := strconv.ParseBool(os.Getenv("WHATSAPP_DEBUG"))
	var dbLog waLog.Logger
	if WHATSAPP_DEBUG {
		dbLog = waLog.Stdout("Database", "DEBUG", true)
	} else {
		dbLog = nil
	}

	// 3. connect to sqlite3 database for this CompanyPhone
	container, err := sqlstore.New("sqlite3", "file:"+os.Getenv("SESSIONS")+"/"+CompanyPhone+".db?_foreign_keys=on", dbLog)
	if err != nil {
		logMessage := "Client " + CompanyPhone + ": Error in sqlite3 database conection. Error is: " + fmt.Sprintf("%v", err)
		panic(logMessage)
	}

	// 4. get first device
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	// 5. open new wss connection to Whatsapp Server
	var clientLog waLog.Logger
	if WHATSAPP_DEBUG {
		clientLog = waLog.Stdout("Client", "DEBUG", true)
	} else {
		clientLog = nil
	}
	deviceClient := whatsmeow.NewClient(deviceStore, clientLog)
	deviceClient.EnableAutoReconnect = true

	// 6. add handler events
	structs.Dns[CompanyPhone] = dns
	structs.Connections[CompanyPhone] = deviceClient
	classClient := new(webhooks.Client)
	classClient.WAClient = deviceClient
	classClient.CompanyPhone = CompanyPhone
	deviceClient.AddEventHandler(classClient.EventHandler)

	// 7. verify if exist a stored session
	if deviceClient.Store.ID == nil {
		logMessage := "Client " + CompanyPhone + ": Fail restoring client connection. That is: client.Store.ID is nil"
		logg.Log("[WhatsappSessionRepository-RestoreClient] ", "", "", logMessage, true)

		structs.Connections[CompanyPhone].Disconnect()
		delete(structs.Connections, CompanyPhone)

		app, _ := models.ApplicationModel.GetApp("S4h_EPRZm-b46kyoUbUJ")
		go models.WPAccountModel.DeleteWPNumberSession(CompanyPhone, app)
		panic("client.Store.ID is nil")
	}

	// 8. try connect to Whatsapp Severs
	err = structs.Connections[CompanyPhone].Connect()
	if err != nil {
		logMessage := "Client " + CompanyPhone + ": Fail restoring client connection. Error is" + fmt.Sprintf("%v", err)
		logg.Log("[WhatsappSessionRepository-RestoreClient] ", "", "", logMessage, true)

		structs.Connections[CompanyPhone].Disconnect()
		delete(structs.Connections, CompanyPhone)

		app, _ := models.ApplicationModel.GetApp("S4h_EPRZm-b46kyoUbUJ")
		go models.WPAccountModel.DeleteWPNumberSession(CompanyPhone, app)

		panic(err)
	}

	// 9. connection restored successfully
	logMessage := "Client " + CompanyPhone + ": client connection was restored successfully."
	logg.Log("[WhatsappSessionRepository-RestoreClient] ", "", "", logMessage, false)
	return nil
}

// Repositories Session
func (sr *SessionRepositories) IsLogged(CompanyPhone string) {
	if structs.Connections[CompanyPhone] == nil {
		panic("CompanyPhone is not in Connections array.")
	}
	if !structs.Connections[CompanyPhone].IsLoggedIn() {
		panic("Device is not logged on Whatsapp.")
	}
}

func (sr *SessionRepositories) VerifyClientIsLogged(phonenumber string) error {
	if structs.Connections[phonenumber] == nil {
		logMessage := "Client " + phonenumber + ": Client not LoggedIn because no is in Wacs array."
		logg.GeneralLogger.Println("[SessionUtil-VerifyClientIsLogged] ", logMessage)

		structs.Connections[phonenumber].Disconnect()
		delete(structs.Connections, phonenumber)
		structs.Connections[phonenumber] = nil
		app, _ := models.ApplicationModel.GetApp("S4h_EPRZm-b46kyoUbUJ")
		go models.WPAccountModel.DeleteWPNumberSession(phonenumber, app)

		return errors.New("client not loggedin: phonenumber is not in structs.Connections array")
	}

	if structs.Connections[phonenumber].IsLoggedIn() {
		logMessage := "Client " + phonenumber + ": IsLoggedIn ok."
		logg.GeneralLogger.Println("[SessionUtil-VerifyClientIsLogged] ", logMessage)
		return nil
	} else {
		logMessage := "Client " + phonenumber + ": Client not LoggedIn."
		logg.GeneralLogger.Println("[SessionUtil-VerifyClientIsLogged] ", logMessage)

		structs.Connections[phonenumber].Disconnect()
		delete(structs.Connections, phonenumber)
		structs.Connections[phonenumber] = nil
		app, _ := models.ApplicationModel.GetApp("S4h_EPRZm-b46kyoUbUJ")
		go models.WPAccountModel.DeleteWPNumberSession(phonenumber, app)

		return errors.New("client not loggedin")
	}
}
