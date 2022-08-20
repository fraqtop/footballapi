package connection

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var instance *sql.DB

func GetStorage() (*sql.DB, error) {
	if instance == nil {
		err := Init()
		if err != nil {
			return nil, err
		}
	}

	return instance, nil
}
