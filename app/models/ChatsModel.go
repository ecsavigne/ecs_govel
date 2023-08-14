package models

import (
	"errors"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"
)

func (m *ChatModels) Create(chat migrations.Chats) error {
	m.DB.Create(&chat)
	if m.DB.Error != nil {
		return m.DB.Error
	}
	return nil
}

func (m *ChatModels) SaveMessageIntoChats(chat migrations.Chats) error {
	if !m.DB.NewRecord(chat) {
		return m.DB.Error
	}
	m.DB.Create(&chat)
	if m.DB.Error != nil {
		return m.DB.Error
	}
	return nil
}

func (m *ChatModels) UpdateChat(chat migrations.Chats) error {
	m.DB.Model(&chat).Update("path", "")
	if m.DB.Error != nil {
		return m.DB.Error
	}
	return nil
}

func (m *ChatModels) AlreadyExistChat(messageid string) error {
	var chat migrations.Chats
	var count int
	m.DB.Where("message_id = ?", messageid).First(&chat).Count(&count)
	if count == 1 || count >= 1 {
		return errors.New("message already exist in chats")
	}
	return nil
}

type ChatWithQuote struct {
	Chat     migrations.Chats
	QuoteMsg *migrations.Chats `gorm:"-"`
}

func (m *ChatModels) LoadFromChats(CompanyPhone string, ContactPhone string, companyId string, companyWhatsappId string, perPage int64, page int64) []migrations.Chats {
	chats := make([]migrations.Chats, 0)

	ChatWithQuotes := make([]migrations.Chats, 0)
	// m.DB.Raw("SELECT * from chats where company_whatsapp_id = ? and contact_phone = ? order by whatsapp_date desc OFFSET ? limit ?", companyWhatsappId, ContactPhone, ((page - 1) * perPage), perPage).Scan(&chats)
	m.DB.Find(&chats)
	for _, chat := range chats {
		quoteMsg := new(migrations.Chats)
		if chat.ResponseMessageID != "" {
			m.DB.Where("message_id = ?", chat.ResponseMessageID).
				First(quoteMsg)
		} else {
			quoteMsg = nil
		}
		chat.QuoteMsg = quoteMsg
		ChatWithQuotes = append(ChatWithQuotes, chat)
	}
	return ChatWithQuotes
}

func (m *ChatModels) LoadFromChatsApi(phonenumber string, remoteJid string, companyId string, perPage int64, page int64, firstMessageId string) []structs.ChatsApi {
	var firstChat migrations.Chats

	m.DB.Where("company_id = ? and message_id = ?", companyId, firstMessageId).First(&firstChat)
	if firstChat.ID == 0 {
		panic("Invalid or non-existent message_id " + firstMessageId)
	}

	var chats []structs.ChatsApi
	if m.DB.Error != nil {
		return chats
	}

	m.DB.Raw("SELECT id, source, message, created_at, message_id, path, client_original_name, type_id, company_phone, contact_phone from chats where company_id = ? and contact_phone = ? and id >= ? order by whatsapp_date desc OFFSET ? limit ?", companyId, remoteJid, firstChat.ID, ((page - 1) * perPage), perPage).Scan(&chats)
	return chats
}

func (m *ChatModels) DeteleChat(message_id string, ContactPhone string) error {
	m.DB.Unscoped().Where("contact_phone = ? and message_id = ?", ContactPhone, message_id).Delete(&migrations.Chats{})
	if m.DB.Error != nil {
		return m.DB.Error
	}
	return nil
}

func (m *ChatModels) LoadOldestChats(perPage int64, page int64, date string) []migrations.Chats {
	var chats []migrations.Chats
	m.DB.Raw("SELECT * from chats where type_id > 1 and created_at < ? and created_at > '2022-05-16 00:00:00' order by created_at desc OFFSET ? limit ?", date, ((page - 1) * perPage), perPage).Scan(&chats)
	return chats
}
