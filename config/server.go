package config

type ServerConfig struct {
	port string
}

func (sc ServerConfig) Port() string {
	return sc.port
}
