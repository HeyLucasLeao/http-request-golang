package main

import (
	"http-request-golang/config"
	pipe "http-request-golang/pipeline"
	"io/fs"
	"net/http"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var loggerError = config.NewErrorLogger()
var loggerInfo = config.NewInfoLogger()

func main() {
	var wg sync.WaitGroup
	chanResp := make(chan *http.Response)
	err := godotenv.Load()

	if err != nil {
		loggerError.Fatal(err)
	}

	start := time.Now()

	files := config.NewFile()

	wg.Add(len(files))
	for _, file := range files {

		//Read and Unmarshall the files concurrently
		go func(file fs.DirEntry) {
			defer wg.Done()
			requests := config.NewJSON(file.Name())
			nestedWg := sync.WaitGroup{}

			//For every value, send concurrently HTTP Request
			nestedWg.Add(len(requests))
			for _, request := range requests {

				go func(request any) {
					defer nestedWg.Done()
					resp := pipe.NewRequest(request)
					chanResp <- resp
				}(request)

				//Save response in a log file
				go config.LoggingResponse(<-chanResp)
			}

			nestedWg.Wait()
		}(file)
	}

	wg.Wait()
	elapsed := time.Since(start)

	loggerInfo.Printf("Function took %s", elapsed)
}
