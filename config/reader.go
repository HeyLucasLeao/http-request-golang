package config

import (
	"fmt"
	"io/fs"
	"os"
)

var loggerError = NewErrorLogger()

func NewFile() []fs.DirEntry {
	root := "data"
	pattern := os.Getenv("JSON_FOLDER")
	path := fmt.Sprintf("%s/%s", root, pattern)
	file, err := os.ReadDir(path)

	if err != nil {
		loggerError.Fatal(err)
	}

	return file
}
