package server

import (
	"context"
	"fmt"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/labstack/echo/v4"
)

var serverInstance *echo.Echo

func Serve() error {
	serverConfig := config.GetServerConfig()
	serverInstance = echo.New()
	registerRoutes(serverInstance, serverConfig)
	if err := serverInstance.Start(fmt.Sprintf(":%s", serverConfig.Port())); err != nil {
		return err
	}

	return nil
}

func Destroy(ctx context.Context) error {
	err := serverInstance.Shutdown(ctx)

	return err
}
