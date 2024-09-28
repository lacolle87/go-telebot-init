package middlewares

import (
	"go-telebot-init/pkg/database"
	"go-telebot-init/pkg/utils"
	tele "gopkg.in/telebot.v3"
	"log/slog"
)

func IsAdmin(db *database.DBImpl) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
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
