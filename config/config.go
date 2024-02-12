package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	PublicUrl   string
	TelegramUrl string
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	PublicUrl = os.Getenv("PUBLIC_URL")
	TelegramUrl = os.Getenv("TELEGRAM_URL")
}
