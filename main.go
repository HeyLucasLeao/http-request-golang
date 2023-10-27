package main

import (
	"http-request-golang/config"
	pipe "http-request-golang/pipeline"
	"io/fs"
	"sync"

	"github.com/joho/godotenv"
)

var loggerError = config.NewErrorLogger()

//var loggerInfo = config.NewInfoLogger()

func main() {

	var wg sync.WaitGroup
	err := godotenv.Load()

	if err != nil {
		loggerError.Fatal(err)
	}

	files := config.NewFile()

	wg.Add(len(files))
	for _, file := range files {

		go func(file fs.DirEntry) {
			defer wg.Done()
			requests := config.NewJSON(file.Name())
			nestedWg := sync.WaitGroup{}

			nestedWg.Add(len(requests))
			for _, request := range requests {
				go func(request any) {
					defer nestedWg.Done()
					pipe.NewRequest(request)
				}(request)
			}

			nestedWg.Wait()
		}(file)
	}
	wg.Wait()
}
