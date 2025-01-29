package model

import (
	"ecs_govel/database/migration"
	"fmt"

	"gorm.io/gorm"
)

type ContactTag migration.ContactTag

// TagID, Arg0: TagID, Arg1: ContactID, Arg2: id of ContactTag puede o no pasarse si no se pasa gorm da un id...,
func NewContactTag(arg ...interface{}) (*ContactTag, error) {
	if len(arg) > 0 {
		if arg[0] == nil || arg[1] == nil {
			return nil, fmt.Errorf("Arg0 TagID, Arg1 ChatsID are required")
		}

		TagID, ok := arg[0].(uint)
		if !ok {
			return nil, fmt.Errorf("Arg0 TagID is not uint")
		}

		ContactID, ok := arg[1].(uint)
		if !ok {
			return nil, fmt.Errorf("Arg1 ContactID is not uint")
		}

		if arg[2] != nil {
			id, ok := arg[2].(uint)
			if !ok {
				return nil, fmt.Errorf("Arg2 ContactTagID is not uint")
			}
			return &ContactTag{
				TagID:     TagID,
				ContactID: ContactID,
				ID:        id,
			}, nil
		} else {
			return &ContactTag{
				TagID:     TagID,
				ContactID: ContactID,
			}, nil
		}
	} else {
		return &ContactTag{}, nil
	}
}

func (ct *ContactTag) Insert(db *gorm.DB) *gorm.DB {
	return db.Create(ct)
}

func (ct *ContactTag) Update(db *gorm.DB) *gorm.DB {
	return db.Save(ct)
}

// Param2: TagID, Param3: ChatID
func (ct *ContactTag) Get(db *gorm.DB, TagID, ContactID uint) {
	db.First(ct, "tag_id = ? AND contact_id = ?", TagID, ContactID)
}
