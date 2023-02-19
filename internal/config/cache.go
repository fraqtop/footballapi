package config

type CacheConfig struct {
	host     string
	password string
	port     string
}

func (this CacheConfig) Host() string {
	return this.host
}

func (this CacheConfig) Password() string {
	return this.password
}

func (this CacheConfig) Port() string {
	return this.port
}
