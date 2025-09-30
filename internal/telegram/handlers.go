package telegram

import (
	"strings"
	"tg-template/internal/telegram/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// @ commands handler
func HandleMessage(bot *Bot, message *tgbotapi.Message) {
	for _, cmd := range commands.AllCommands {
		if strings.HasPrefix(message.Text, cmd.Name) {
			cmd.Execute(bot.API, message)
			return
		}
	}
	bot.API.Send(tgbotapi.NewMessage(message.Chat.ID, "I don't understand that command."))
}
