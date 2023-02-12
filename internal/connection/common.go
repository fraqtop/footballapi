package connection

import (
	"database/sql"
	"fmt"
	"github.com/fraqtop/footballapi/internal/config"
)

func Destroy() error {
	if instance != nil {
		return instance.Close()
	}

	return nil
}

func Init() error {
	databaseConfig := config.GetStorageConfig()
	newInstance, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
			databaseConfig.User(),
			databaseConfig.Password(),
			databaseConfig.Host(),
			databaseConfig.Port(),
			databaseConfig.Name(),
			"disable",
		),
	)
	instance = newInstance

	return err
}
