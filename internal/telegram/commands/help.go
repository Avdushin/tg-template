package commands

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// * Command: /something
func init() {
	Register(&Command{
		Name:        "/help",
		Description: "Show help message",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			var b strings.Builder
			b.WriteString("Available commands:\n")
			for _, cmd := range AllCommands {
				b.WriteString(cmd.Name + " — " + cmd.Description + "\n")
			}
			botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, b.String()))
		},
	})
}
