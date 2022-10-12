package models

import (
	"gorm.io/datatypes"
)

type Holiday struct {
	Id          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Date        datatypes.Date      `json:"date"`//datatypes.Date
}

// type StoredMessage struct {
// 	MessageID int   `sql:"message_id" json:"message_id"`
// 	ChatID    int64 `sql:"chat_id" json:"chat_id"`
// }

type Users struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"name"`
	UserID   int64    `json:"userID"`
}
