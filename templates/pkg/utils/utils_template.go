package utils

import (
	"go-telebot-init/pkg/cache"
	"go-telebot-init/pkg/database"
	"go-telebot-init/pkg/database/models"
)

func GetUserByChatID(db *database.DBImpl, chatID int64) (*models.User, error) {
	c := cache.GetInstance()

	user, ok := c.Get(chatID)
	if ok {
		return user.(*models.User), nil
	}
	user, err := db.DBS.GetUserByChatID(chatID)
	if err != nil {
		return nil, err
	}
	c.Set(chatID, user)
	return user.(*models.User), err
}
