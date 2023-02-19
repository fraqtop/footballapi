package server

import (
	"context"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/fraqtop/footballapi/internal/connection"
	"github.com/fraqtop/footballapi/internal/repository/competition"
	"github.com/labstack/echo/v4"
	"net/http"
)

func competitionListHandler(ctx echo.Context) error {
	competitionRepository, err := competition.NewReadRepository()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorResponse{Reason: err.Error()})
	}
	cache := connection.GetRedisClient(config.GetCacheConfig())
	competitionRepository = competition.NewRepositoryCacheProxy(context.Background(), cache, competitionRepository)

	formatter := newCompetitionListFormatter()

	return ctx.JSON(http.StatusOK, formatter.format(competitionRepository.All()))
}
