package main

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/tashanemclean/genai-rest-api/router"
	"github.com/tashanemclean/genai-rest-api/util/config"
	"github.com/tashanemclean/genai-rest-api/util/logger"
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