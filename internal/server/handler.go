package server

import (
	"net/http"

	"github.com/fraqtop/footballapi/internal/container"
	"github.com/fraqtop/footballapi/internal/output"
	corecompetition "github.com/fraqtop/footballcore/competition"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func competitionListHandler(ctx echo.Context) error {
	err := container.Get().Invoke(func(repository corecompetition.ReadRepository, formatter *output.CompetitionListFormatter) {
		responses := formatter.Format(repository.All())

		err := ctx.JSON(http.StatusOK, responses)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, output.ErrorResponse{Message: err.Error()})
		}
	})

	if err != nil {
		log.Warn(err)
	}

	return err
}
