package output

import (
	"github.com/fraqtop/footballcore/competition"
)

type CompetitionListFormatter struct {
	responses []competitionResponse
}

func (clf CompetitionListFormatter) Format(competitions []competition.Competition) []competitionResponse {
	clf.responses = make([]competitionResponse, 0, len(competitions))
	for i := 0; i < len(competitions); i++ {
		clf.responses = append(clf.responses, competitionResponse{
			Id:    competitions[i].Id(),
			Title: competitions[i].Title(),
		})
	}

	return clf.responses
}
