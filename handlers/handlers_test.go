package handlers_test

import (
	"os"
	"testing"

	"github.com/mrpiggy97/rest/server"
)

func TestHandlers(testCase *testing.T) {
	testCase.Run("action=test-home", testHome)
	testCase.Run("action=test-signup-handler", testSignUpHandler)
}

func TestMain(m *testing.M) {
	go server.Runserver(server.GetTestingConfig())
	os.Exit(m.Run())
}
