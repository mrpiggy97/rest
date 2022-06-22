package repository

type IHub interface {
	Run()
	GetNumberOfActiveClients() int
	OnConnect(client IClient)
	OnDisconnect(client IClient)
	BroadCast(message interface{}, ignoreClientId string)
	RegisterClient(client IClient)
	DeregisterClient(client IClient)
}

var AppHub IHub

func SetHub(hub IHub) {
	AppHub = hub
}

func GetHub() IHub {
	return AppHub
}

func GetNumberOfActiveClients() int {
	return AppHub.GetNumberOfActiveClients()
}

func BroadCast(message interface{}, ignoreClientId string) {
	AppHub.BroadCast(message, ignoreClientId)
}

func RegisterClient(client IClient) {
	AppHub.RegisterClient(client)
}

func DeregisterClient(client IClient) {
	AppHub.DeregisterClient(client)
}

func RunHub() {
	go AppHub.Run()
}
