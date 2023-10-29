package pipe

import (
	"bytes"
	"encoding/json"
	"http-request-golang/config"
	"io"
	"net/http"
	"os"
)

var loggerError = config.NewErrorLogger()
var requestLogger = config.NewRequestLogger()

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
