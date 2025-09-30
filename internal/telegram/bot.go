package telegram

import (
	"log"
	"tg-template/internal/telegram/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

func NewBot(token string, debug bool) *Bot {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	//@ Debug режим
	botAPI.Debug = debug

	//@ Инициализация всех команд
	// commands.Init()

	//@ Установка команд в меню Telegram
	setCommandMenu(botAPI)

	return &Bot{
		API: botAPI,
	}
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.API.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go HandleMessage(b, update.Message)
		}
	}
}

// setCommandMenu устанавливает список команд для автоподсказки
func setCommandMenu(botAPI *tgbotapi.BotAPI) {
	var botCommands []tgbotapi.BotCommand
	for _, cmd := range commands.AllCommands {
		botCommands = append(botCommands, tgbotapi.BotCommand{
			Command:     cmd.Name[1:],
			Description: cmd.Description,
		})
	}

	if _, err := botAPI.Request(tgbotapi.NewSetMyCommands(botCommands...)); err != nil {
		log.Printf("Failed to set bot commands: %v", err)
	}
}
