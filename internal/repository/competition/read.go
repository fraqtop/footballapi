package competition

import (
	"database/sql"
	"github.com/fraqtop/footballcore/competition"
	"github.com/labstack/gommon/log"
)

type readRepository struct {
	connection *sql.DB
}

var _ competition.ReadRepository = (*readRepository)(nil)

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

func NewReadRepository(connectionInstance *sql.DB) competition.ReadRepository {
	return readRepository{connection: connectionInstance}
}
