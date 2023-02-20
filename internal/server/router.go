package server

import (
	"github.com/labstack/echo/v4"
)

func registerRoutes(serverInstance *echo.Echo) {
	serverInstance.GET("/competition", competitionListHandler)
}
