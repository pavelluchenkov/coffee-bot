package app

import (
	"coffee-bot/internal/db"
	"coffee-bot/internal/handler"
	"coffee-bot/internal/httpserver"
	"coffee-bot/internal/repository"
	"coffee-bot/internal/service"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	server      *httpserver.Server
	menuService      *service.MenuService
	userStateService *service.UserStateService
}

func New(connString string) *App {
	pool := db.New(connString)
	repo := repository.New(pool)
	menuService := service.NewMenuService(repo)
	userStateService := service.NewUserStateService(repo)
	menuHandler := handler.NewMenuHandler(menuService)
	

	router := httpserver.NewRouter(menuHandler)
	server := httpserver.New(router)

	return &App{
		server:      server,
		menuService: menuService,
		userStateService: userStateService,
	}
}

func (a *App) Run() {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		a.server.Run()
	}()

	<-stop

	log.Println("📴 shutdown signal received")

	a.server.Shutdown(context.Background())
}
