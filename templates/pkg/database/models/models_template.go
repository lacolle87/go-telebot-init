package models

import "time"

type User struct {
	ID        uint  `gorm:"primaryKey"`
	ChatId    int64 `gorm:"unique;not null"`
	Username  string
	FirstName string
	LastName  string
	IsAdmin   *bool     `gorm:"default:false"`
	Date      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
