package config

type BrokerConfig struct {
	host string
}

func (this BrokerConfig) Host() string {
	return this.host
}
