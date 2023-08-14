package webhooks

import (
	"time"

	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"

	"go.mau.fi/whatsmeow/types/events"
)

type Client structs.ClassClient

func (cl *Client) EventHandler(evt interface{}) {

	CompanyPhone := cl.CompanyPhone
	if cl.WAClient.Store.ID != nil {
		CompanyPhone = cl.WAClient.Store.ID.User
	}
	shnumber := cl.CompanyPhone

	if CompanyPhone != shnumber {
		logMessage := "Client " + shnumber + " had received an message to the similar number " + CompanyPhone
		logg.ErrorLogger.Println("\033[31m[Session-RegisterNewSession] ", logMessage, "\033[0m")
		CompanyPhone = shnumber
	}

	switch v := evt.(type) {
	case *events.Receipt:
		// if v.Type.GoString() == "events.ReceiptTypeRead" {
		// }
		// if v.Type.GoString() == "events.ReceiptTypePlayed" {  Mensaje de multimeda played
		// }
		break

	case *events.Presence:
		break

	case *events.HistorySync:
		break

	case *events.Message:
		if v.Info.MessageSource.IsGroup {
			return //IGNORE MESSAGE FROM GROUP
		}

		// if v.Info.MessageSource.IsFromMe {
		// return //IGNORE MESSAGE FROM ME
		// }

		if v.Message.ButtonsResponseMessage != nil {
			go ReceiveButtonsResponseWebhook(v.Info, v.Message, CompanyPhone, structs.Connections[CompanyPhone])

		} else if v.Message.Conversation != nil {
			ReceiveTextWebhook(v.Info, v.Message, CompanyPhone, structs.Connections[CompanyPhone])

		} else if v.Message.ExtendedTextMessage != nil {
			go ReceiveTextWebhook(v.Info, v.Message, CompanyPhone, structs.Connections[CompanyPhone])

		} else if v.Message.ImageMessage != nil {
			go ReceiveImageWebhook(v.Info, v.Message, CompanyPhone, structs.Connections[CompanyPhone], structs.Dns[CompanyPhone])

		} else if v.Message.AudioMessage != nil {
			go ReceiveAudioWebhook(v.Info, v.Message, CompanyPhone, structs.Connections[CompanyPhone], structs.Dns[CompanyPhone])

		} else if v.Message.DocumentMessage != nil {
			go ReceiveDocumentWebhook(v.Info, v.Message, CompanyPhone, structs.Connections[CompanyPhone], structs.Dns[CompanyPhone])

		} else if v.Message.ContactMessage != nil {
			go ReceiveContactWebhook(v.Info, v.Message, CompanyPhone, structs.Connections[CompanyPhone])

		} else if v.Message.VideoMessage != nil {
			go ReceiveVideoWebhook(v.Info, v.Message, CompanyPhone, structs.Connections[CompanyPhone], structs.Dns[CompanyPhone])

			// -----------------Device Sent Message------------------------
		} else if v.Message.DeviceSentMessage != nil {
			if v.Message.DeviceSentMessage.Message.Conversation != nil {
				go ReceiveTextWebhook(v.Info, v.Message.DeviceSentMessage.Message, CompanyPhone, structs.Connections[CompanyPhone])

			} else if v.Message.DeviceSentMessage.Message.ContactMessage != nil {
				go ReceiveContactWebhook(v.Info, v.Message.DeviceSentMessage.Message, CompanyPhone, structs.Connections[CompanyPhone])

			} else if v.Message.DeviceSentMessage.Message.AudioMessage != nil {
				go ReceiveAudioWebhook(v.Info, v.Message.DeviceSentMessage.Message, CompanyPhone, structs.Connections[CompanyPhone], structs.Dns[CompanyPhone])

			} else if v.Message.DeviceSentMessage.Message.ImageMessage != nil {
				go ReceiveImageWebhook(v.Info, v.Message.DeviceSentMessage.Message, CompanyPhone, structs.Connections[CompanyPhone], structs.Dns[CompanyPhone])

			} else if v.Message.DeviceSentMessage.Message.DocumentMessage != nil {
				go ReceiveDocumentWebhook(v.Info, v.Message.DeviceSentMessage.Message, CompanyPhone, structs.Connections[CompanyPhone], structs.Dns[CompanyPhone])

			} else if v.Message.DeviceSentMessage.Message.VideoMessage != nil {
				go ReceiveVideoWebhook(v.Info, v.Message.DeviceSentMessage.Message, CompanyPhone, structs.Connections[CompanyPhone], structs.Dns[CompanyPhone])
			}

		}

	case *events.QRScannedWithoutMultidevice:
		go ReciveQRScannedWithoutMultidevice(CompanyPhone, v)

	case *events.PairSuccess:
		if CompanyPhone != v.ID.User {
			structs.Connections[CompanyPhone].Disconnect()
			delete(structs.Connections, CompanyPhone)
			structs.Connections[CompanyPhone] = nil
			apihash := "S4h_EPRZm-b46kyoUbUJ"
			app, _ := models.ApplicationModel.GetApp(apihash)
			go models.WPAccountModel.DeleteWPNumberSession(CompanyPhone, app)
		}
		go RecivePairSuccess(CompanyPhone, v)

	case *events.PairError:
		go RecivePairError(CompanyPhone, v)

	case *events.TemporaryBan:
		go ReciveTemporaryBan(CompanyPhone, v)
		apihash := "S4h_EPRZm-b46kyoUbUJ"
		app, _ := models.ApplicationModel.GetApp(apihash)
		go models.WPAccountModel.DeleteWPNumberSession(CompanyPhone, app)
		structs.Connections[CompanyPhone].Logout()
		delete(structs.Connections, CompanyPhone)
		structs.Connections[CompanyPhone] = nil

	case *events.TempBanReason:
		go ReciveTempBanReason(CompanyPhone, v)

	case *events.ClientOutdated:
		logMessage := "Client " + CompanyPhone + ": Has received a ClientOutdated event"
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)

	case *events.ConnectFailure:
		logMessage := "Client " + CompanyPhone + ": Has received a ConnectFailure event"
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)

	case *events.ConnectFailureReason:
		logMessage := "Client " + CompanyPhone + ": Has received a ConnectFailureReason event. Reason was: " + v.String()
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)

	case *events.KeepAliveTimeout:
		logMessage := "Client " + CompanyPhone + ": Has received a KeepAliveTimeout event"
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)

	case *events.KeepAliveRestored:
		logMessage := "Client " + CompanyPhone + ": Has received a KeepAliveRestored event"
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)

	case *events.StreamError:
		logMessage := "Client " + CompanyPhone + ": Has received a StreamError event. Code is: " + v.Code + "Raw.XMLString is: " + v.Raw.XMLString()
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)

	case *events.UndecryptableMessage:
		go ReceiveUndecryptableMessageWebhook(v.Info, CompanyPhone, structs.Connections[CompanyPhone])

	case *events.StreamReplaced:
		logMessage := "Client " + CompanyPhone + ": Has received a StreamReplaced event. This event is emitted when the client is disconnected by another client connecting with the same keys. This can happen if you accidentally start another process with the same session or otherwise try to connect twice with the same session."
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)

	case *events.AppState:
		// logMessage := "Client " + CompanyPhone + ": Has received a AppState event. Event is: " + fmt.Sprintf("%+v", v)
		// logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, false)
		break

	case *events.LoggedOut:
		OnConnectMessage := "The event was triggered by a stream:error message"
		if v.OnConnect {
			OnConnectMessage = "The event was triggered by a connect failure message, and the reason was: " + v.Reason.String()
		}
		mandatoryLogout := ""
		if v.Reason.IsLoggedOut() {
			structs.Connections[CompanyPhone].Disconnect()
			delete(structs.Connections, CompanyPhone)
			structs.Connections[CompanyPhone] = nil
			app, _ := models.ApplicationModel.GetApp("S4h_EPRZm-b46kyoUbUJ")
			go models.WPAccountModel.DeleteWPNumberSession(CompanyPhone, app)
			mandatoryLogout = " The client device has deleted session datas."
		}
		logMessage := "Client " + CompanyPhone + ": Has received a LoggedOut event. LoggedOut was caused by: " + OnConnectMessage + mandatoryLogout
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)
		go LoggedOut(CompanyPhone)

	case *events.Disconnected:
		logMessage := "Client " + CompanyPhone + ": Has received a Disconnected event. Disconnected is emitted when the websocket is closed by the server."
		logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)
		time.Sleep(8 * time.Second)
		err := structs.Connections[CompanyPhone].Connect()
		structs.Connections[CompanyPhone].EnableAutoReconnect = true
		if err != nil {
			logMessage = "Client " + CompanyPhone + ": Was re-connected successfully after recive a Disconnected event and sleep by 8 seconds."
			logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, false)
		} else {
			logMessage = "Client " + CompanyPhone + ": Imposible re-connect after recive a Disconnected event and sleep 8 seconds."
			logg.Log("[ReceiveTextWebhook-eventHandler] ", CompanyPhone, CompanyPhone, logMessage, true)
		}
	}
}
