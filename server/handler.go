package server

import (
	"github.com/fraqtop/footballapi/core/competition"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CompetitionListHandler(ctx echo.Context) error {
	competitionRepository, err := competition.NewReadRepository()
	if err != nil {
		return err
	}

	var (
		responses []competitionResponse
		response  competitionResponse
	)

	for _, competitionEntity := range competitionRepository.All() {
		response = competitionResponse{
			Id:    competitionEntity.Id(),
			Title: competitionEntity.Title(),
		}
		responses = append(responses, response)
	}

	return ctx.JSON(http.StatusOK, responses)
}
