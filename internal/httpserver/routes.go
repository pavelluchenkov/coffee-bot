package httpserver

import (
	"coffee-bot/internal/handler"
	"net/http"
)

type Router struct {
	menuHandler *handler.MenuHandler
}

func NewRouter(menuHandler *handler.MenuHandler) http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/menu/categories", menuHandler.GetCategories)
	mux.HandleFunc("/menu/", menuHandler.GetByCategory)

	return mux
}
