package telegram

import (
	"fmt"
	"strings"
	"tg-template/internal/services/weather"
	"tg-template/internal/telegram/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Обработка обычных команд
func HandleMessage(bot *Bot, message *tgbotapi.Message) {
	for _, cmd := range commands.AllCommands {
		if strings.HasPrefix(message.Text, cmd.Name) {
			cmd.Execute(bot.API, message)
			return
		}
	}
	bot.API.Send(tgbotapi.NewMessage(message.Chat.ID, "I don't understand that command."))
}

// Обработка callback query для выбора города
func HandleCallback(bot *Bot, callback *tgbotapi.CallbackQuery) {
	if callback.Data == "" {
		return
	}

	if strings.HasPrefix(callback.Data, "city:") {
		city := strings.TrimPrefix(callback.Data, "city:")
		userID := callback.From.ID

		// Сохраняем город
		commands.UserCities[userID] = city

		// Мгновенно уведомляем Telegram
		_, err := bot.API.Request(tgbotapi.NewCallback(callback.ID, fmt.Sprintf("City %s saved!", city)))
		if err != nil {
			fmt.Println("Callback request error:", err)
		}

		// Отправляем сообщение с погодой
		go func() {
			result, err := weather.GetWeather(city)
			msgText := fmt.Sprintf("City saved: %s\n%s", city, result)
			if err != nil {
				msgText = fmt.Sprintf("Error: %v", err)
			}
			_, err = bot.API.Send(tgbotapi.NewMessage(callback.Message.Chat.ID, msgText))
			if err != nil {
				fmt.Println("Send message error:", err)
			}
		}()
	}
}
