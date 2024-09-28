package bot

import (
	"fmt"
	"github.com/spf13/viper"
	"go-telebot-init/pkg/bot/fsm"
	ah "go-telebot-init/pkg/bot/handlers/admin"
	ph "go-telebot-init/pkg/bot/handlers/public"
	"go-telebot-init/pkg/bot/middlewares"
	"go-telebot-init/pkg/database"
	tele "gopkg.in/telebot.v3"
	"log/slog"
	"time"
)

type TelegramBot struct {
	Bot *tele.Bot
	DB  *database.DBImpl
	FSM *fsm.FSM
}

func NewTelegramBot() (*TelegramBot, error) {
	bot, err := createBot()
	if err != nil {
		return nil, err
	}
	db, err := database.Init()
	if err != nil {
		return nil, err
	}

	finiteStateMachine := fsm.NewFSM()

	return &TelegramBot{
		Bot: bot,
		DB:  db,
		FSM: finiteStateMachine,
	}, nil
}

func createBot() (*tele.Bot, error) {
	retries := viper.GetInt("config.retries")
	for i := 0; i < retries; i++ {
		pref := tele.Settings{
			Token:  viper.GetString("TELEGRAM_TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}
		bot, err := tele.NewBot(pref)
		if err == nil {
			return bot, nil
		}

		slog.Error("Failed to create bot, retrying", "error", err)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("failed to create bot after %d attempts", retries)
}

func (tb *TelegramBot) registerHandlers() {
	admin := tb.Bot.Group()
	admin.Use(middlewares.IsAdmin(tb.DB))
	admin.Handle("/post", ah.HandlePost(tb.FSM))
	admin.Handle(tele.OnText, ah.HandleText(tb.DB, tb.FSM))
	admin.Handle(tele.OnPhoto, ah.HandlePhoto(tb.DB, tb.FSM))

	public := tb.Bot.Group()
	public.Handle("/start", ph.HandleStart(tb.DB))
}

func Start() error {
	tb, err := NewTelegramBot()
	if err != nil {
		return err
	}

	tb.registerHandlers()

	slog.Info("Starting bot", "Bot", tb.Bot.Me.Username)
	tb.Bot.Start()
	return nil
}
