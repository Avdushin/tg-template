package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendSettings(botAPI *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Choose your city:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Moscow", "city:Moscow"),
			tgbotapi.NewInlineKeyboardButtonData("London", "city:London"),
			tgbotapi.NewInlineKeyboardButtonData("Paris", "city:Paris"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("New York", "city:New York"),
			tgbotapi.NewInlineKeyboardButtonData("Chicago", "city:Chicago"),
		),
	)
	botAPI.Send(msg)
}

func init() {
	Register(&Command{
		Name:        "/settings",
		Description: "Set your city for weather",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			SendSettings(botAPI, msg.Chat.ID)
		},
	})
}
