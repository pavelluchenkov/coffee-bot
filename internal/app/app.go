package app

import (
	"coffee-bot/internal/db"
	"coffee-bot/internal/httpserver"
	"coffee-bot/internal/repository"
	"coffee-bot/internal/services"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	server      *httpserver.Server
	menuService *services.MenuService
}

func New(connString string, server *httpserver.Server) *App {
	pool := db.New(connString)
	repo := repository.New(pool)
	menuService := services.NewMenuService(repo)
	return &App{
		server:      server,
		menuService: menuService,
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
