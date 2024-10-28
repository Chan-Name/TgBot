package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Start() {

	updates, _ := b.initStart()

	for update := range updates {
		if update.Message != nil { // ignore any non-Message Updates
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, valuterStart())
			b.bot.Send(msg)
		}
	}
}

func (b *Bot) initStart() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}
