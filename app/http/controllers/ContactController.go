package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/repositories"
	"new_whatsmeow/app/structs"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"go.mau.fi/whatsmeow/types"
)

func (c *Contact) ContactInfo(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-ContactInfo")
	debugMessage := "1"
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Recovered from ecxeption in WhatsappController-GetContactInfo. ContactPhone " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err) + ". DebugMessage is: " + debugMessage
			logg.Log("[WhatsappController-CheckItsWhatsapp] ", vars["CompanyPhone"], "", logMessage, true)
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
	helpers.ValidateContactPhone(r.FormValue("RemoteJid"))
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])

	// 3. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 4. send message
	response := repositories.ContactRepository.ContactInfo(vars["CompanyPhone"], r.FormValue("RemoteJid"))

	// 5. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *Contact) CheckItsWhatsapp(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-CheckItsWhatsapp")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Recovered from ecxeption in WhatsappController-CheckItsWhatsapp. ContactPhone " + r.FormValue("RemoteJid") + ". Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[WhatsappController-CheckItsWhatsapp] ", vars["CompanyPhone"], "", logMessage, true)
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
	helpers.ValidateContactPhone(r.FormValue("RemoteJid"))
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])

	// 3. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 4. send message
	response := repositories.ContactRepository.CheckItsWhatsapp(vars["CompanyPhone"], r.FormValue("RemoteJid"))

	// 5. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *Contact) CheckAreWhatsapp(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-CheckAreWhatsapp")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Recovered from ecxeption in WhatsappController-CheckAreWhatsapp. Interface in defer is: " + fmt.Sprintf("%+v", err)
			logg.Log("[WhatsappController-CheckAreWhatsapp] ", vars["CompanyPhone"], "", logMessage, true)
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
	helpers.ValidateContactPhone(r.FormValue("Contacts"))
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])

	// 3. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])

	// 4. send message
	response := repositories.ContactRepository.CheckAreWhatsapp(vars["CompanyPhone"], r.FormValue("Contacts"))

	// 5. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *Contact) GetContactList(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println("\033[36m"+r.FormValue("RequestOrigin"), "-GetContactList\033[0m")
	debugMessage := "1"
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	// 1. catch exceptions
	defer func() {
		if err := recover(); err != nil {
			logMessage := "Client " + vars["CompanyPhone"] + ": Recovered from ecxeption in WhatsappController-GetContactList. Interface in defer is: " + fmt.Sprintf("%+v", err) + ". DebugMessage: " + debugMessage
			logg.Log("[WhatsappController-GetContactList] ", vars["RemoteJid"], "", logMessage, true)
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
	}()
	defer r.Body.Close()

	// 2. params validation
	helpers.ValidateApiHash(vars["ApiHash"])
	debugMessage = "2"
	helpers.ValidateCompanyPhone(vars["CompanyPhone"])
	debugMessage = "2"

	// 3. verify if client is logged in Whatsapp
	repositories.SessionRepository.IsLogged(vars["CompanyPhone"])
	debugMessage = "3"

	// 4. send message
	response := repositories.ContactRepository.GetContactList(vars["CompanyPhone"])
	debugMessage = "4"

	// 5. http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *Contact) CheckItsGroup(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-CheckItsGroup")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	phonenumber := vars["phonenumber"]
	remoteJid := r.FormValue("RemoteJid")

	defer func() {
		remoteJid := strings.Replace(r.FormValue("RemoteJid"), "@s.whatsapp.net", "", -1)
		if r := recover(); r != nil {
			logMessage := "Client " + phonenumber + ": Recovered from ecxeption when CheckItsGroup for RemoteJid " + remoteJid + ". Interface in defer is: " + fmt.Sprintf("%+v", r)
			logg.Log("[WhatsappController-CheckItsGroup] ", vars["RemoteJid"], "", logMessage, true)

			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(structs.Response{
				Status:  false,
				Code:    401,
				Message: logMessage,
			})
		}
	}()

	defer r.Body.Close()

	helpers.ValidateCompanyPhone(r.FormValue("phonenumber"))
	helpers.ValidateContactPhone(r.FormValue("RemoteJid"))

	err := repositories.SessionRepository.VerifyClientIsLogged(phonenumber)
	if err != nil {
		logMessage := "Client " + phonenumber + ": Impossible reconnect client from saved client to CheckItsGroup. Error is: " + err.Error()
		panic(logMessage)
	}

	re := regexp.MustCompile(`^\d+(-\d+)?$`)
	if !re.MatchString(remoteJid) {
		logMessage := "Client " + phonenumber + ": Impossible reconnect client from saved client to CheckItsGroup. Error is: remoteJid is an invalid group_id"
		panic(logMessage)
	}

	cnt, err := structs.Connections[phonenumber].GetGroupInfo(types.JID{User: remoteJid, Server: types.GroupServer})

	if err != nil {
		logMessage := "Client " + phonenumber + ": Impossible check if Exist the remoteJid " + remoteJid + ". Error is: " + err.Error()
		panic(logMessage)
	}

	if cnt != nil {
		logMessage := "Client " + phonenumber + ":  RemoteJid " + remoteJid + " is a VALID group id."
		logg.Log("WhatsappController-CheckItsGroup] ", vars["RemoteJid"], "", logMessage, false)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(structs.Response{
			Status:  true,
			Code:    200,
			Message: logMessage,
		})

	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(structs.Response{
			Status:  false,
			Code:    200,
			Message: "No is group",
		})
	}
}

func (c *Contact) GetJoinedGroups(w http.ResponseWriter, r *http.Request) {
	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-GetJoinedGroups")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	phonenumber := vars["phonenumber"]

	defer func() {
		if r := recover(); r != nil {
			logMessage := "Client " + phonenumber + ": Recovered from ecxeption when GetContactList. Interface in defer is: " + fmt.Sprintf("%+v", r)
			logg.Log("[ContactController-GetJoinedGroups] ", "", "", logMessage, true)
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(structs.Response{
				Name:   "",
				PicURL: "",
			})
		}
	}()
	defer r.Body.Close()

	helpers.ValidateCompanyPhone(phonenumber)
	phonenumber = strings.Replace(phonenumber, "@s.whatsapp.net", "", -1)

	groupList, err := structs.Connections[phonenumber].GetJoinedGroups()
	if err != nil {
		panic("Error obtain group")
	}

	groups := make([]structs.Groups, 0)
	for _, group := range groupList {
		c := structs.Groups{
			JID:  group.JID.User,
			Name: group.GroupName.Name}
		groups = append(groups, c)
	}
	logMessage := "Client " + phonenumber + ": Contacts info profiles obtained successfully"
	logg.GeneralLogger.Println("[ContactController-GetContactList] ", logMessage)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}

// func (mc *Contact) getAvatar(w http.ResponseWriter, r *http.Request) {
// 	logg.GeneralLogger.Println(r.FormValue("RequestOrigin"), "-getAvatar")
// 	vars := mux.Vars(r)
// 	filename := vars["filename"]

// 	defer func() {
// 		if err := recover(); err != nil {
// 			logMessage := "Recovered from ecxeption when downloding picImage for remoteJid " + strings.Split(filename, ".")[0] + ". Interface in defer is: " + fmt.Sprintf("%+v", r)
// 			logg.Log("[WhatsappController-GetContactList] ", vars["RemoteJid"], "", logMessage, true)
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(
// 				structs.Response{
// 					Status:     false,
// 					Message:    "",
// 					LogMessage: logMessage,
// 				},
// 			)
// 		}
// 	}()
// 	defer r.Body.Close()

// 	helpers.ValidateFilneame(vars["filename"])

// 	fileBytes, err := ioutil.ReadFile(os.Getenv("AVATAR_FILES") + "/" + filename)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/octet-stream")
// 	w.Write(fileBytes)

// 	defer r.Body.Close()
// }
