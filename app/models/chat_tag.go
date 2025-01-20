package models

import (
	"fmt"
	"oficial_gin/database/migration"

	"gorm.io/gorm"
)

type ChatTag migration.ChatTag

// TagID, Arg0: TagID, Arg1: ChatID, Arg2: id of ChatTag puede o no pasarse si no se pasa gorm da un id...,
func NewChatTag(arg ...interface{}) (*ChatTag, error) {
	if len(arg) > 0 {
		if arg[0] == nil || arg[1] == nil {
			return nil, fmt.Errorf("Arg0 TagID, Arg1 ChatsID are required")
		}

		TagID, ok := arg[0].(uint)
		if !ok {
			return nil, fmt.Errorf("Arg0 TagID is not uint")
		}

		ChatsID, ok := arg[1].(uint)
		if !ok {
			return nil, fmt.Errorf("Arg1 ChatsID is not uint")
		}

		if arg[2] != nil {
			id, ok := arg[2].(uint)
			if !ok {
				return nil, fmt.Errorf("Arg2 ChatTagId is not uint")
			}
			return &ChatTag{
				TagID:  TagID,
				ChatID: ChatsID,
				ID:     id,
			}, nil
		} else {
			return &ChatTag{
				TagID:  TagID,
				ChatID: ChatsID,
			}, nil
		}
	} else {
		return &ChatTag{}, nil
	}
}

func (ct *ChatTag) Insert(db *gorm.DB) *gorm.DB {
	return db.Create(ct)
}

func (ct *ChatTag) Update(db *gorm.DB) *gorm.DB {
	return db.Save(ct)
}

// Param2: TagID, Param3: ChatID
func (ct *ChatTag) Get(db *gorm.DB, TagID, ChatsID uint) {
	db.First(ct, "tag_id = ? AND chats_id = ?", TagID, ChatsID)
}
