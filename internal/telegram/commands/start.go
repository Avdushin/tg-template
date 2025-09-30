package commands

import (
	"tg-template/internal/telegram/keyboard"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// * Command: /start
func init() {
	Register(&Command{
		Name:        "/start",
		Description: "Start the bot",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			message := tgbotapi.NewMessage(msg.Chat.ID, "Welcome to the bot!")
			message.ReplyMarkup = keyboard.ReplyCommandKeyboard()
			// Удалить клавиатуру
			// message.ReplyMarkup = keyboard.RemoveKeyboard()
			botAPI.Send(message)
		},
	})
}

// package commands

// import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// // * Command: /start
// func init() {
// 	Register(&Command{
// 		Name:        "/start",
// 		Description: "Start the bot",
// 		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
// 			botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, "Welcome to the bot!"))
// 		},
// 	})
// }
