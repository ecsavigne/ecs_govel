package routes

import (
	"ecs_govel/routes/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RutaTestFunc(c *gin.Context) {
	fmt.Print("RutaTestFunc")
	c.String(http.StatusOK, "RutaTestFunc")
}

func RutaTestFunc1(c *gin.Context) {
	fmt.Print("RutaTestFunc")
	// Response + :Param, ?Query, Data POST, Form POST, Multipart POST
	c.JSON(http.StatusOK, gin.H{
		"Param 1":      c.Param("name"),                     // Con parametro router.GET("/user/:id", function(){})
		"Query2(name)": c.Query("name"),                     // Con Query GET /path?id=1234&name=Manu&value
		"ValorTest1":   c.Request.FormValue("name"),         //Deprecated Desde Formulario(Form-Data)
		"T":            c.PostForm("t"),                     //Ok urlencoded form or multipart form when it exists
		"T1":           c.DefaultPostForm("Name", "asassd"), //urlencoded form or multipart form when it exists si no coloca valor por defecto
	})

	// -------Guardando binario file
	// -------Obtenerlo desde la peticion
	// _, fileHeader, err := c.Request.FormFile("file")
	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// 	return
	// }
	// --- Salva el archivo en el disco
	// if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
	// 	c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
	// 	return
	// }
}

func loadApi(g *gin.Engine) {
	// Sin parametro
<<<<<<< HEAD
	g.POST("/Test", RutaTestFunc)
	// g.PUT("/Test", RutaTestFun
	// g.PATCH("/Test", RutaTestFunc)
	// g.DELETE("/Test", RutaTestFunc)
	//Rutas con parametros
	// ejg.GET("/p/:name/:name1/Test", RutaTestFunc), g.POST("/l/:n/Test1", RutaTestFunc1)

	// AudioController:
	// g.POST("/:apihash/:phonenumber/SendAudioMessage", new(controller.AdController).SendAudioMessage)
	// ButtonController:
	// g.POST("/:apihash/:phonenumber/SendButtonMessage", new(controller.BtnController).SendButtonMessage)
	// CompaniesWhatsappsController:
	/*g.POST("/:apihash/createCompanyWhatsapp", new(controller.CompaniesWhatsappsController).CreateCompanyWhatsapp)
	g.POST("/:apihash/updateCompanyWhatsapp", new(controller.CompaniesWhatsappsController).UpdateCompanyWhatsapp)
	g.POST("/:apihash/deleteCompanyWhatsapp", new(controller.CompaniesWhatsappsController).DeleteCompanyWhatsapp)
	g.GET("/:apihash/updateWPAccount", new(controller.CompaniesWhatsappsController).UpdateWPAccount)
	g.POST("/:apihash/verifyWhatsapp", new(controller.CompaniesWhatsappsController).VerifyWhatsappA_B)
	g.POST("/:apihash/verifyWhatsappBA", new(controller.CompaniesWhatsappsController).VerifyWhatsappB_A)
	g.POST("/:apihash/deleteUnloggedSession", new(controller.CompaniesWhatsappsController).DeleteUnloggedSession)
	g.POST("/:apihash/deleteSession", new(controller.CompaniesWhatsappsController).DeleteSession)
	g.POST("/:apihash/:phonenumber/ChargeTags", new(controller.CompaniesWhatsappsController).ChargeTags)*/
	// ContactsController:
	/*g.POST("/:apihash/:phonenumber/GetContactList/:importContactsFromGroups", new(controller.ContController).GetContactList)
	g.POST("/:apihash/:phonenumber/CheckItsGroup", new(controller.ContController).CheckItsGroup)
	g.POST("/:apihash/:phonenumber/GetJoinedGroups", new(controller.ContController).GetJoinedGroups)
	g.POST("/:apihash/:phonenumber/GetContactInfo", new(controller.ContController).GetContactInfo)
	g.POST("/:apihash/:phonenumber/CheckItsWhatsapp", new(controller.ContController).CheckItsWhatsapp)
	g.POST("/:apihash/:phonenumber/CheckAreWhatsapp", new(controller.ContController).CheckAreWhatsapp)
	g.POST("/:apihash/:phonenumber/AnalizeServiceLog", new(controller.ContController).AnalizeServiceLog)
	g.POST("/:apihash/:phonenumber/ResendMessage", new(controller.ContController).ResendMessages)
	g.POST("/:apihash/:phonenumber/SharedContact", new(controller.ContController).SharedContacts)
	g.POST("/:apihash/:phonenumber/CheckItsWhatsapps", new(controller.ContController).CheckItsWhatsapps)
	g.POST("/:apihash/getAvatar/:filename", new(controller.ContController).GetAvatar)
	g.POST("/:apihash/transferChatContacts", new(controller.ContController).TransferChatContact)*/
	// DocumentController:
	// g.POST("/:apihash/:phonenumber/SendDocumentMessage", new(controller.DocController).SendDocumentMessage)
	// HistoryController:
	// g.POST("/:apihash/ComputeMessagesDelay", new(controller.HistController).ComputeMessagesDelay)
	// g.GET("/:apihash/:phonenumber/getHistoryPage", new(controller.HistController).GetHistoryPage)
	// g.GET("/:apihash/:phonenumber/getHistoryPageApi", new(controller.HistController).GetHistoryPageApi)
	// ImageController:
	// g.POST("/:apihash/:phonenumber/SendImageMessage", new(controller.ImgController).SendImageMessage)
	// LocationController:
	// g.POST("", )
	//LstController
	// g.POST("/:apihash/:phonenumber/SendListMessage", new(controller.LstController).SendListMessage)
	// MessageController:
	/*g.POST("/:apihash/:phonenumber/SendTextMessage", new(controller.MessageController).SendTextMessage)
	g.POST("/:apihash/:phonenumber/SendCommentaryMessage", new(controller.MessageController).SendCommentaryMessage)
	g.POST("/:apihash/:phonenumber/RevokeMessage/:remoteJid/:msgid", new(controller.MessageController).RevokeMessage)
	g.POST("/:apihash/getFile/:remoteJid/:path", new(controller.MessageController).GetFile)
	g.GET("/:apihash/UpdateFilePaths", new(controller.MessageController).UpdateFilePaths)
	g.POST("/:apihash/UpdateWhatsappDate", new(controller.MessageController).UpdateWhatsappDate)
	g.POST("/:apihash/DeteleChat", new(controller.MessageController).DeteleChat)
	g.POST("/:apihash/:phonenumber/EditMessage", new(controller.MessageController).EditTextMessage)*/
	//PresenseController:
	// g.POST("/:apihash/:phonenumber/SendChatPresence/:remoteJid", new(controller.PresenseController).SendChatPresence)
	// ReactionsController:
	// SessionController:
	/*g.POST("/:apihash/:phonenumber/qrcode", new(controller.SessionController).RegisterNewSession)
	g.GET("/:apihash/:phonenumber/isconnected", new(controller.SessionController).IsConnected)
	g.GET("/:apihash/:phonenumber/logout", new(controller.SessionController).Logout)
	g.GET("/:apihash/:phonenumber/reconnect", new(controller.SessionController).Reconnect)
	g.POST("/:apihash/:phonenumber/checkUpdate", new(controller.SessionController).CheckUpdate)
	g.POST("/:apihash/:phonenumber/amountWPAccounts", new(controller.SessionController).AmountWPAccounts)
	g.POST("/:apihash/:phonenumber/fetchState", new(controller.SessionController).TestState)
	g.POST("/", new(controller.SessionController).Index)*/
	// SystemController:
	/*g.POST("/:apihash/availableStorage", new(controller.SystemController).AvailableStorage)
	g.POST("/availableMainProcess", new(controller.SystemController).AvailableMainProcess)*/
	// TestsController:
	/*g.POST("/:apihash/:phonenumber/deleteChatContact", new(controller.Tests).DeleteChatContact)
	g.POST("/:apihash/Gorm", new(controller.Tests).GormConsult)
	g.POST("/deleteChatContact", new(controller.Tests).DeleteChatContact)
	g.POST("/:apihash/:phonenumber/analizarDelay", new(controller.Tests).AnalizarDelay)
	g.POST("/testButton", new(controller.Tests).TestButton)
	g.POST("/seederColor", new(controller.Tests).SeederColor)
	g.POST("/insertChat", new(controller.Tests).IdParentFolderDrivensertChat)
	g.GET("/test1", new(controller.Tests).TestRoute)*/
	// VideoController:
	/*g.POST("/:apihash/:phonenumber/SendVideoMessage", new(controller.VidController).SendVideoMessage)*/
	// UrlController:
	//g.POST("/:apihash/:phonenumber/SendUrlMessage", new(controller.UrlController).SendUrlMessage)
=======
	g.GET("/Test", middleware.CorsMiddleware(), RutaTestFunc)
>>>>>>> 1424e55af40b1cc94cb7ef395becc50dfd93e7ef
}
