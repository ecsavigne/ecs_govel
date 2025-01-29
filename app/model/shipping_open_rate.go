package model

import (
	"ecs_govel/database/migration"
	"fmt"

	"gorm.io/gorm"
)

type ShippingOpenRate migration.ShippingOpenRate

// Arg0: MessageId, Arg1: CompaingId, Arg2: Releaded, Arg3: id of ShippingOpenRate puede o no pasarse si no se pasa gorm da un id...,
func NewShippingOpenRate(arg ...interface{}) (*ShippingOpenRate, error) {
	if len(arg) > 0 {
		if arg[0] == nil || arg[1] == nil || arg[2] == nil {
			return nil, fmt.Errorf("Arg0 MessageId, Arg1 CompaingId, Arg2 Readed are required")
		}
		MessageId, ok := arg[0].(string)
		if !ok {
			return nil, fmt.Errorf("Arg0 MessageId is not String")
		}

		CompaingId, ok := arg[1].(uint64)
		if !ok {
			return nil, fmt.Errorf("Arg1 CompaingId is not uint64")
		}

		Readed, ok := arg[2].(bool)
		if !ok {
			return nil, fmt.Errorf("Arg2 Readed is not bool")
		}

		if arg[3] != nil {
			id, ok := arg[3].(uint)
			if !ok {
				return nil, fmt.Errorf("Arg3 ShippingOpenRateId is not uint")
			}

			return &ShippingOpenRate{
				MessageId:  MessageId,
				CompaingId: CompaingId,
				Readed:     Readed,
				ID:         id,
			}, nil
		} else {
			return &ShippingOpenRate{
				MessageId:  MessageId,
				CompaingId: CompaingId,
				Readed:     Readed,
			}, nil
		}
	} else {
		return &ShippingOpenRate{}, nil
	}
}

func (sor *ShippingOpenRate) Insert(db *gorm.DB) *gorm.DB {
	return db.Create(sor)
}

func (sor *ShippingOpenRate) Update(db *gorm.DB) *gorm.DB {
	return db.Save(sor)
}

// Param2: Colorumber = [0, ........19]
func (sor *ShippingOpenRate) Get(db *gorm.DB, MessageId string) {
	db.First(sor, "ColorNumber = ?", MessageId)
}
