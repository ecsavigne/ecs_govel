package controllers

import (
	"github.com/gin-gonic/gin"
)

type AdController struct{}

func (mc *AdController) SendAudioMessage(g *gin.Context) {
	/*var debugMessage string = ""
	c_.GeneralLogger.Println(g.PostForm("RequestOrigin"), "-SendAudioMessage")

	logMessage := ""

	phonenumber := g.Param("phonenumber")
	remoteJid := strings.Replace(g.PostForm("RemoteJid"), "@s.whatsapp.net", "", -1)
	c_.JSONResponse = c_.GeneralResponse{}

	defer func() {
		if r := recover(); r != nil {
			logMessage = filepath.Base(os.Args[0]) + " :  Client " + phonenumber + ": Recovered from exception when sending a audio message to RemoteJid " + remoteJid + ". Interface in defer is: " + fmt.Sprintf("%+v", r) + ". DebugMessage is: " + debugMessage
			c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
			//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, true)
			c_.JSONResponse.Status = false
			c_.JSONResponse.Message = logMessage
			c_.JSONResponse.ErrorCode = -1
			c_.JSONResponse.Code = "-1"
			g.JSON(http.StatusAccepted, c_.JSONResponse)
		}
	}()

	dateFormat := time.Now()
	serviceDate := dateFormat.Format("2006-01-02 15:04:05")
	IsGroup := g.PostForm("IsGroup")
	UserId, _ := strconv.ParseUint(g.PostForm("UserId"), 10, 64)
	durationTime, errDur := strconv.ParseUint(g.PostForm("DurationTime"), 10, 32)
	CompanyId, _ := strconv.ParseUint(g.PostForm("CompanyId"), 10, 64)
	CompanyWhatsappId, _ := strconv.ParseUint(g.PostForm("CompanyWhatsappId"), 10, 64)
	Source, _ := strconv.Atoi(g.PostForm("Source"))
	viewOnce, errViewOnce := strconv.ParseBool(g.PostForm("ViewOne"))
	if errViewOnce != nil {
		viewOnce = false
	}

	previousMessage := g.PostForm("PreviousMessage")
	previousMessageId := g.PostForm("PreviousMessageId")
	previousMessageType, _ := strconv.Atoi(g.PostForm("PreviousMessageType"))
	previousMessageSource, _ := strconv.Atoi(g.PostForm("PreviousMessageSource"))
	previousMessageUrl := g.PostForm("PreviousMessageUrl")

	debugMessage = "2"

	if CompanyId == 0 || CompanyWhatsappId == 0 {
		company, err := c_.Database.GetCompanyWhatsapp(phonenumber)
		if err != nil {
			logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": CompanyWhatsap " + phonenumber + " no exist in database. Error is: " + err.Error()
			panic(logMessage)
		} else {
			CompanyId = company.CompanyId
			CompanyWhatsappId = company.CompanyWhatsappId
		}
	}

	if remoteJid == "" || phonenumber == "" {
		logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": Some of the required fields were not informed. Informed field were: (remoteJid: " + remoteJid + ", phonenumber: " + phonenumber + ")"
		c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
		//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, true)
		c_.JSONResponse.Status = false
		c_.JSONResponse.Message = logMessage
		g.JSON(http.StatusOK, c_.JSONResponse)
		return
	}
	debugMessage = "3"

	// err := whoock.VerifyClientIsLogged(phonenumber)
	apihash := g.Param("apihash")
	_, err := whook.VerifyClientIsLogged1(phonenumber, g.Request.Host, apihash, g.Param("RequestOrigin"))
	if err != nil {
		logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": Impossible reconnect client from saved client to send a audio message. Error is: " + err.Error() + ". Debug message is " + debugMessage
		c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
		//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, phonenumber, logMessage, true)

		c_.JSONExist.Status = false
		c_.JSONExist.Code = 401
		c_.JSONExist.Message = logMessage
		g.JSON(http.StatusOK, c_.JSONExist)
		return
	}
	debugMessage = "4"

	file, handler, err := g.Request.FormFile("Audio")
	nameAudio := handler.Filename
	if handler.Filename == "" {
		logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": Missing audio filename."
		c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
		//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, true)

		c_.JSONResponse.Status = false
		c_.JSONResponse.Message = logMessage
		g.JSON(http.StatusOK, c_.JSONResponse)
		return
	}
	debugMessage = "5"
	if err != nil {
		logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": Impossible get audio from FormFile. Error was: " + err.Error()
		c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
		//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, true)

		c_.JSONResponse.Status = false
		c_.JSONResponse.Message = logMessage
		g.JSON(http.StatusOK, c_.JSONResponse)
		return
	}
	defer file.Close()

	if handler.Size == 0 {
		panic(" occured one error, file with 0 Bytes detected")
	}

	path, mime, err := whoock.ReceiveAPIFile(file, handler, strings.Split(remoteJid, "@")[0])
	debugMessage = "6"

	if errDur != nil {
		if mime == "audio/mpeg" {
			durationTime = getDurationAudio(path)
		}
	}

	if err != nil {
		logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": Impossible ReceiveAPIFile. Path was: " + path + ". Mime was: " + mime + ". Error was:" + err.Error()
		c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
		//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, true)

		c_.JSONResponse.Status = false
		c_.JSONResponse.Message = logMessage
		g.JSON(http.StatusOK, c_.JSONResponse)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": Failed to ReadFile: " + fmt.Sprintf("%v", path) + ": " + fmt.Sprintf("%v", err)
		c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
		//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, true)

		c_.JSONResponse.Status = false
		c_.JSONResponse.Message = logMessage
		g.JSON(http.StatusOK, c_.JSONResponse)
		return
	}
	debugMessage = "7"

	Url := ""

	uploaded, err := whoock.Wacs[phonenumber].Upload(context.Background(), data, whatsmeow.MediaAudio)
	if err != nil {
		logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": Failed to Upload audio file to Whatsapp: " + fmt.Sprintf("%v", path) + ": " + fmt.Sprintf("%v", err) + ". DebugMessage is: " + debugMessage
		c_.ErrorLogger.Println("\033[31m[AudioController-SendAudioMessage] ", logMessage+"\033[0m")
		//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, true)

		c_.JSONResponse.Status = false
		c_.JSONResponse.Message = logMessage
		g.JSON(http.StatusOK, c_.JSONResponse)
		return
	}
	//if Source != 4 {
	Url = whoock.SendFileToStore(phonenumber, remoteJid, path, Source)
	//}
	debugMessage = "8"

	wppmsg := &waProto.Message{
		AudioMessage: &waProto.AudioMessage{
			URL:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey, // Mimetype: proto.String(http.DetectContentType(data)),
			Mimetype:      proto.String(mime),
			FileEncSHA256: uploaded.FileEncSHA256,
			FileSHA256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uploaded.FileLength), // FileLength:    proto.Uint64(uint64(len(data))),
			PTT:           proto.Bool(true),
			Seconds:       proto.Uint32(uint32(durationTime)),
			ViewOnce:      proto.Bool(viewOnce),
			// MediaKeyTimestamp: new(int64),
		}}
	if previousMessageId != "" {
		QuotedTarget := ""
		if previousMessageSource == 1 {
			QuotedTarget = remoteJid + "@s.whatsapp.net"
		} else {
			QuotedTarget = phonenumber + "@s.whatsapp.net"
		}

		wppmsg.AudioMessage.ContextInfo = &waProto.ContextInfo{
			StanzaID:      proto.String(previousMessageId),
			Participant:   proto.String(QuotedTarget),
			QuotedMessage: getProtocolMessageQuotedMessageByType(previousMessageType, previousMessageId, phonenumber, remoteJid),
		}

		wppmsg.MessageContextInfo = &waProto.MessageContextInfo{
			DeviceListMetadataVersion: proto.Int32(2),
			DeviceListMetadata:        &waProto.DeviceListMetadata{},
		}
	}

	ContactJID := types.NewJID(remoteJid, types.DefaultUserServer)
	fmt.Printf("ContactPhone: %+v\n", ContactJID)
	if IsGroup == "1" {
		re := regexp.MustCompile(`^\d+(-\d+)?$`)
		if !re.MatchString(remoteJid) {
			logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ". Error is: remoteJid " + remoteJid + " is an invalid group_id"
			c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
			//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, phonenumber, logMessage, true)

			c_.JSONExist.Status = false
			c_.JSONExist.Code = 500
			g.JSON(http.StatusOK, c_.JSONExist)
			return
		}
		ContactJID.Server = types.GroupServer
	}

	resp, err := whoock.Wacs[phonenumber].SendMessage(context.Background(), ContactJID, wppmsg, whatsmeow.SendRequestExtra{
		Timeout: time.Minute * 2,
	})
	if err != nil {
		debugMessage = "9"
		logMessage := filepath.Base(os.Args[0]) + ": Client " + phonenumber + ": Error sending audio. Error was:" + err.Error()
		c_.ErrorLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
		//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, true)

		c_.JSONResponse.Status = false
		c_.JSONResponse.Message = logMessage
		c_.JSONResponse.ErrorCode = -3
		c_.JSONResponse.Code = err.Error()
		g.JSON(http.StatusOK, c_.JSONResponse)
		return
	}
	whoock.SyncState(whoock.Wacs[phonenumber])

	whatsappDate := resp.Timestamp.Format("2006-01-02 15:04:05")
	msgid := resp.ID

	// _, _, filename, _ := whoock.StoreAPIFile2(file, handler, strings.Split(remoteJid, "@")[0], msgid, mime, path)
	_, filename, _ := whoock.GetMimeNameFile(handler, msgid)
	debugMessage = "10"

	logMessage = filepath.Base(os.Args[0]) + " :  Client " + phonenumber + ": Audio message sended successfully to RemoteJid " + remoteJid + ". Filename is: " + handler.Filename + ". Msgid is: " + msgid
	c_.GeneralLogger.Println("[AudioController-SendAudioMessage] ", logMessage)
	//c_.Database.SaveMessageIntoWhatchDog("[AudioController-SendAudioMessage]", phonenumber, remoteJid, logMessage, false)
	debugMessage = "11"

	// if Url == "" {
	// 	Dns := whoock.MyClients[phonenumber].Dns
	// 	switch Source {
	// 	case 0:
	// 		Url = "https://" + Dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + remoteJid + "/" + filename
	// 	case 2:
	// 		Url = "https://" + Dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + remoteJid + "/" + filename
	// 	// case 4:
	// 	// 	Url = "https://shippingnew.socialhub.pro/" + g.PostForm("ShippingPath")
	// 	// 	os.Remove(filepath)
	// 	case 5:
	// 		Url = "https://" + Dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + remoteJid + "/" + filename
	// 	case 6:
	// 		Url = "https://" + Dns + ".socialhub.pro/S4h_EPRZm-b46kyoUbUJ/getFile/" + remoteJid + "/" + filename
	// 	}
	// } else {
	switch Source {
	case 4:
		Url = "https://shippingnew.socialhub.pro/" + g.PostForm("ShippingPath")
	}
	//}

	chat := models.Chat{
		CompanyId:             CompanyId,
		CompanyWhatsappId:     CompanyWhatsappId,
		StatusId:              2,
		Status:                1,
		TypeId:                3,
		CompanyPhone:          phonenumber,
		ContactPhone:          remoteJid,
		Source:                Source,
		MessageId:             msgid,
		WhatsappDate:          whatsappDate,
		ServiceDate:           serviceDate,
		Seconds:               int32(durationTime),
		Path:                  filename,
		ResponseMessageID:     previousMessageId,
		ResponseMessageText:   previousMessage,
		ResponseMessageSource: previousMessageSource,
		ResponseMessageType:   previousMessageType,
		ResponseMessageUrl:    previousMessageUrl,
		ClientOriginalName:    nameAudio,
		// IdParentFolderDrive:   idFolderParent,
		// IdFileDrive:           fileUploadId,
		// ClientOriginalName: "",
		Url: Url,
		// Url: srv.BaseUrl + fileUploadId,
	}

	if UserId != 0 {
		chat.UserId = UserId
	}
	if previousMessageId != "" {
		chat.ResponseMessageID = previousMessageId
	}
	if previousMessage != "" {
		chat.ResponseMessageText = previousMessage
	}
	c_.Database.SaveMessageIntoChats(chat)
	debugMessage = "12"

	c_.JSONResponse.Status = true
	c_.JSONResponse.MsgID = msgid
	c_.JSONResponse.ErrorCode = 0
	c_.JSONResponse.Code = ""
	c_.JSONResponse.ResponseMessageText = previousMessage
	c_.JSONResponse.ResponseMessageID = previousMessageId
	c_.JSONResponse.ResponseMessageSource = previousMessageSource
	c_.JSONResponse.ResponseMessageType = previousMessageType
	c_.JSONResponse.ResponseMessageUrl = previousMessageUrl
	// c_.JSONResponse.ResponseMessageUrl = path + " <<>> Mime: " + mime
	g.JSON(http.StatusOK, c_.JSONResponse)
	debugMessage = "13"*/
}
