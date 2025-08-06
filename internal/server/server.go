package server

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler, port string) *Server {
	if port == "" {
		port = "8080"
	}

	httpServer := &http.Server{
		Addr:              net.JoinHostPort("0.0.0.0", port),
		Handler:           handler,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	log.Printf("Server starting on port %s", s.httpServer.Addr)
	log.Printf("API available at: http://localhost%s/api/v1/", s.httpServer.Addr)
	log.Printf("Swagger UI available at: http://localhost%s/swagger", s.httpServer.Addr)

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Ошибка запуска сервера: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Завершение работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("Сервер остановлен")
	return err
}
