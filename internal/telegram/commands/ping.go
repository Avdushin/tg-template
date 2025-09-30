package commands

import (
	"strings"
	"tg-template/internal/services/ping"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	Register(&Command{
		Name:        "/ping",
		Description: "Ping a host and show response time",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			args := strings.Fields(msg.Text) // разбиваем сообщение на слова
			host := "google.com"
			if len(args) > 1 {
				host = args[1] // второй аргумент — это хост
			}
			result := ping.PingHost(host)
			botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, result))
		},
	})
}
