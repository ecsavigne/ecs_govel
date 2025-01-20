package models

import (
	"fmt"
	"oficial_gin/database/migration"

	"gorm.io/gorm"
)

type Tag migration.Tag

// Arg0: ColorId, Arg1: TagName, Arg2: id of tag puede o no pasarse si no se pasa gorm da un id...,
func NewTag(arg ...interface{}) (*Tag, error) {
	if len(arg) > 0 {
		if arg[0] == nil || arg[1] == nil {
			return nil, fmt.Errorf("Arg0 ColorId, Arg1 TagName are required")
		}

		colorId, ok := arg[0].(uint)
		if !ok {
			return nil, fmt.Errorf("Arg0 ColorId is not uint")
		}
		name, ok := arg[1].(string)
		if !ok {
			return nil, fmt.Errorf("Arg1 TagName is not string")
		}
		if arg[2] != nil {
			id, ok := arg[2].(uint)
			if !ok {
				return nil, fmt.Errorf("Arg2 TagId is not uint")
			}
			return &Tag{
				ColorID: colorId,
				Name:    name,
				ID:      id,
			}, nil
		} else {
			return &Tag{
				ColorID: colorId,
				Name:    name,
			}, nil
		}
	} else {
		return &Tag{}, nil
	}
}

func (t *Tag) Insert(db *gorm.DB) *gorm.DB {
	return db.Create(t)
}

func (t *Tag) Update(db *gorm.DB) *gorm.DB {
	return db.Create(t)
}

// Parameter: id = id of tag = *events.LabelEdit.LabelID
func (t *Tag) Get(db *gorm.DB, id uint) {
	db.First(t, "id = ?", id)
}
