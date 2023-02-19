package team

import (
	"database/sql"
	"fmt"
	"github.com/fraqtop/footballcore/team"
	"strings"
)

var repository *writeRepository

type writeRepository struct {
	connection *sql.DB
}

var _ team.WriteRepository = (*writeRepository)(nil)

func (w writeRepository) BatchUpdate(teams []team.Team) error {
	var insertQueryParts []string

	for _, entity := range teams {
		currentQueryPart := fmt.Sprintf("('%s', '%s')", entity.TitleShort(), entity.TitleFull())
		insertQueryParts = append(insertQueryParts, currentQueryPart)
	}

	queryPattern := "insert into team (title_short, title_full) " +
		"values %s " +
		"on conflict (title_short, title_full) do update " +
		"set title_full = excluded.title_full " +
		"returning id"

	queryToExecute := fmt.Sprintf(queryPattern, strings.Join(insertQueryParts, ","))
	rows, err := w.connection.Query(queryToExecute)

	if err != nil {
		return err
	}
	defer rows.Close()

	var id int
	i := 0
	for rows.Next() {
		_ = rows.Scan(&id)
		teams[i].SetId(id)
		i++
	}

	return nil
}

func NewWriteRepository(connection *sql.DB) team.WriteRepository {
	if repository == nil {
		repository = &writeRepository{connection: connection}
	}

	return repository
}
