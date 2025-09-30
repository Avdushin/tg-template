package main

import (
	"tg-template/internal/config"
	"tg-template/internal/database"
	"tg-template/internal/telegram"
)

func main() {
	cfg := config.LoadConfig()

	//? Подключение к БД (можно пропустить, если не нужно)
	db := database.ConnectPostgres(cfg.PostgresDSN)
	defer db.Close()

	//? Создаем и запускаем Telegram бота
	bot := telegram.NewBot(cfg.TelegramToken, cfg.DebugMode)
	bot.Start()
}
