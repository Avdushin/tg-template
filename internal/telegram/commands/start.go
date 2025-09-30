package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type StartCommand struct{}

func (c *StartCommand) Name() string {
	return "/start"
}

func (c *StartCommand) Execute(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, "Welcome to the bot!"))
}
