package commands

import (
	"strings"
	"tg-template/internal/services/ping"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	Register(&Command{
		Name:        "/ping",
		Description: "Ping a host and show response time. Example: `/ping github.com`",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			args := strings.Fields(msg.Text)
			host := "google.com"
			if len(args) > 1 {
				host = args[1]
			}

			result := ping.PingHost(host)

			// Отправляем результат с Markdown
			msgConfig := tgbotapi.NewMessage(msg.Chat.ID, result)
			botAPI.Send(msgConfig)
		},
	})
}
