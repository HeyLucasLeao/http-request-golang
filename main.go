package main

import (
	"http-request-golang/config"
	pipe "http-request-golang/pipeline"
	"io/fs"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var loggerError = config.NewErrorLogger()
var loggerInfo = config.NewInfoLogger()

//var loggerInfo = config.NewInfoLogger()

func main() {
	var wg sync.WaitGroup

	err := godotenv.Load()

	if err != nil {
		loggerError.Fatal(err)
	}

	start := time.Now()

	files := config.NewFile()

	wg.Add(len(files))
	for _, file := range files {

		go func(file fs.DirEntry) {
			defer wg.Done()
			requests := config.NewJSON(file.Name())
			nestedWg := sync.WaitGroup{}
			nestedMu := sync.Mutex{}

			nestedWg.Add(len(requests))
			for _, request := range requests {
				go func(request any) {
					nestedMu.Lock()
					defer nestedWg.Done()
					pipe.NewRequest(request)
					nestedMu.Unlock()
				}(request)
			}

			nestedWg.Wait()
		}(file)
	}
	wg.Wait()
	elapsed := time.Since(start)

	loggerInfo.Printf("Function took %s", elapsed)
}
