package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/mrpiggy97/rest/handlers"
	"github.com/mrpiggy97/rest/server"
)

func testSignUpHandler(testCase *testing.T) {
	// initialize server
	go server.Runserver(server.GetTestingConfig())
	// give time for server to come up
	time.Sleep(time.Millisecond * 1500)
	// set data to send and encode it to json
	var data map[string]string = make(map[string]string)
	data["Email"] = "testin20@email.com"
	data["Password"] = "tes2tingpassword10"
	jsonData, _ := json.Marshal(data)
	var bufferer *bytes.Buffer = bytes.NewBuffer(jsonData)
	// make http request
	req, _ := http.NewRequest("POST", "http://localhost:5000/signup", bufferer)
	var client *http.Client = &http.Client{}
	response, responseError := client.Do(req)
	// make tests
	if responseError != nil {
		testCase.Errorf("expected response error to be nil got %v instead", responseError.Error())
		os.Exit(2)
	}
	if response.StatusCode != 201 {
		testCase.Errorf("expected status code to be 201 got %d instead", response.StatusCode)
		os.Exit(2)
	}

	var decodedResponse *handlers.SignUpResponse = new(handlers.SignUpResponse)
	json.NewDecoder(response.Body).Decode(decodedResponse)
	var log string = fmt.Sprintf("token %v expiration date %v", decodedResponse.Token, decodedResponse.ExpirationDate)
	fmt.Println(log)
}
