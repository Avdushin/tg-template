package commands

import (
	"tg-template/internal/services/example"
	"tg-template/internal/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// * Command: /something
func init() {
	Register(&Command{
		Name:        "/something",
		Description: "Do something",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			//? Пример использования стороннего сервиса
			example.NewExampleService().DoSomething()
			//@ Сообщение пользователю
			// botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, "Doing something"))
			utils.SendHTML(botAPI, msg.Chat.ID, `<b>Doing something</b> ✅`)
		},
	})
}
