package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	port      int
	logger    *RequestLogger
	server    *http.Server
	isRunning bool
}

func NewServer(port int, logFilePath string) *Server {
	addr := fmt.Sprintf(":%d", port)
	logger, err := NewRequestLogger(logFilePath)

	if err != nil {
		panic(err)
	}

	return &Server{
		port:   port,
		logger: logger,
		server: &http.Server{
			Addr: addr,
		},
		isRunning: true,
	}
}

func (s *Server) Start() {
	if s.isRunning {
		fmt.Println("Server is already running.")
		return
	}

	go func() {
		fmt.Printf("\nServer is starting on port: %d\n", s.port)
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
}

func (s *Server) Stop() {
	if !s.isRunning {
		fmt.Println("Server is not running.")
		return
	}

	fmt.Println("Server is shutting down.")
	var shutdownTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Printf("Error stopping server: %s\n", err)
		os.Exit(1)
	}
	s.isRunning = false
}

func (s *Server) sendRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	s.logger.LogRequest(request)
	response, err := client.Do(request)
	s.logger.LogResponse(response)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response, nil
}
