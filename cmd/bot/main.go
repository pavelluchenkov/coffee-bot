package main

import (
	"coffee-bot/internal/app"
	"coffee-bot/internal/config"
	"fmt"
)

func main() {
	cfg := config.Load()

	fmt.Println("Bot starting..")

	// database := db.New(cfg.DB)
	// defer database.Close()

	// vkClient := vk.New(cfg.VKToken, cfg.GroupID)
	// _ = vkClient.SendMessage(123456789, "Бот запущен")

	// select {}
	application := app.New(cfg.DB)
	application.Test()
	application.Run()

}
