package migration

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

/*
external_id, sender_id, external_id
*/
type Message struct {
	ID                           uint           `gorm:"primaryKey"`
	ConversationID               uint           `gorm:"index;uniqueIndex:idx_messages_unique;"`
	TypeID                       uint           `gorm:"index"`
	StatusID                     uint           `gorm:"index"`
	UserID                       uint           `gorm:"index"`
	CreatedAt                    time.Time      `gorm:"primaryKey;uniqueIndex:idx_messages_unique"`
	UpdatedAt                    time.Time      `gorm:"index"`
	DeletedAt                    gorm.DeletedAt `gorm:"index"`
	SenderAt                     time.Time      `gorm:"index"`
	ConversationCreatedAt        time.Time      `gorm:"index"`
	SenderID                     string         `gorm:"index"`
	ExternalID                   string         `gorm:"index:idx_external_id;uniqueIndex:idx_messages_unique;"`
	TextContent                  string
	Payload                      datatypes.JSON
	ReplyToMessageExternalID     *string
	ReplyToMessageConversationID *uint
	ReplyToMessageCreatedAt      *time.Time
	ReplyToMessage               *Message  `gorm:"foreignKey:ReplyToMessageExternalID,ReplyToMessageConversationID,ReplyToMessageCreatedAt;references:ExternalID,ConversationID,CreatedAt"`
	Replies                      []Message `gorm:"foreignKey:ReplyToMessageExternalID;references:ExternalID"`
	// Conversation                 *Conversation
	// Type                         Type
	// Status                       Status
	// IGContact                    IGContact `gorm:"foreignKey:SenderID;references:IGAccountIDExternal"`
}
