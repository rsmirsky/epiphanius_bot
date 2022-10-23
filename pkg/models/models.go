package models

import (
	"gorm.io/datatypes"
)

type Holiday struct {
	Id          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Date        datatypes.Date      `json:"date"`
}


type Users struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"name"`
	UserID   int64    `json:"userID"`
}
