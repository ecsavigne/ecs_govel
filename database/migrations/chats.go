package migrations

import (
	"github.com/jinzhu/gorm"
)

type Chats struct {
	gorm.Model
	CompanyId                    uint64 `gorm:"index"`
	CompanyWhatsappId            uint64 `gorm:"index"`
	UserId                       uint64 // sender attendant id
	StatusId                     int    // SH status
	Status                       int    // WP status
	TypeId                       int    // Media type [1: text, 2: image, 3: audio, 4: video, 5: document, 8: contact]
	MessageId                    string `gorm:"UNIQUE_INDEX:compositeindex;index;not null"`
	CompanyPhone                 string `gorm:"UNIQUE_INDEX:compositeindex;index;not null"`
	ContactPhone                 string `gorm:"UNIQUE_INDEX:compositeindex;index;not null"`
	Source                       int    // [0: sended messages, 1: incomming messages]
	Message                      string
	Path                         string
	ClientOriginalName           string
	WhatsappDate                 string
	ServiceDate                  string
	ContactJson                  string // contact card message as json
	ButtonsJson                  string // buttons as json
	ListJson                     string // buttons as json
	ResponseMessageID            string
	ResponseMessageText          string
	ResponseMessagePath          string
	ResponseMessageSource        int
	ResponseMessageType          int
	ResponseMessageDuration      string
	ResponseMessageOriginalName  string
	ResponseMessageJpegThumbnail string
	IdParentFolderDrive          string
	IdFileDrive                  string
	PathDrive                    string
	Url                          string
	QuoteMsg                     *Chats `gorm:"-"`
}
