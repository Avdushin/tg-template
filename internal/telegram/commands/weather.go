package commands

import (
	"fmt"
	"tg-template/internal/services/weather"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	Register(&Command{
		Name:        "/weather",
		Description: "Show weather for your city",
		Execute: func(botAPI *tgbotapi.BotAPI, msg *tgbotapi.Message) {
			userID := msg.From.ID
			city, ok := UserCities[userID]
			if !ok {
				botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, "You haven't set a city yet. Use /settings"))
				return
			}
			result, err := weather.GetWeather(city)
			if err != nil {
				botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Error: %v", err)))
				return
			}
			botAPI.Send(tgbotapi.NewMessage(msg.Chat.ID, result))
		},
	})
}
