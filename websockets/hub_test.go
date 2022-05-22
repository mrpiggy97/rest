package websockets_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mrpiggy97/rest/server"
)

func testHandleWebSocket(testCase *testing.T) {
	go server.Runserver()
	time.Sleep(time.Second * 2)
	var url string = "ws://localhost:5050/ws"
	_, _, socketErr := websocket.DefaultDialer.Dial(url, nil)
	//defer socketConn.Close()
	if socketErr != nil {
		fmt.Println("socket connection error")
		testCase.Error(socketErr)
	}
}

func TestHandleWebsocket(testCase *testing.T) {
	testCase.Run("action=test-handle-websocket", testHandleWebSocket)
}
