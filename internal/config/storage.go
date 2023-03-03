package config

type StorageConfig struct {
	host     string
	port     string
	user     string
	password string
	name     string
}

func (this StorageConfig) Host() string {
	return this.host
}

func (this StorageConfig) Port() string {
	return this.port
}

func (this StorageConfig) User() string {
	return this.user
}

func (this StorageConfig) Password() string {
	return this.password
}

func (this StorageConfig) Name() string {
	return this.name
}
