package main

import (
	"go-telebot-init/internal/configs"
	"go-telebot-init/pkg/bot"
	"log/slog"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	err = bot.Start()
	if err != nil {
		slog.Error("Failed to start bot", "Fatal", err)
	}
}
