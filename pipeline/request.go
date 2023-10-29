package pipe

import (
	"bytes"
	"encoding/json"
	"http-request-golang/config"
	"net/http"
	"os"
)

var loggerError = config.NewErrorLogger()

func NewRequest(request interface{}) *http.Response {
	client := &http.Client{}
	endPoint := os.Getenv("HTTP_ENDPOINT")
	bearerToken := "Bearer " + os.Getenv("BEARER_TOKEN")
	body, err := json.Marshal(request)

	if err != nil {
		loggerError.Fatal(err)
	}

	r, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(body))

	if err != nil {
		loggerError.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", bearerToken)

	resp, err := client.Do(r)

	if err != nil {
		loggerError.Fatal(err)
	}

	return resp

}
