package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
	}
}

func (s *Server) Run() {
	log.Println("🚀 HTTP server started on :8080")

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		log.Fatal(err)
	}
}

func (s *Server) Shutdown(ctx context.Context){
	log.Println("🛑 shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil{
		log.Fatal("shutdown error:", err)
	}
}