package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *Bot, message *tgbotapi.Message) {
	switch message.Text {
	case "/start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to the bot!")
		bot.API.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "I don't understand that command.")
		bot.API.Send(msg)
	}
}
