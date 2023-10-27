package config

import (
	"log"
	"os"
	"path/filepath"
)

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
