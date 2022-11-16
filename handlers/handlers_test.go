package handlers_test

import "testing"

func TestHandlers(testCase *testing.T) {
	testCase.Run("action=test-home", testHome)
	testCase.Run("action=test-signup-handler", testSignUpHandler)
}
