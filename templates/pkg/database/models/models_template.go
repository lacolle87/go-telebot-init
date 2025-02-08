package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Username  string
	FirstName string
	LastName  string
	ChatID    int64 `gorm:"unique;not null"`
	IsAdmin   *bool `gorm:"default:false"`
}
