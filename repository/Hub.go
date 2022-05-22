package repository

import "github.com/mrpiggy97/rest/websockets"

var AppHub *websockets.Hub = new(websockets.Hub)

func SetHub(hub *websockets.Hub) {
	AppHub = hub
}
