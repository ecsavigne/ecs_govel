package migration

import (
	"time"

	"gorm.io/gorm"
)

// Chats :
type Chat struct {
	ID                           uint           `json:"id" gorm:"primarykey"`
	CreatedAt                    time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt                    time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt                    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CompanyId                    uint64         `json:"company_id" gorm:"index"`
	CompanyWhatsappId            uint64         `json:"company_whatsapp_id" gorm:"index"`
	UserId                       uint64         `json:"user_id" gorm:"index"`
	StatusId                     int            `json:"status_id" gorm:"index"`
	Status                       int            `json:"status" gorm:"index"`
	TypeId                       int            `json:"type_id" gorm:"index"`
	MessageId                    string         `json:"message_id" gorm:"UNIQUE_INDEX:compositeindex;index;not null"`
	CompanyPhone                 string         `json:"company_phone" gorm:"UNIQUE_INDEX:compositeindex;index;not null"`
	ContactPhone                 string         `json:"contact_phone" gorm:"UNIQUE_INDEX:compositeindex;index;not null"`
	Source                       int            `json:"source"` // [0: sended messages, 1: incomming messages no mios, 3: Recivido. Msg enviado para mi mismo desde, 4: sent by shipping ]
	Message                      string         `json:"message"`
	Path                         string         `json:"path"`
	ClientOriginalName           string         `json:"client_original_name"`
	WhatsappDate                 string         `json:"whatsapp_date"`
	ServiceDate                  string         `json:"service_date"`
	ContactJson                  string         `json:"contact_json"`
	ButtonsJson                  string         `json:"buttons_json"`
	ListJson                     string         `json:"list_json"`
	DriveId                      string         `json:"drive_id"`
	IdParentFolderDrive          string         `json:"id_parent_folder_drive"`
	IdFileDrive                  string         `json:"id_file_drive"`
	Url                          string         `json:"url"`
	Seconds                      int32          `json:"seconds"`
	ResponseMessageID            string         `json:"response_message_id"`
	ResponseMessageText          string         `json:"response_message"`
	ResponseMessagePath          string         `json:"response_message_path"`
	ResponseMessageSource        int            `json:"response_message_source"`
	ResponseMessageType          int            `json:"response_message_type"`
	ResponseMessageSeconds       int32          `json:"response_message_seconds"`
	ResponseMessageOriginalName  string         `json:"response_message_original_name"`
	ResponseMessageJpegThumbnail string         `json:"response_message_jpeg_thumbnail"`
	ResponseMessageUrl           string         `json:"response_message_url"`
	Resended                     bool           `json:"resended"`
	GroupSenderData              string         `json:"group_sender_data"`
	Location                     string         `json:"location"`
	ChatTags                     []ChatTag      `json:"chat_tags" gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
