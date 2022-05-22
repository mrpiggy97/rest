package websockets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	locker     *sync.Mutex
}

func (hub *Hub) HandleWebSocket(writer http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	var client *Client = NewClient(hub, socket)
	hub.register <- client
	go client.Write()
}

func (hub *Hub) onConnect(client *Client) {
	fmt.Println("client has connected ", client.socket.RemoteAddr())
	hub.locker.Lock()
	defer hub.locker.Unlock()
	client.id = client.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, client)
}

func (hub *Hub) onDisconnect(client *Client) {
	fmt.Println("client has disconnected ", client.socket.RemoteAddr().String())
	client.socket.Close()
	hub.locker.Lock()
	var newSlice []*Client = make([]*Client, 0)
	defer hub.locker.Unlock()
	for _, currentClient := range hub.clients {
		if currentClient.id != client.id {
			newSlice = append(newSlice, currentClient)
		}
	}
	hub.clients = newSlice
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) BroadCast(message interface{}, ignore *Client) {
	data, _ := json.Marshal(message)
	for _, client := range hub.clients {
		if client != ignore {
			client.outbound <- data
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client, 1),
		unregister: make(chan *Client, 1),
		locker:     new(sync.Mutex),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
