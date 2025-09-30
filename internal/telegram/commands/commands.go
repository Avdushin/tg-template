package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Command interface {
	Name() string
	Execute(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message)
}
