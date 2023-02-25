package stats

import (
	"database/sql"
	"errors"
	"fmt"
	corecompetition "github.com/fraqtop/footballcore/competition"
	"github.com/fraqtop/footballcore/stats"
	coreteam "github.com/fraqtop/footballcore/team"
	"strings"
)

type writeRepository struct {
	connection                 *sql.DB
	competitionWriteRepository corecompetition.WriteRepository
	teamWriteRepository        coreteam.WriteRepository
}

var (
	_               stats.WriteRepository = (*writeRepository)(nil)
	ErrStatsInvalid                       = errors.New("stats are invalid and were not saved")
)

func (this writeRepository) BatchUpdate(stats []stats.Stats) error {
	var insertQueryParts []string
	inputStatsLen := len(stats)
	stats = this.filterValid(stats)
	if err := this.syncTeams(stats); err != nil {
		return err
	}
	distinctCompetitions := make(map[int]corecompetition.Competition)
	for _, currentStats := range stats {
		distinctCompetitions[currentStats.Competition().Id()] = currentStats.Competition()
		insertQueryParts = append(
			insertQueryParts,
			fmt.Sprintf(
				"(%d, %d, '%s', %d, %d, %d, %d, %d, %d, %d)",
				currentStats.Team().Id(),
				currentStats.Competition().Id(),
				currentStats.Season(),
				currentStats.Games(),
				currentStats.Points(),
				currentStats.Wins(),
				currentStats.Draws(),
				currentStats.Losses(),
				currentStats.Scored(),
				currentStats.Passed(),
			),
		)
	}

	if err := this.syncCompetitions(distinctCompetitions); err != nil {
		return err
	}

	sqlPattern := "insert into stats (team_id, competition_id, season, games, points, wins, draws, losses, scored, passed)" +
		"values %s " +
		"on conflict(team_id, competition_id) do update set " +
		"season = excluded.season," +
		"games = excluded.games," +
		"points = excluded.points," +
		"wins = excluded.wins," +
		"draws = excluded.draws," +
		"losses = excluded.losses," +
		"scored = excluded.scored," +
		"passed = excluded.passed;"

	queryToExecute := fmt.Sprintf(sqlPattern, strings.Join(insertQueryParts, ","))
	_, err := this.connection.Exec(queryToExecute)
	if err != nil {
		return err
	}

	if inputStatsLen != len(stats) {
		err = ErrStatsInvalid
	}

	return err
}

func (this writeRepository) syncTeams(stats []stats.Stats) error {
	distinctTeams := make(map[string]coreteam.Team)
	for _, currentStats := range stats {
		distinctTeams[currentStats.Team().TitleShort()+currentStats.Team().TitleFull()] = currentStats.Team()
	}

	var teams []coreteam.Team
	for _, currentTeam := range distinctTeams {
		teams = append(teams, currentTeam)
	}

	return this.teamWriteRepository.BatchUpdate(teams)
}

func (this writeRepository) syncCompetitions(competitions map[int]corecompetition.Competition) error {
	for _, currentCompetition := range competitions {
		if err := this.competitionWriteRepository.Save(currentCompetition); err != nil {
			return err
		}
	}

	return nil
}

func (this writeRepository) filterValid(inputStats []stats.Stats) []stats.Stats {
	result := make([]stats.Stats, 0, len(inputStats))

	for _, entity := range inputStats {
		if entity.IsValid() {
			result = append(result, entity)
		}
	}

	return result
}

func NewWriteRepository(
	teamRepository coreteam.WriteRepository,
	competitionRepository corecompetition.WriteRepository,
	connection *sql.DB,
) stats.WriteRepository {
	return &writeRepository{
		connection:                 connection,
		competitionWriteRepository: competitionRepository,
		teamWriteRepository:        teamRepository,
	}
}
