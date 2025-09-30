package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Command struct {
	Name        string
	Description string
	Execute     func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message)
}

var AllCommands []*Command

// @ Функция регистрации команды
// ? Команды регистрируются автоматически
func Register(cmd *Command) {
	AllCommands = append(AllCommands, cmd)
}
