package routes

import (
	"github.com/gorilla/mux"

	"new_whatsmeow/app/http/controllers"
)

var Router *mux.Router = new(mux.Router)

func init() {
	// messages
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/SendTextMessage", controllers.TextController.SendTextMessage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/SendImageMessage", controllers.ImageController.SendImageMessage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/SendAudioMessage", controllers.AudioController.SendAudioMessage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/SendVideoMessage", controllers.VideoController.SendVideoMessage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/SendDocumentMessage", controllers.DocumentController.SendDocumentMessage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/SendButtonMessage", controllers.ButtonController.SendButtonMessage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/SendCommentaryMessage", controllers.TextController.SendCommentaryMessage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/SendChatPresence/{RemoteJid}", controllers.PresenseController.SendChatPresence).Methods("GET", "POST")
	Router.HandleFunc("/{apihash}/getFile/{remoteJid}/{path}", controllers.TextController.GetFile).Methods("GET", "POST")

	// history
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/RevokeMessage", controllers.HistoryController.RevokeMessage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/getHistoryPage", controllers.HistoryController.GetHistoryPage).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/getHistoryPageApi", controllers.HistoryController.GetHistoryPageApi).Methods("GET", "POST")

	// contacts
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/GetContactList", controllers.ContactController.GetContactList).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/GetContactInfo", controllers.ContactController.ContactInfo).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/CheckItsWhatsapp", controllers.ContactController.CheckItsWhatsapp).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/CheckAreWhatsapp", controllers.ContactController.CheckAreWhatsapp).Methods("GET", "POST")
	Router.HandleFunc("/{apihash}/{phonenumber}/CheckItsGroup", controllers.ContactController.CheckItsGroup).Methods("GET", "POST")
	Router.HandleFunc("/{apihash}/{phonenumber}/GetJoinedGroups", controllers.ContactController.GetJoinedGroups).Methods("GET", "POST")
	//Router.HandleFunc("/{apihash}/{phonenumber}/GetJoinedGroups", controllers.ContactController.getAvatar).Methods("GET", "POST")

	// sessions
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/qrcode", controllers.SessionController.Qrcode).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/isconnected", controllers.SessionController.IsLogged).Methods("GET")
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/logout", controllers.SessionController.Logout).Methods("GET", "POST")

	// companies whatsapps
	Router.HandleFunc("/{ApiHash}/createCompanyWhatsapp", controllers.CompaniesWhatsappsController.CreateCompanyWhatsapp).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/updateCompanyWhatsapp", controllers.CompaniesWhatsappsController.UpdateCompanyWhatsapp).Methods("GET", "POST")
	Router.HandleFunc("/{ApiHash}/deleteCompanyWhatsapp", controllers.CompaniesWhatsappsController.DeleteCompanyWhatsapp).Methods("GET", "POST")
	Router.HandleFunc("/{apihash}/verifyWhatsapp", controllers.CompaniesWhatsappsController.VerifyWhatsappA_B).Methods("GET", "POST")
	Router.HandleFunc("/{apihash}/verifyWhatsappBA", controllers.CompaniesWhatsappsController.VerifyWhatsappB_A).Methods("GET", "POST")

	// system config
	Router.HandleFunc("/{ApiHash}/{CompanyPhone}/AnalizeServiceLog", controllers.SystemController.AnalizeServiceLog).Methods("GET", "POST")
	Router.HandleFunc("/{apihash}/availableStorage", controllers.SystemController.AvailableStorage).Methods("GET", "POST")
}
