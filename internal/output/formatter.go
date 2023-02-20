package output

import (
	"github.com/fraqtop/footballcore/competition"
)

type CompetitionListFormatter struct {
	responses []competitionResponse
}

func (this CompetitionListFormatter) Format(competitions []competition.Competition) []competitionResponse {
	var response competitionResponse
	for i := 0; i < len(competitions); i++ {
		response = competitionResponse{
			Id:    competitions[i].Id(),
			Title: competitions[i].Title(),
		}
		this.responses = append(this.responses, response)
	}

	return this.responses
}
