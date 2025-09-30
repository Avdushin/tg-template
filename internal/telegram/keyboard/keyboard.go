package keyboard

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

// ReplyCommandKeyboard возвращает обычную клавиатуру с командами
func ReplyCommandKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
			tgbotapi.NewKeyboardButton("/help"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/something"),
			tgbotapi.NewKeyboardButton("/ping"),
		),
	)
}

// RemoveKeyboard возвращает убирающую клавиатуру
func RemoveKeyboard() tgbotapi.ReplyKeyboardRemove {
	return tgbotapi.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      false,
	}
}
