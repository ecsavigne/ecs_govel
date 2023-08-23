package controllers

type Audio struct{}
type Button struct{}
type CompaniesWhatsapps struct{}
type Document struct{}
type History struct{}
type Image struct{}
type List struct{}
type Presense struct{}
type System struct{}
type Text struct{}
type Video struct{}
type Session struct{}
type Contact struct {
	Name   string `json:"name"`
	PicURL string `json:"picurl"`
}

var AudioController = new(Audio)
var ButtonController = new(Button)
var CompaniesWhatsappsController = new(CompaniesWhatsapps)
var DocumentController = new(Document)
var HistoryController = new(History)
var ImageController = new(Image)
var ListController = new(List)
var PresenseController = new(Presense)
var SystemController = new(System)
var TextController = new(Text)
var VideoController = new(Video)
var ContactController = new(Contact)
var SessionController = new(Session)
