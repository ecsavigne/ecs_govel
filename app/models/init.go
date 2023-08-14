package models

import "new_whatsmeow/config"

type ApplicationModels struct {
	*config.DbInstance
}

type ChatModels struct {
	*config.DbInstance
}

type WhatchDogModels struct {
	*config.DbInstance
}

type WPAccountModels struct {
	*config.DbInstance
}

type CompanyWhatsappModels struct {
	*config.DbInstance
}

var ApplicationModel = new(ApplicationModels)
var ChatModel = new(ChatModels)
var WhatchDogModel = new(WhatchDogModels)
var WPAccountModel = new(WPAccountModels)
var CompanyWhatsappModel = new(CompanyWhatsappModels)

func init() {
	ApplicationModel.DbInstance = new(config.DbInstance)
	ChatModel.DbInstance = new(config.DbInstance)
	WhatchDogModel.DbInstance = new(config.DbInstance)
	WPAccountModel.DbInstance = new(config.DbInstance)
	CompanyWhatsappModel.DbInstance = new(config.DbInstance)

	WhatchDogModel.DbInstance = config.Database
	ApplicationModel.DbInstance = config.Database
	ChatModel.DbInstance = config.Database
	WPAccountModel.DbInstance = config.Database
	CompanyWhatsappModel.DbInstance = config.Database
}
