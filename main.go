package main

import (
	"TG/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, _ := tgbotapi.NewBotAPI(jsonTokenDecoder())

	b := telegram.NewBot(bot)
	b.Start()
}
