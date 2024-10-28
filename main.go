package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, _ := tgbotapi.NewBotAPI(jsonTokenDecoder())

	b := NewBot(bot)
	b.Start()
}
