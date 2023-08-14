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

func (mc *System) AnalizeServiceLog(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-AnalizeServiceLog")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception in SystemController-AnalizeServiceLog. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[SystemController-AnalizeServiceLog] ", vars["CompanyPhone"], "", logMessage, true)
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
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	helpers.ValidateInitDate(r.FormValue("InitDate"))
	helpers.ValidateEndDate(r.FormValue("EndDate"))

	// 3. send message
	response := repositories.SystemConfigRepository.AnalizeServiceLog(vars["CompanyPhone"], r.FormValue("InitDate"), r.FormValue("EndDate"))

	// 4. return response
	json.NewEncoder(w).Encode(response)
}

func (system *System) AvailableStorage(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-AvailableStorage")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Exception checking Available Storage. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[SystemController-AvailableStorage] ", vars["CompanyPhone"], "", logMessage, true)
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

	// 3. send message
	response := repositories.SystemConfigRepository.AvailableStorage()

	// 4. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (system *System) AvailableWhatsmeowProcess(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-AvailableWhatsmeowProcess")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "An exception occur when checking AvailableWhatsmeowProcess. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[SystemController-AvailableWhatsmeowProcess] ", vars["CompanyPhone"], "", logMessage, true)
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

	// 3. send message
	response := repositories.SystemConfigRepository.AvailableWhatsmeowProcess()

	// 4. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *System) FreeHdSpace(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-FreeHdSpace")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "An exception occur when checking AvailableWhatsmeowProcess. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[SystemController-AvailableWhatsmeowProcess] ", vars["CompanyPhone"], "", logMessage, true)
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

	// 3. send message
	response := repositories.SystemConfigRepository.FreeHdSpace()

	// 4. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
