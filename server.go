package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	port   int
	logger *RequestLogger
	server *http.Server
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
	}
}

func (s *Server) Start() {
	go func() {
		fmt.Printf("Server is starting on port: %d\n", s.port)
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("Server is shutting down.")
	var shutdownTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Printf("Error stopping server: %s\n", err)
		os.Exit(1)
	}
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
