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
	//A partir da demanda em goroutina, carrega em memoria, processa requisição a requisição e envia HTTP
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
			requests := pipe.NewJSON(file.Name())

			for _, request := range requests {
				pipe.NewRequest(request)
			}

		}(file)
	}
	wg.Wait()

}
