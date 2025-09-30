package commands

import (
	"tg-template/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// * Command: /something
func init() {
	Register(&Command{
		Name:        "/something",
		Description: "Do something",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			services.NewExampleService().DoSomething()
			botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, "Doing something"))
		},
	})
}
