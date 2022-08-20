package config

type StorageConfig struct {
	host     string
	port     string
	user     string
	password string
	name     string
}

func (sc StorageConfig) Host() string {
	return sc.host
}

func (sc StorageConfig) Port() string {
	return sc.port
}

func (sc StorageConfig) User() string {
	return sc.user
}

func (sc StorageConfig) Password() string {
	return sc.password
}

func (sc StorageConfig) Name() string {
	return sc.name
}
