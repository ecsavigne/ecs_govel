package repositories

import (
	"encoding/json"
	"new_whatsmeow/app/helpers"
	"new_whatsmeow/app/structs"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

type ContactData struct { //TODOColocar en el fichero de Types
	Name     string `json:"Name"`
	PicUrl   string `json:"PicUrl"`
	Verified string `json:"Verified"`
	Whatsapp string `json:"Whatsapp"`
}

// Repositories Contact
func (cr *ContactRepositories) ContactInfo(CompanyPhone string, ContactPhone string) structs.Response {
	// 1. get basic info
	Contact := types.JID{
		User:   ContactPhone,
		Server: types.DefaultUserServer,
	}
	ContactInfo, err := structs.Connections[CompanyPhone].Store.Contacts.GetContact(Contact)
	if err != nil {
		panic("Impossible get the Contacts list of CompanyPhone " + CompanyPhone + ". Error is: " + err.Error())
	}

	contactName := ""
	if ContactInfo.Found {
		if ContactInfo.FullName != "" {
			contactName = ContactInfo.FullName
		} else {
			contactName = ContactInfo.PushName
		}
	}

	// 2. get picture
	var picUrl string
	picthumb, err := structs.Connections[CompanyPhone].GetProfilePictureInfo(Contact, &whatsmeow.GetProfilePictureParams{
		Preview:     true,
		IsCommunity: false,
		ExistingID:  "",
	})
	if err != nil {
		picUrl = "https://back1.socialhub.pro/images/images.png"
	} else {
		picUrl = picthumb.URL
	}

	helpers.DownloadProfilePicture(picUrl, ContactPhone)

	// 7. return response
	c := ContactData{
		Name:   contactName,
		PicUrl: picUrl,
	}
	str, _ := json.Marshal(c)
	return structs.Response{
		Status:  true,
		Message: string(str),
	}
}

// Repositories Contact
func (cr *ContactRepositories) CheckItsWhatsapp(CompanyPhone string, ContactPhone string) structs.Response {
	phones := []string{
		"+" + ContactPhone,
	}

	data, err := structs.Connections[CompanyPhone].IsOnWhatsApp(phones)
	if err != nil {
		panic("Impossible check if Exist the ContactPhone " + ContactPhone + ". Error is: " + err.Error())
	}

	if !data[0].IsIn {
		logMessage := "Client " + CompanyPhone + ":  ContactPhone: " + ContactPhone + " is a invalid or absent whatsapp number."
		panic(logMessage)
	} else {
		return structs.Response{
			Status:  true,
			Message: strings.Replace(data[0].JID.User, "s.whatsapp.net", "", -1),
		}
	}
}

// Repositories Contact
func (cr *ContactRepositories) CheckAreWhatsapp(CompanyPhone string, Contacts string) structs.Response {
	inputContacts := make([]ContactData, 0)
	json.Unmarshal([]byte(Contacts), &inputContacts)
	phones := make([]string, 0)

	for _, contact := range inputContacts {
		contact.Whatsapp = "+" + contact.Whatsapp
		phones = append(phones, contact.Whatsapp)
	}

	data, err := structs.Connections[CompanyPhone].IsOnWhatsApp(phones)

	if err != nil {
		logMessage := "Client: " + CompanyPhone + " - ContactPhone: " + Contacts + " : Error checking a contact list. Error is: " + err.Error()
		panic(logMessage)
	}

	for i := 0; i < len(data); i++ {
		number := strings.Replace(data[i].Query, "+", "", -1)
		for j := 0; j < len(inputContacts); j++ {
			if inputContacts[j].Whatsapp == number {
				if data[i].IsIn {
					inputContacts[j].Verified = strings.Replace(data[i].JID.User, "s.whatsapp.net", "", -1)
				} else {
					inputContacts[j].Verified = ""
				}
			}
		}
	}

	resp, _ := json.Marshal(inputContacts)
	return structs.Response{
		Status:     true,
		Message:    string(resp),
		LogMessage: "",
	}
}

// Repositories Contact
func (cr *ContactRepositories) GetContactList(CompanyPhone string) structs.Response {
	contactList, _ := structs.Connections[CompanyPhone].Store.Contacts.GetAllContacts()
	var list []ContactData
	for jid, contact := range contactList {
		c := ContactData{
			Name:     contact.FullName,
			Whatsapp: strings.Replace(jid.User, "@s.whatsapp.net", "", -1),
		}
		list = append(list, c)
	}

	str, _ := json.Marshal(list)

	return structs.Response{
		Status:  true,
		Message: string(str),
	}
}
