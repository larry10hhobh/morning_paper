package main

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config.yaml file: %v", err)
	}
}

func main() {
	_, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch weather and news
	// weather := getWeather()
	// news := getNews()

	// Send weather and news to Telegram
	// sendToTelegram(weather, news, bot)
}


func sendToTelegram(weather, news string, bot *tgbotapi.BotAPI) {
	message := fmt.Sprintf("Weather: %s\nNews: %s", weather, news)
	msg := tgbotapi.NewMessage(config.TelegramUserID, message)

	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Message sent: %s", message)
	}
}
