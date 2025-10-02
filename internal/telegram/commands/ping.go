package commands

import (
	"fmt"
	"strings"
	"tg-template/internal/services/ping"
	"tg-template/internal/telegram/utils"

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

			text := fmt.Sprintf(
				"<b>Ping result</b> for <a href=\"https://%s\">%s</a>:\n%s",
				utils.EscapeHTML(host),
				utils.EscapeHTML(host),
				utils.EscapeHTML(result),
			)

			utils.SendHTML(botAPI, msg.Chat.ID, text)

		},
	})
}
