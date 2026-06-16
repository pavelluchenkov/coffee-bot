package main

import (
	"coffee-bot/internal/app"
	"coffee-bot/internal/config"
	"coffee-bot/internal/httpserver"
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

	router := httpserver.NewRouter()
	server := httpserver.New(router)
	application := app.New(cfg.DB, server)
	application.Run()

}
