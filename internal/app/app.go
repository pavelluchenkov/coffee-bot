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
	server           *httpserver.Server
	menuService      *service.MenuService
	userStateService *service.UserStateService
	messageService   *service.MessageService
	orderService     *service.OrderService
}

func New(connString string) *App {
	pool := db.New(connString)
	repo := repository.New(pool)
	menuService := service.NewMenuService(repo)
	userStateService := service.NewUserStateService(repo)
	orderService := service.NewOrderService(repo)
	menuHandler := handler.NewMenuHandler(menuService)
	messageService := service.NewMessageService(menuService, userStateService, orderService)

	router := httpserver.NewRouter(menuHandler)
	server := httpserver.New(router)

	return &App{
		server:           server,
		menuService:      menuService,
		userStateService: userStateService,
		messageService:   messageService,
		orderService:     orderService,
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

func (a *App) Test() {
	ctx := context.Background()

	userID := int64(123)

	tests := []string{
		"menu:start",
		"category:classic",
		"drink:Латте",
		"size:M",
	}

	for _, t := range tests {
		resp, err := a.messageService.ProcessMessage(ctx, userID, t)
		if err != nil {
			log.Println("ERR:", err)
			continue
		}

		log.Println("INPUT:", t)
		log.Println("OUTPUT:", resp)
		log.Println("-------------------")
	}
}
