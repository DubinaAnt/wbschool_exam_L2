package internal

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Run запускает сервер
func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown остановдивает сервер
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
