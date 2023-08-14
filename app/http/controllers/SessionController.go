package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"

	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/repositories"
	"new_whatsmeow/app/structs"
)

func (c *Session) Qrcode(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-Qrcode")
	debugMessage := "1"

	w.Header().Set("Content-Type", "application/json")

	// 1. catch exceptions
	vars := mux.Vars(r)
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception in SessionController-Qrcode. Error was: " + fmt.Sprintf("%+v", err) + " Debug message is: " + debugMessage
			logg.Log("[SessionController-Qrcode] ", vars["CompanyPhone"], "", logMessage, true)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				structs.Response{
					Status:     false,
					Message:    "",
					LogMessage: logMessage,
				},
			)
		}
	}()
	defer r.Body.Close()

	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	if !models.ApplicationModel.AppIsAuth(vars["ApiHash"]) {
		panic("Unauthorized apihash")
	}
	helpers.ValidateWebHook(r.FormValue("webhook"))
	debugMessage = "2"
	fmt.Println("Here 4")

	// 4. get qrcode
	response := repositories.SessionRepository.Qrcode(vars["CompanyPhone"], vars["ApiHash"], r.FormValue("webhook"), r.Host)
	fmt.Println("Here 5")

	// 6. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	fmt.Println("Here 6")
}

func (c *Session) IsLogged(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-IsLogged")
	debugMessage := "1"
	w.Header().Set("Content-Type", "application/json")

	// 1. catch exceptions
	vars := mux.Vars(r)
	defer func() {
		if err := recover(); err != nil {
			app, _ := models.ApplicationModel.GetApp(vars["ApiHash"])
			go models.WPAccountModel.DeleteWPNumberSession(vars["CompanyPhone"], app)
			delete(structs.Connections, vars["CompanyPhone"])

			logMessage := "Client " + vars["CompanyPhone"] + ": an exception occur in SessionController-IsLogged. Interface in defer is: " + fmt.Sprintf("%+v", err) + " DebugMessage is: " + debugMessage
			logg.Log("[SessionController-IsLogged] ", vars["CompanyPhone"], "", logMessage, true)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				structs.Response{
					Status:     false,
					Message:    "",
					LogMessage: logMessage,
				},
			)
		}
	}()
	defer r.Body.Close()

	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	debugMessage = "1.1"
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	debugMessage = "1.2"
	if !models.WPAccountModel.WPNumberExists(vars["CompanyPhone"]) {
		panic("No session in wp_accounts table of service database with number " + vars["CompanyPhone"] + "")
	}
	debugMessage = "2"

	// 3. verify if deviceClient is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])
	debugMessage = "4"
	logMessage := "Client " + vars["CompanyPhone"] + ": Client is Connected and Logged in Whatsapp"
	logg.Log("[SessionController-Qrcode] ", vars["CompanyPhone"], "", logMessage, false)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		structs.Response{
			Status:     true,
			Message:    "",
			LogMessage: logMessage,
		},
	)
	debugMessage = "5"
}

func (c *Session) Logout(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-Logout")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": an exception occur. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[SessionController-Logout] ", vars["CompanyPhone"], "", logMessage, true)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				structs.Response{
					Status:     false,
					Message:    "",
					LogMessage: logMessage,
				},
			)
		}
	}()
	defer r.Body.Close()

	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	if !models.WPAccountModel.WPNumberExists(vars["CompanyPhone"]) {
		panic("No session in service database with number " + vars["CompanyPhone"] + "")
	}
	if structs.Connections[vars["CompanyPhone"]] == nil {
		panic("CompanyPhone is not in Connections array.")
	}

	// 3. logout
	if !structs.Connections[vars["CompanyPhone"]].IsLoggedIn() {
		app, _ := models.ApplicationModel.GetApp(vars["ApiHash"])
		go models.WPAccountModel.DeleteWPNumberSession(vars["CompanyPhone"], app)
		delete(structs.Connections, vars["CompanyPhone"])
		logMessage := "Client " + vars["CompanyPhone"] + ": Impossible reconnect client from saved client to do logout."
		logg.Log("[SessionController-Logout] ", vars["CompanyPhone"], "", logMessage, true)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			structs.Response{
				Status:     false,
				Message:    "",
				LogMessage: logMessage,
			},
		)
		return
	}

	if models.WPAccountModel.WPNumberExists(vars["CompanyPhone"]) {
		app, _ := models.ApplicationModel.GetApp(vars["ApiHash"])
		go models.WPAccountModel.DeleteWPNumberSession(vars["CompanyPhone"], app)
		structs.Connections[vars["CompanyPhone"]].Logout()
		delete(structs.Connections, vars["CompanyPhone"])
		logMessage := "Client " + vars["CompanyPhone"] + ": Successfully Logout."
		logg.Log("[SessionController-Logout] ", vars["CompanyPhone"], "", logMessage, true)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			structs.Response{
				Status:     false,
				Message:    "",
				LogMessage: logMessage,
			},
		)
		return
	}
}
