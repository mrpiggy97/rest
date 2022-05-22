package server

import "github.com/mrpiggy97/rest/websockets"

func GetHub() *websockets.Hub {
	return websockets.NewHub()
}
