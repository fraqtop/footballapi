package config

type ServerConfig struct {
	port string
	debug bool
}

func (this ServerConfig) Port() string {
	return this.port
}

func (this ServerConfig) Debug() bool {
	return this.debug
}
