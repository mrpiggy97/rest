package websockets

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub      *Hub
	Id       string
	Socket   *websocket.Conn
	Outbound chan []byte
}

func (client *Client) Write() {
	for {
		select {
		case message, ok := <-client.Outbound:
			if !ok {
				client.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			fmt.Println("sending message")
			var sendingErr error = client.Socket.WriteMessage(websocket.TextMessage, message)
			if sendingErr != nil {
				fmt.Println(sendingErr.Error())
			}
		}
	}
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		Hub:      hub,
		Socket:   socket,
		Outbound: make(chan []byte, 1),
	}
}
