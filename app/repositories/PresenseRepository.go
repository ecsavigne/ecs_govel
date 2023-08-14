package repositories

import (
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/structs"

	"go.mau.fi/whatsmeow/types"
)

// Repositories Presence
func (pr *PresenseRepositories) SendChatPresence(CompanyPhone string, ContactPhone string, action string) structs.Response {
	ContactJID := types.JID{
		User:   ContactPhone,
		Server: types.DefaultUserServer,
	}

	var State types.ChatPresence
	var Media types.ChatPresenceMedia
	switch action {
	case "typing_start":
		State = types.ChatPresenceComposing
		Media = types.ChatPresenceMediaText
	case "typing_pause":
		State = types.ChatPresencePaused
		Media = types.ChatPresenceMediaText
	case "audio_start":
		State = types.ChatPresenceComposing
		Media = types.ChatPresenceMediaAudio
	case "audio_pause":
		State = types.ChatPresencePaused
		Media = types.ChatPresenceMediaAudio
	}

	err := structs.Connections[CompanyPhone].SendChatPresence(ContactJID, State, Media)
	if err != nil {
		panic("Error when sending a presense message. The error was: " + err.Error())
	}

	// 5. save logs
	logMessage := "Client " + CompanyPhone + ": have send presence chat succesfully."
	logg.Log("[PresenseRepository-SendChatPresence]", CompanyPhone, ContactPhone, logMessage, false)

	// return response
	return structs.Response{
		Status:  true,
		Message: "",
	}
}
