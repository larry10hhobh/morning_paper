package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	// Telegram
	TelegramUserID   int64  `toml:"telegram_user_id"`
	TelegramBotToken string `toml:"telegram_bot_token"`

	// Proxy
	UseProxy     bool   `toml:"use_proxy"`
	ProxyAddress string `toml:"proxy_address"`

	// Seniverse
	SeniverseAPIKey   string `toml:"seniverse_api_key"`
	SeniverseLocation string `toml:"seniverse_location"`
}

var config Config

func loadConfig() error {
	_, err := toml.DecodeFile("config.toml", &config)
	return err
}

func getConfig() Config {
	return config
}
