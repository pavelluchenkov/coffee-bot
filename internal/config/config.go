package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	VKToken string
	GroupID string
	DB   string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system env")
	}
	return Config{
		VKToken: os.Getenv("VK_TOKEN"),
		GroupID: os.Getenv("GROUP_ID"),
		DB:     os.Getenv("DB_DSN"),
	}
}
