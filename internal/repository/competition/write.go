package competition

import (
	"database/sql"
	"errors"
	"github.com/fraqtop/footballcore/competition"
)

type writeRepository struct {
	connection *sql.DB
}

var (
	_                     competition.WriteRepository = (*writeRepository)(nil)
	ErrInvalidCompetition                             = errors.New("competition is invalid, can't save")
)

func (w writeRepository) Save(competition competition.Competition) error {
	if !competition.IsValid() {
		return ErrInvalidCompetition
	}

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
