package models

import (
	"fmt"
	"oficial_gin/database/migration"

	"gorm.io/gorm"
)

type Contact migration.Contact

// Arg0: id of ContactTag puede o no pasarse si no se pasa gorm da un id...,
func NewContact(arg ...interface{}) (*Contact, error) {
	if len(arg) > 0 {
		// if arg[0] == nil || arg[1] == nil {
		// 	return nil, fmt.Errorf("Arg0 TagID, Arg1 ChatsID are required")
		// }

		// TagID, ok := arg[0].(uint)
		// if !ok {
		// 	return nil, fmt.Errorf("Arg0 TagID is not uint")
		// }

		// ContactID, ok := arg[1].(uint)
		// if !ok {
		// 	return nil, fmt.Errorf("Arg1 ContactID is not uint")
		// }

		if arg[0] != nil {
			id, ok := arg[2].(uint)
			if !ok {
				return nil, fmt.Errorf("Arg0 ContactID is not uint")
			}
			return &Contact{
				ID: id,
			}, nil
		} else {
			return &Contact{}, nil
		}
	} else {
		return &Contact{}, nil
	}
}

func (c *Contact) Insert(db *gorm.DB) *gorm.DB {
	return db.Create(c)
}

func (c *Contact) Update(db *gorm.DB) *gorm.DB {
	return db.Save(c)
}

// Param2: ContactID
func (c *Contact) Get(db *gorm.DB, ContactID uint) {
	db.First(c, "contact_id = ?", ContactID)
}
