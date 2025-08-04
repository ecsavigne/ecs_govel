package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Crear colección Postman
func getVarParam() map[string]Body {
	fmt.Println("Mode:", string(modeDefault))
	m := make(map[string]Body)
	m["/Test"] = Body{
		modeDefault,
		[]Variable{
			{Key: "key", Value: "Prueb"},
			{Key: "key1", Value: "Prueb1"},
		},
	}
	// AudioController
	m["/apihash/phonenumber/SendAudioMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "UserId", Value: ""},
			{Key: "DurationTime", Value: "5"},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
			{Key: "Source", Value: "0"},
			{Key: "ViewOne", Value: "false"},
			{Key: "PreviousMessage", Value: ""},
			{Key: "PreviousMessageId", Value: ""},
			{Key: "PreviousMessageType", Value: ""},
			{Key: "PreviousMessageSource", Value: ""},
			{Key: "PreviousMessageUrl", Value: ""},
			{Key: "Audio", Value: "", Type: "file"},
		},
	}
	// ButtonController
	m["/apihash/phonenumber/SendButtonMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "Buttons", Value: ""},
			{Key: "Message", Value: "Test desde social Hub"},
			{Key: "Source", Value: "0"},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
		},
	}
	// CompaniesWhatsappsController:
	m["/apihash/createCompanyWhatsapp"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
			{Key: "Whatsapp", Value: ""},
		},
	}
	m["/apihash/updateCompanyWhatsapp"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
		},
	}
	m["/apihash/deleteCompanyWhatsapp"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "Whatsapp", Value: ""},
		},
	}
	m["/apihash/updateWPAccount"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "OldWhatsapp", Value: ""},
			{Key: "NewWhatsapp", Value: ""},
			{Key: "WhatsappDns", Value: ""},
		},
	}
	m["/apihash/verifyWhatsapp"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "Whatsapp", Value: ""},
		},
	}
	m["/apihash/verifyWhatsappBA"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
		},
	}
	m["/apihash/deleteSession"] = Body{
		modeDefault,
		[]Variable{
			{Key: "whastapp", Value: ""},
		},
	}

	//ContactsController:
	m["/apihash/phonenumber/GetContactList/importContactsFromGroups"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
		},
	}
	m["/apihash/phonenumber/CheckItsGroup"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
		},
	}
	m["/apihash/phonenumber/GetJoinedGroups"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
		},
	}
	m["/apihash/phonenumber/GetContactInfo"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "IsGroup", Value: ""},
		},
	}
	m["/apihash/phonenumber/CheckItsWhatsapp"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
		},
	}
	m["/apihash/phonenumber/CheckAreWhatsapp"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "Contacts", Value: ""},
		},
	}
	m["/apihash/phonenumber/AnalizeServiceLog"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "phonenumber", Value: ""},
			{Key: "dateInit", Value: ""},
			{Key: "dateEnd", Value: ""},
		},
	}
	m["/apihash/phonenumber/ResendMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJids", Value: ""},
			{Key: "MessageIds", Value: ""},
		},
	}
	m["/apihash/phonenumber/SharedContact"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "contacts", Value: ""},
			{Key: "remoteJid", Value: "5354701239"},
		},
	}
	m["/apihash/phonenumber/CheckItsWhatsapps"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "Numbers", Value: ""},
		},
	}

	// DocumentController:
	m["/apihash/phonenumber/SendDocumentMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "UserId", Value: ""},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
			{Key: "IsGroup", Value: ""},
			{Key: "Message", Value: "Test desde social Hub"},
			{Key: "Source", Value: "0"},
			{Key: "PreviousMessage", Value: ""},
			{Key: "PreviousMessageId", Value: ""},
			{Key: "PreviousMessageType", Value: ""},
			{Key: "PreviousMessageSource", Value: ""},
			{Key: "PreviousMessageUrl", Value: ""},
			{Key: "Document", Value: "", Type: "file"},
		},
	}
	// HistoryController:
	m["/apihash/phonenumber/getHistoryPage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
			{Key: "PerPage", Value: ""},
			{Key: "Page", Value: ""},
		},
	}
	m["/apihash/phonenumber/getHistoryPageApi"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "CompanyId", Value: "1"},
			{Key: "PerPage", Value: ""},
			{Key: "Page", Value: ""},
			{Key: "FirstMessageId", Value: ""},
		},
	}
	// ImageController:
	m["/apihash/phonenumber/SendImageMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "IsGroup", Value: ""},
			{Key: "Message", Value: "Test desde social Hub"},
			{Key: "UserId", Value: ""},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
			{Key: "Source", Value: "0"},
			{Key: "PreviousMessageType", Value: ""},
			{Key: "PreviousMessageSource", Value: ""},
			{Key: "PreviousMessageId", Value: ""},
			{Key: "PreviousMessageUrl", Value: ""},
			{Key: "PreviousMessage", Value: ""},
			{Key: "ViewOne", Value: "false"},
			{Key: "Image", Value: "", Type: "file"},
		},
	}

	// ListController:
	m["/apihash/phonenumber/SendListMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "Message", Value: "Test desde social Hub"},
			{Key: "Lists", Value: ""},
			{Key: "Source", Value: "0"},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
		},
	}
	// MessageController:
	m["/apihash/phonenumber/SendTextMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "Message", Value: "Test desde social Hub"},
			{Key: "IsGroup", Value: ""},
			{Key: "UserId", Value: ""},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
			{Key: "Source", Value: "0"},
			{Key: "PreviousMessageType", Value: ""},
			{Key: "PreviousMessageSource", Value: ""},
			{Key: "PreviousMessageId", Value: ""},
			{Key: "PreviousMessageUrl", Value: ""},
			{Key: "PreviousMessage", Value: ""},
		},
	}
	m["/apihash/phonenumber/SendCommentaryMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "Message", Value: "Test desde social Hub"},
			{Key: "TypeId", Value: ""},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
			{Key: "UserId", Value: ""},
		},
	}
	m["/apihash/phonenumber/EditMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "RequestOrigin-SISTEMA_ONLINE-"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "MsgId", Value: ""},
			{Key: "Message", Value: "Test desde social Hub"},
		},
	}
	//PresenseController:
	m["/apihash/phonenumber/SendChatPresence/remoteJid"] = Body{
		modeDefault,
		[]Variable{
			{Key: "Action", Value: ""},
		},
	}
	// SessionController:
	m["/apihash/phonenumber/qrcode"] = Body{
		modeDefault,
		[]Variable{
			{Key: "webhook", Value: ""},
			{Key: "os", Value: "Ubuntu"},
			{Key: "platform", Value: "FireFox"},
		},
	}
	m["/apihash/phonenumber/isconnected"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: "PostMan"},
		},
	}
	// TestsController:
	m["/apihash/phonenumber/deleteChatContact"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RemoteJid", Value: "5354701239"},
		},
	}
	m["/deleteChatContact"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RemoteJid", Value: "b5354701239"},
		},
	}
	// ---> // Hasta aqui seguir hacia arriba
	m["/apihash/phonenumber/analizarDelay"] = Body{
		modeDefault,
		[]Variable{
			{Key: "CantMsgAnalizar", Value: "0"},
		},
	}
	m["/testButton"] = Body{
		modeDefault,
		[]Variable{
			{Key: "CompanyPhone", Value: "Prueb"},
			{Key: "ContactPhone", Value: "Prueb1"},
		},
	}
	m["/seederColor"] = Body{
		modeDefault,
		[]Variable{
			{Key: "rgb", Value: ""},
			{Key: "hexa", Value: ""},
		},
	}

	m["/Insert"] = Body{
		modeDefault,
		[]Variable{
			{Key: "company_phone", Value: ""},
			{Key: "created_at", Value: ""},
			{Key: "updated_at", Value: ""},
			{Key: "company_id", Value: ""},
			{Key: "company_whatsapp_id", Value: ""},
			{Key: "user_id", Value: ""},
			{Key: "status_id", Value: ""},
			{Key: "status", Value: ""},
			{Key: "type_id", Value: ""},
			{Key: "message_id", Value: ""},
			{Key: "contact_phone", Value: ""},
			{Key: "source", Value: "0"},
			{Key: "message", Value: "Test desde social Hub"},
			{Key: "path", Value: ""},
			{Key: "client_original_name", Value: ""},
			{Key: "whatsapp_date", Value: ""},
			{Key: "service_date", Value: ""},
			{Key: "contact_json", Value: ""},
			{Key: "buttons_json", Value: ""},
			{Key: "list_json", Value: ""},
			{Key: "id_parent_folder_drive", Value: ""},
			{Key: "id_file_drive", Value: ""},
			{Key: "url", Value: ""},
			{Key: "seconds", Value: ""},
			{Key: "drive_id", Value: ""},
			{Key: "response_message_id", Value: ""},
			{Key: "response_message_text", Value: ""},
			{Key: "response_message_path", Value: ""},
			{Key: "response_message_source", Value: ""},
			{Key: "response_message_type", Value: ""},
			{Key: "response_message_seconds", Value: ""},
			{Key: "response_message_original_name", Value: ""},
			{Key: "response_message_jpeg_thumbnail", Value: ""},
			{Key: "response_message_url", Value: ""},
			{Key: "resended", Value: ""},
			{Key: "group_sender_data", Value: ""},
			{Key: "location", Value: ""},
		},
	}
	// VideoController:
	m["/apihash/phonenumber/SendVideoMessage"] = Body{
		modeDefault,
		[]Variable{
			{Key: "RequestOrigin", Value: ""},
			{Key: "Video", Value: "", Type: "file"},
			{Key: "RemoteJid", Value: "5354701239"},
			{Key: "IsGroup", Value: ""},
			{Key: "Message", Value: "Test desde social Hub"},
			{Key: "UserId", Value: ""},
			{Key: "CompanyId", Value: "1"},
			{Key: "CompanyWhatsappId", Value: ""},
			{Key: "Source", Value: "0"},
			{Key: "ViewOne", Value: "false"},
			{Key: "PreviousMessageType", Value: ""},
			{Key: "PreviousMessageSource", Value: ""},
			{Key: "PreviousMessageId", Value: ""},
			{Key: "PreviousMessageUrl", Value: ""},
			{Key: "PreviousMessage", Value: ""},
			{Key: "Response", Value: ""},
		},
	}

	return m
}

func F_prepare_collection_postman(host string) {
	fmt.Println("Cantidad de Routas", len(engine.Routes()))
	itemsReqPost := []PostmanItem{}
	itemsReqGet := []PostmanItem{}
	itemsReqPut := []PostmanItem{}
	itemsReqDelete := []PostmanItem{}
	itemsReqPatch := []PostmanItem{}

	Host := host
	path := ""
	if Host == "" {
		Host = HTTP_SERVER_HOST
	}

	m := getVarParam()

	for _, route := range engine.Routes() {
		path = strings.ReplaceAll(route.Path, "/:", "/")
		base := filepath.Base(path)
		body := Body{}
		if _, existe := m[path]; existe {
			body = m[path]
		}
		itemPostman := PostmanItem{
			strings.ToTitle(base),
			Request{
				Method: route.Method,
				URL: URL{
					Raw:  fmt.Sprintf("%s:%s%s", Host, HTTP_SERVER_PORT, path),
					Host: []string{fmt.Sprintf("%s:%s", Host, HTTP_SERVER_PORT)},
					Path: []string{path},
				},
				Body: &body,
			},
		}

		switch route.Method {
		case "POST":
			itemsReqPost = append(itemsReqPost, itemPostman)
		case "GET":
			itemsReqGet = append(itemsReqGet, itemPostman)
		case "PUT":
			itemsReqPut = append(itemsReqPut, itemPostman)
		case "DELETE":
			itemsReqDelete = append(itemsReqDelete, itemPostman)
		case "PATCH":
			itemsReqPatch = append(itemsReqPatch, itemPostman)
		}
	}

	folderColletionReqPost := ColletionFolder{"Post", itemsReqPost}
	folderColletionReqGet := ColletionFolder{"Get", itemsReqGet}
	folderColletionReqPut := ColletionFolder{"Put", itemsReqPut}
	folderColletionReqDel := ColletionFolder{"Delete", itemsReqDelete}
	folderColletionReqPatch := ColletionFolder{"Patch", itemsReqPatch}
	folders := []ColletionFolder{
		folderColletionReqPost, folderColletionReqGet, folderColletionReqPut,
		folderColletionReqDel, folderColletionReqPatch,
	}

	postman := PostmanCollection{
		Info{
			Name:   "Gin Coleçâo Serviço",
			Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		folders,
	}

	file, _ := json.MarshalIndent(postman, "", " ")
	_ = os.WriteFile("collection.json", file, 0644)
}
