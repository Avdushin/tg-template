package utils

import (
	"html"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendMarkdown отправляет сообщение с форматированием MarkdownV2
func SendMarkdown(botAPI *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := botAPI.Send(msg); err != nil {
		log.Printf("Failed to send Markdown message to chat %d: %v", chatID, err)
	}
}

// SendHTML отправляет сообщение с форматированием HTML
func SendHTML(botAPI *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML

	if _, err := botAPI.Send(msg); err != nil {
		log.Printf("Failed to send HTML message to chat %d: %v", chatID, err)
	}
}

// EscapeHTML экранирует текст для безопасного использования в HTML-сообщениях
func EscapeHTML(text string) string {
	return html.EscapeString(text)
}
