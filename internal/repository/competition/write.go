package competition

import (
	"database/sql"
	"github.com/fraqtop/footballcore/competition"
)

type writeRepository struct {
	connection *sql.DB
}

var _ competition.WriteRepository = (*writeRepository)(nil)

func (w writeRepository) Save(competition competition.Competition) error {
	var err error
	_, err = w.connection.Exec("insert into competition (id, title) "+
		"values ($1, $2) "+
		"on conflict(id) do update "+
		"set title = excluded.title", competition.Id(), competition.Title())
	if err != nil {
		return err
	}

	return err
}

func NewWriteRepository(connection *sql.DB) competition.WriteRepository {
	return &writeRepository{connection: connection}
}
