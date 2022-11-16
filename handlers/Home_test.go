package handlers_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/mrpiggy97/rest/server"
)

func testHome(testCase *testing.T) {
	var config *server.Config = server.GetTestingConfig()
	go server.Runserver(config)
	time.Sleep(time.Millisecond * 1500)
	req, _ := http.NewRequest("GET", "http://localhost:5000/", nil)
	var client *http.Client = &http.Client{}
	response, responseError := client.Do(req)
	if responseError != nil {
		testCase.Errorf("expected responseError to be nil got %v", responseError.Error())
		os.Exit(2)
	}
	if response.StatusCode != 200 {
		testCase.Errorf("expected status code to be 200 got %d instead", response.StatusCode)
	}
}
