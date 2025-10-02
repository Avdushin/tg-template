package commands

import (
	"fmt"
	"html"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	Register(&Command{
		Name:        "/help",
		Description: "Show help message",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			var b strings.Builder

			b.WriteString("*Available commands:*\n")

			for _, cmd := range AllCommands {
				// Экранируем только основное содержимое
				name := html.EscapeString(cmd.Name)
				desc := cmd.Description

				// Проверяем, есть ли уже тег <code> в Description
				// (например "Example: <code>/ping github.com</code>")
				// Если есть, оставляем как есть, иначе экранируем
				line := fmt.Sprintf("%s — %s", name, desc)

				b.WriteString(line + "\n")
			}

			msgConfig := tgbotapi.NewMessage(msg.Chat.ID, b.String())
			msgConfig.ParseMode = tgbotapi.ModeMarkdown

			if _, err := botAPI.Send(msgConfig); err != nil {
				log.Printf("Failed to send help message: %v", err)
			}
		},
	})
}
