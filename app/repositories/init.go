package repositories

import "github.com/joho/godotenv"

type AudioRepositories struct{}
type ButtonRepositories struct{}
type CompaniesWhatsappsRepositories struct{}
type ContactRepositories struct{}
type DocumentRepositories struct{}
type HistoriesRepositories struct{}
type ImageRepositories struct{}
type ListRepositories struct{}
type PresenseRepositories struct{}
type SessionRepositories struct{}
type SystemConfigRepositories struct{}
type TextRepositories struct{}
type VideoRepositories struct{}

var AudioRepository *AudioRepositories
var ButtonRepository *ButtonRepositories
var CompaniesWhatsappRepository *CompaniesWhatsappsRepositories
var ContactRepository *ContactRepositories
var DocumentRepository *DocumentRepositories
var HistoriesRepository *HistoriesRepositories
var ImageRepository *ImageRepositories
var ListRepository *ListRepositories
var PresenseRepository *PresenseRepositories
var SessionRepository *SessionRepositories
var SystemConfigRepository *SystemConfigRepositories
var TextRepository *TextRepositories
var VideoRepository *VideoRepositories

func init() {
	godotenv.Load()
	AudioRepository = new(AudioRepositories)
	ButtonRepository = new(ButtonRepositories)
	CompaniesWhatsappRepository = new(CompaniesWhatsappsRepositories)
	ContactRepository = new(ContactRepositories)
	DocumentRepository = new(DocumentRepositories)
	HistoriesRepository = new(HistoriesRepositories)
	ImageRepository = new(ImageRepositories)
	ListRepository = new(ListRepositories)
	PresenseRepository = new(PresenseRepositories)
	SessionRepository = new(SessionRepositories)
	SystemConfigRepository = new(SystemConfigRepositories)
	TextRepository = new(TextRepositories)
	VideoRepository = new(VideoRepositories)
}
