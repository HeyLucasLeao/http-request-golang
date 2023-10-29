package config

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var loggerError = NewErrorLogger()
var requestLogger = NewRequestLogger()

func NewErrorLogger() *log.Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	logger := log.New(os.Stdout, "ERROR🚨: ", flags)
	return logger
}

func NewRequestLogger() *log.Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	path := filepath.Join("log", "response.log")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		panic(err)
	}
	logger := log.New(f, "", flags)

	return logger
}

func NewInfoLogger() *log.Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	logger := log.New(os.Stdout, "INFO✅: ", flags)
	return logger
}

func LoggingResponse(resp *http.Response) {
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		loggerError.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		requestLogger.SetPrefix("INFO✅: ")
		requestLogger.Printf("Code: %d - Body: %s", resp.StatusCode, body)
		return
	}

	requestLogger.SetPrefix("ERROR🚨: ")
	requestLogger.Printf("Code: %d - Body: %s", resp.StatusCode, body)
}
