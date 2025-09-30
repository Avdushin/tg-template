package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// CommandKeyboard возвращает клавиатуру с кнопками команд
func CommandKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("/start", "/start"),
			tgbotapi.NewInlineKeyboardButtonData("/help", "/help"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("/something", "/something"),
			tgbotapi.NewInlineKeyboardButtonData("/ping", "/ping"),
		),
	)
}
