package websockets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients    []*Client
	Register   chan *Client
	Unregister chan *Client
	Number     chan int
	Locker     *sync.Mutex
}

func (hub *Hub) HandleWebSocket(writer http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	var client *Client = NewClient(hub, socket)
	fmt.Println(client.Id, client.Socket.RemoteAddr().String())
	hub.Register <- client
	fmt.Println("it gets here")
	go client.Write()
}

func (hub *Hub) onConnect(client *Client) {
	fmt.Println("client has connected ", client.Socket.RemoteAddr())
	hub.Locker.Lock()
	defer hub.Locker.Unlock()
	client.Id = client.Socket.RemoteAddr().String()
	hub.Clients = append(hub.Clients, client)
}

func (hub *Hub) onDisconnect(client *Client) {
	fmt.Println("client has disconnected ", client.Socket.RemoteAddr().String())
	client.Socket.Close()
	hub.Locker.Lock()
	var newSlice []*Client = make([]*Client, 0)
	defer hub.Locker.Unlock()
	for _, currentClient := range hub.Clients {
		if currentClient.Id != client.Id {
			newSlice = append(newSlice, currentClient)
		}
	}
	hub.Clients = newSlice
}

func (hub *Hub) Run() {
	for {
		fmt.Println("waiting")
		select {
		case client := <-hub.Register:
			fmt.Println("registering client")
			hub.onConnect(client)
		case client := <-hub.Unregister:
			fmt.Println("disconnecting client")
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) BroadCast(message interface{}, ignore *Client) {
	data, _ := json.Marshal(message)
	for _, client := range hub.Clients {
		if client != ignore {
			client.Outbound <- data
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make([]*Client, 0),
		Register:   make(chan *Client, 100),
		Unregister: make(chan *Client, 100),
		Locker:     new(sync.Mutex),
		Number:     make(chan int, 1),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
