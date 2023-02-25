package team

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/fraqtop/footballcore/team"
	"strings"
)

var repository *writeRepository

type writeRepository struct {
	connection *sql.DB
}

var (
	_              team.WriteRepository = (*writeRepository)(nil)
	ErrInvalidTeam                      = errors.New("team is invalid and was not saved")
)

func (this writeRepository) BatchUpdate(teams []team.Team) error {
	inputTeamsLen := len(teams)
	teams = this.filterValid(teams)
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
	rows, err := this.connection.Query(queryToExecute)

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

	if inputTeamsLen != len(teams) {
		return ErrInvalidTeam
	}

	return nil
}

func (this writeRepository) filterValid(teams []team.Team) []team.Team {
	result := make([]team.Team, 0, len(teams))

	for _, entity := range teams {
		if entity.IsValid() {
			result = append(result, entity)
		}
	}

	return result
}

func NewWriteRepository(connection *sql.DB) team.WriteRepository {
	if repository == nil {
		repository = &writeRepository{connection: connection}
	}

	return repository
}
