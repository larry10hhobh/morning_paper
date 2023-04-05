package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
	"strings"
	"time"
)

func getBotWithSocks5Proxy() (bot *tgbotapi.BotAPI, err error) {

	if getConfig().UseProxy {
		dialer, err := proxy.SOCKS5("tcp", getConfig().ProxyAddress, nil, proxy.Direct)
		if err != nil {
			log.Fatalf("Error creating SOCKS5 proxy dialer: %v", err)
		}

		httpTransport := &http.Transport{
			Dial: dialer.Dial,
		}
		httpClient := &http.Client{
			Timeout:   time.Second * 30,
			Transport: httpTransport,
		}

		bot, err = tgbotapi.NewBotAPIWithClient(getConfig().TelegramBotToken, httpClient)
		if err != nil {
			log.Fatalf("Error creating Telegram bot with proxy: %v", err)
		}
	} else {
		bot, err = tgbotapi.NewBotAPI(getConfig().TelegramBotToken)
		if err != nil {
			log.Fatalf("Error creating Telegram bot: %v", err)
		}
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	log.Printf("Authorized on account %s", bot.Self.ID)

	return bot, nil
}

func sendToTelegram(message string, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(config.TelegramUserID, message)

	_, err := bot.Send(msg)
	if err != nil {
		if strings.Contains(err.Error(), "chat not found") {
			log.Println("Error: chat not found. Please start a conversation with the bot first.")
		} else {
			log.Printf("Error sending message: %v", err)
		}
	}
}
