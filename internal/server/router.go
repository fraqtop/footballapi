package server

import (
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/pprof"
)

func registerRoutes(serverInstance *echo.Echo, config *config.ServerConfig) {
	serverInstance.GET("/competition", competitionListHandler)

	if config.Debug() {
		serverInstance.GET("/debug", echoRouterAdapter(pprof.Index))
		serverInstance.GET("/debug/cmdline", echoRouterAdapter(pprof.Cmdline))
		serverInstance.GET("/debug/profile", echoRouterAdapter(pprof.Profile))
		serverInstance.GET("/debug/symbol", echoRouterAdapter(pprof.Symbol))
		serverInstance.POST("/debug/symbol", echoRouterAdapter(pprof.Symbol))
		serverInstance.GET("/debug/trace", echoRouterAdapter(pprof.Trace))
	}
}

func echoRouterAdapter(handlerFunc http.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		handlerFunc(c.Response().Writer, c.Request())

		return nil
	}
}
