package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

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

func NewJSON(file string) []any {
	root := "data"
	pattern := os.Getenv("JSON_FOLDER")
	path := filepath.Join(root, pattern, file)
	f, err := os.ReadFile(path)

	if err != nil {
		loggerError.Fatal(err)
	}

	requests := []any{}

	err = json.Unmarshal(f, &requests)

	if err != nil {
		loggerError.Fatal(err)
	}

	return requests
}
