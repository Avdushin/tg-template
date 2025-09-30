package commands

import (
	"tg-template/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SomethingCommand struct{}

func (c *SomethingCommand) Name() string {
	return "/something"
}

func (c *SomethingCommand) Execute(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	services.NewExampleService().DoSomething()
	botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, "Doing something"))
}
