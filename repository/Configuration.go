package repository

type IConfig interface {
	GetJWTSecret() string
	GetDatabaseUrl() string
	GetPort() string
}

var ConfigImplementation IConfig

func SetConfig(newConfig IConfig) {
	ConfigImplementation = newConfig
}

func GetJWTSecret() string {
	return ConfigImplementation.GetJWTSecret()
}

func GetDatabaseUrl() string {
	return ConfigImplementation.GetDatabaseUrl()
}

func GetPort() string {
	return ConfigImplementation.GetPort()
}
