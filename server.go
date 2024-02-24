package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	logger    *RequestLogger
	server    *http.Server
	isRunning bool
}

func NewServer(logFilePath string) *Server {
	port, err := getFreePort()
	if err != nil {
		panic(err)
	}

	logger, err := NewRequestLogger(logFilePath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Server Running on: %d\n", port)

	return &Server{
		logger: logger,
		server: &http.Server{
			Addr: strconv.Itoa(port),
		},
		isRunning: true,
	}
}

func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

func (s *Server) Start() {
	if s.isRunning {
		fmt.Println("Server is already running.")
		return
	}

	go func() {
		fmt.Printf("\nServer is starting on port: %s\n", s.server.Addr)
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
