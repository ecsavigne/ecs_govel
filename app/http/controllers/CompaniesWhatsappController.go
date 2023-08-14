package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/repositories"
	"new_whatsmeow/app/structs"

	"github.com/gorilla/mux"
)

func (c *CompaniesWhatsapps) CreateCompanyWhatsapp(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-CreateCompanyWhatsapp")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + r.FormValue("Whatsapp") + ": Exception in CompaniesWhatsappController-CreateCompanyWhatsapp " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[CompaniesWhatsappController-CreateCompanyWhatsapp] ", r.FormValue("Whatsapp"), r.FormValue("Whatsapp"), logMessage, true)
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
	helpers.ValidateCompanyPhone(r.FormValue("Whatsapp"))
	CompanyId := helpers.ValidateCompanyId(r.FormValue("CompanyId"))
	CompanyWhatsappId := helpers.ValidateCompanyWhatsappId(r.FormValue("CompanyWhatsappId"))

	// 3. create company whatsapp
	response := repositories.CompaniesWhatsappRepository.CreateCompanyWhatsapp(r.FormValue("Whatsapp"), CompanyId, CompanyWhatsappId)

	// 4. logg messga
	logMessage := "Whatsapp " + r.FormValue("Whatsapp") + " created successfully."
	logg.Log("[CompaniesWhatsapps-CreateCompanyWhatsapp]", r.FormValue("Whatsapp"), r.FormValue("Whatsapp"), logMessage, false)

	// 5. return response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *CompaniesWhatsapps) UpdateCompanyWhatsapp(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-SendAudioMessage")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + r.FormValue("Whatsapp") + ": Exception in CompaniesWhatsappController-UpdateCompanyWhatsapp " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[CompaniesWhatsappController-UpdateCompanyWhatsapp] ", r.FormValue("Whatsapp"), r.FormValue("Whatsapp"), logMessage, true)
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
	helpers.ValidateWhatsapp(r.FormValue("Whatsapp"))
	Whatsapp := r.FormValue("Whatsapp")
	CompanyId := helpers.ValidateCompanyId(r.FormValue("CompanyId"))
	CompanyWhatsappId := helpers.ValidateCompanyWhatsappId(r.FormValue("CompanyWhatsappId"))

	// 3. update company whatsapp
	response := repositories.CompaniesWhatsappRepository.UpdateCompanyWhatsapp(Whatsapp, CompanyId, CompanyWhatsappId)

	// 4. logg messga
	logMessage := "CompanyWhatsapp " + Whatsapp + " updated successfully."
	logg.Log("[CompaniesWhatsapps-UpdateCompanyWhatsapp]", Whatsapp, Whatsapp, logMessage, false)

	// 5. return response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *CompaniesWhatsapps) DeleteCompanyWhatsapp(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println("\033[36m"+r.FormValue("RequestOrigin"), "-DeleteCompanyWhatsapp\033[0m")
	w.Header().Set("Content-Type", "application/json")

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + r.FormValue("Whatsapp") + ":  Exception in CompaniesWhatsappController-DeleteCompanyWhatsapp " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[CompaniesWhatsappController-DeleteCompanyWhatsapp] ", r.FormValue("Whatsapp"), r.FormValue("Whatsapp"), logMessage, true)
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
	helpers.ValidateWhatsapp(r.FormValue("Whatsapp"))
	Whatsapp := r.FormValue("Whatsapp")
	CompanyId := helpers.ValidateCompanyId(r.FormValue("CompanyId"))
	CompanyWhatsappId := helpers.ValidateCompanyWhatsappId(r.FormValue("CompanyWhatsappId"))
	// 3. update company whatsapp
	response := repositories.CompaniesWhatsappRepository.DeleteCompanyWhatsapp(Whatsapp, CompanyId, CompanyWhatsappId)

	// 4. logg messga
	logMessage := "CompanyWhatsapp " + Whatsapp + " deleted successfully."
	logg.Log("[CompaniesWhatsapps-DeleteCompanyWhatsapp]", Whatsapp, Whatsapp, logMessage, false)

	// 5. return response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *CompaniesWhatsapps) UpdateWPAccount(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println("\033[31m"+r.FormValue("RequestOrigin"), "-UpdateWPAccount \033[0m")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + r.FormValue("Whatsapp") + ": Exception in CompaniesWhatsappController-UpdateWPAccount " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[CompaniesWhatsappController-UpdateWPAccount] ", r.FormValue("Whatsapp"), r.FormValue("Whatsapp"), logMessage, true)
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
	helpers.ValidateWhatsapp(r.FormValue("OldWhatsapp"))
	helpers.ValidateWhatsapp(r.FormValue("NewWhatapp"))
	OldWhatsapp := r.FormValue("OldWhatsapp")
	NewWhatsapp := r.FormValue("NewWhatapp")
	helpers.ValidateWhatsappDns(r.FormValue("WhatsappDns"))
	WhatsappDns := r.FormValue("WhatsappDns")
	if r.FormValue("WhatsappDns") == "" {
		logg.ErrorLogger.Println("\033[31mMissing WhatsappDns in request\033[0m")
	} else {
		WhatsappDns = r.FormValue("WhatsappDns")
	}

	// 3. update company whatsapp
	response := repositories.CompaniesWhatsappRepository.UpdateWPAccount(OldWhatsapp, NewWhatsapp, WhatsappDns)

	// 4. logg messga
	logMessage := "CompanyWhatsapp " + NewWhatsapp + ", UpdateWPAccount updated successfully."
	logg.Log("[CompaniesWhatsapps-UpdateWPAccount]", NewWhatsapp, OldWhatsapp, logMessage, false)

	// 5. return response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *CompaniesWhatsapps) VerifyWhatsappA_B(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println("\033[36m" + r.FormValue("RequestOrigin") + "\033[0m")
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := recover(); err != nil {
			logMessage := ": Exception in CompaniesWhatsappController-VerifyWhatsappA_B --> " + fmt.Sprintf("%+v", err)
			logg.Log("[CompaniesWhatsappController-VerifyWhatsappB_A] ", "", "", logMessage, true)
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

	Whatsapp := r.FormValue("Whatsapp")
	record, _ := repositories.CompaniesWhatsappRepository.GetCompanyWhatsapp(Whatsapp)
	resp, _ := json.Marshal(&record)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		structs.Response{
			Status:  false,
			Message: string(resp),
		},
	)
}

func (c *CompaniesWhatsapps) VerifyWhatsappB_A(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println("\033[36m" + r.FormValue("RequestOrigin") + "\033[0m")
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		if err := recover(); err != nil {
			logMessage := ": Exception in CompaniesWhatsappController-VerifyWhatsappB_A --> " + fmt.Sprintf("%+v", err)
			logg.Log("[CompaniesWhatsappController-VerifyWhatsappB_A] ", "", "", logMessage, true)
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

	records := repositories.CompaniesWhatsappRepository.GetAllCompanyWhatsapp()

	for _, record := range records {
		repositories.CompaniesWhatsappRepository.Verify(record)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		structs.Response{
			Status:  true,
			Message: "Verified",
		},
	)
}
