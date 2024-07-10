package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Token           string
	TelegramBotHost string
	StoragePath     string
	BatchSize       int
}

func MustLoad() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	token := os.Getenv("BOT_TOKEN")
	host := os.Getenv("TELEGRAM_BOT_HOST")
	storagePath := os.Getenv("STORAGE_PATH")

	if token == "" {
		log.Fatal("BOT_TOKEN is not set")
	}
	if host == "" {
		log.Fatal("TELEGRAM_BOT_HOST is not set")
	}
	if storagePath == "" {
		log.Fatal("STORAGE_PATH is not set")
	}
	if storagePath == "" {
		log.Fatal("BATCH_SIZE is not set")
	}

	return Config{
		Token:           token,
		TelegramBotHost: host,
		StoragePath:     storagePath,
	}

}
