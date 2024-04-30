package main

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/tashanemclean/calendara-rest-api-api/router"
	"github.com/tashanemclean/calendara-rest-api-api/util/config"
	"github.com/tashanemclean/calendara-rest-api-api/util/logger"
)

func init() {
	config.Load()
	logger.SetupLogger()
}

func main() {
	e := echo.New()
	router.RegisterRoutes(e)
	port := config.Config.AppPort
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
	slog.Info("Server Started")
}