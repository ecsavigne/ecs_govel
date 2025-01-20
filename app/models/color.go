package models

import (
	"fmt"
	"oficial_gin/database/migration"

	"gorm.io/gorm"
)

type Color migration.Color

// Arg0: ColorNumber, Arg1: Rgb, Arg2: HexaDecimal, Arg3: id of color puede o no pasarse si no se pasa gorm da un id...,
func NewColor(arg ...interface{}) (*Color, error) {
	if len(arg) > 0 {
		if arg[0] == nil || arg[1] == nil || arg[2] == nil {
			return nil, fmt.Errorf("Arg0 ColorNumber, Arg1 Rgb, Arg2 HexaDecimal are required")
		}

		colorNumber, ok := arg[0].(int)
		if !ok {
			return nil, fmt.Errorf("Arg0 ColorNumber is not int")
		}

		rgb, ok := arg[1].(string)
		if !ok {
			return nil, fmt.Errorf("Arg1 Rgb is not string")
		}

		hexa, ok := arg[2].(string)
		if !ok {
			return nil, fmt.Errorf("Arg2 HexaDecimal is not string")
		}

		if arg[3] != nil {
			id, ok := arg[3].(uint)
			if !ok {
				return nil, fmt.Errorf("Arg3 ColorId is not uint")
			}
			return &Color{
				ColorNumber: colorNumber,
				Rgb:         rgb,
				HexaDecimal: hexa,
				ID:          id,
			}, nil
		} else {
			return &Color{
				ColorNumber: colorNumber,
				Rgb:         rgb,
				HexaDecimal: hexa,
			}, nil
		}
	} else {
		return &Color{}, nil
	}
}

func (c *Color) Insert(db *gorm.DB) *gorm.DB {
	return db.Create(c)
}

func (c *Color) Update(db *gorm.DB) *gorm.DB {
	return db.Save(c)
}

// Param2: Colorumber = [0, ........19]
func (c *Color) Get(db *gorm.DB, colorNumber int) {
	db.First(c, "ColorNumber = ?", colorNumber)
}
