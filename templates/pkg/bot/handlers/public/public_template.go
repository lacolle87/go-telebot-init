package public

import (
	"errors"
	"go-telebot-init/pkg/database"
	"go-telebot-init/pkg/database/models"
	"go-telebot-init/pkg/utils"
	"gorm.io/gorm"
	"log/slog"
)

func HandleStart(db *database.DBImpl) func(telebot.Context) error {
	return func(c telebot.Context) error {
		_, err := utils.GetUserByChatID(db, c.Sender().ID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user := &models.User{
				ChatId:    c.Sender().ID,
				Username:  c.Sender().Username,
				FirstName: c.Sender().FirstName,
				LastName:  c.Sender().LastName,
			}
			if err = db.DBS.Create(user); err != nil {
				slog.Error("Failed to create user", "error", err)
				return err
			}
		} else if err != nil {
			slog.Error("Failed to get user", "error", err)
			return err
		}
		return c.Send("Hello, I'm a bot!")
	}
}
