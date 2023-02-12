package competition

import (
	"database/sql"
	"github.com/fraqtop/footballapi/internal/connection"
	"github.com/fraqtop/footballcore/competition"
	"github.com/labstack/gommon/log"
)

type readRepository struct {
	connection *sql.DB
}

func (r readRepository) All() []competition.Competition {
	rows, err := r.connection.Query("select id, title from competition order by id")
	var competitions []competition.Competition
	if err != nil {
		return competitions
	}

	var (
		id    int
		title string
	)

	for rows.Next() {
		if err = rows.Scan(&id, &title); err == nil {
			competitions = append(competitions, competition.New(id, title))
		} else {
			log.Warn(err)
		}
	}

	return competitions
}

func NewReadRepository() (competition.ReadRepository, error) {
	connectionInstance, err := connection.GetStorage()
	if err != nil {
		return nil, err
	}
	return readRepository{connection: connectionInstance}, nil
}
