package middlewares

import (
	"go-telebot-init/pkg/database"
	"go-telebot-init/pkg/utils"
	"log/slog"
)

func IsAdmin(db *database.DBImpl) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			user, err := utils.GetUserByChatID(db, c.Sender().ID)
			if err != nil {
				slog.Error("Failed to get user", "error", err)
				return err
			}
			if user.IsAdmin == nil || !*user.IsAdmin {
				return nil
			}
			return next(c)
		}
	}
}
