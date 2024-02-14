package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	Port        string
	Host        string
	PublicUrl   string
	TelegramUrl string
)

func init() {
	var err error

	if os.Getenv("GO_ENV") != "production" {
		err = godotenv.Load()
	}
	if err != nil {
		panic(err)
	}

	Port = os.Getenv("PORT")
	Host = os.Getenv("HOST")
	PublicUrl = os.Getenv("PUBLIC_URL")
	TelegramUrl = os.Getenv("TELEGRAM_URL")

	if Port == "" || Host == "" || PublicUrl == "" || TelegramUrl == "" {
		panic("Invalid env variables")
	}
}
