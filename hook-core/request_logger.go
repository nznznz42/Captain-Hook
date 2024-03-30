package hookcore

import (
	"log"
	"net/http"
	"os"
	"sync"
)

type RequestLogger struct {
	mu      sync.Mutex
	logFile *os.File
	logger  *log.Logger
}

func NewRequestLogger(logFileName string) (*RequestLogger, error) {
	LogFilePath := "Logs/" + logFileName
	logFile, err := os.OpenFile(LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &RequestLogger{
		logFile: logFile,
		logger:  log.New(logFile, "[Request Log] ", log.LstdFlags),
	}, nil
}

func (rl *RequestLogger) LogRequest(r *http.Request) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.logger.Printf("Request: %s %s", r.Method, r.Body)
}

func (rl *RequestLogger) LogResponse(response *http.Response) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if response != nil {
		rl.logger.Printf("Response Status: %d", response.StatusCode)
	} else {
		rl.logger.Println("Response is nil")
	}
}

func (rl *RequestLogger) Close() error {
	return rl.logFile.Close()
}
