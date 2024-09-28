package admin

import (
	"fmt"
	"go-telebot-init/pkg/bot/fsm"
	"go-telebot-init/pkg/database"
	"go-telebot-init/pkg/database/models"
	tele "gopkg.in/telebot.v3"
	"log/slog"
)

func HandlePost(fsmBot *fsm.FSM) func(c tele.Context) error {
	return func(c tele.Context) error {
		chatID := c.Chat().ID
		currentState := fsmBot.GetState(chatID)
		if currentState == fsm.Idle || currentState == "" {
			fsmBot.SetState(chatID, fsm.WaitingForContent)
			return c.Send("Please send the post content (text or image)")
		}
		return nil
	}
}

func handleContent(db *database.DBImpl, fsmBot *fsm.FSM, c tele.Context, content string, photo *tele.Photo) error {
	chatID := c.Chat().ID
	if fsmBot.GetState(chatID) == fsm.WaitingForContent {
		fsmBot.ClearState(chatID)

		var users *[]models.User
		err := db.DBS.GetAll(&users)
		if err != nil {
			slog.Error("Failed to retrieve user chat IDs", "error", err)
			return c.Send("Failed to retrieve user chat IDs")
		}

		var userChatIDs []int64
		for _, user := range *users {
			userChatIDs = append(userChatIDs, user.ChatId)
		}

		return broadcastPost(c, userChatIDs, content, photo)
	}
	return nil
}

func HandleText(db *database.DBImpl, fsmBot *fsm.FSM) func(c tele.Context) error {
	return func(c tele.Context) error {
		return handleContent(db, fsmBot, c, c.Message().Text, nil)
	}
}

func HandlePhoto(db *database.DBImpl, fsmBot *fsm.FSM) func(c tele.Context) error {
	return func(c tele.Context) error {
		return handleContent(db, fsmBot, c, c.Message().Caption, c.Message().Photo)
	}
}

func broadcastPost(c tele.Context, userChatIDs []int64, text string, photo *tele.Photo) error {
	if text == "" && photo == nil {
		return fmt.Errorf("no text or photo provided")
	}

	var count uint
	for _, chatID := range userChatIDs {
		if chatID == c.Sender().ID {
			continue
		}
		user := &tele.User{ID: chatID}

		if text != "" {
			if _, err := c.Bot().Send(user, text); err != nil {
				slog.Error("Failed to send message", "userID", chatID, "error", err)
				return c.Send(fmt.Sprintf("Failed to send message: %v", err))
			}
		}
		if photo != nil {
			photo.Caption = text
			if _, err := c.Bot().Send(user, photo); err != nil {
				slog.Error("Failed to send photo", "userID", chatID, "error", err)
				return c.Send(fmt.Sprintf("Failed to send message: %v", err))
			}
		}
		count++
	}
	return c.Send(fmt.Sprintf("Sent to %d users", count))
}
