package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/tashanemclean/calendara-rest-api-api/util/config"
	"github.com/tashanemclean/calendara-rest-api-api/util/logger"
)

type RequestHeaders struct {
	Authorization string
}

type CustomContext struct {
	echo.Context
	Headers *RequestHeaders
}

func Register(e *echo.Echo) {
	e.Use(echoprometheus.NewMiddleware("calendaraBackend"))
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*", fmt.Sprintf("http://localhost:%s", config.Config.AppPort)},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values echoMiddleware.RequestLoggerValues) error {
			logger.Info(
				"request",
				"request | logging",
				values.Error,
				"contentlength",
				values.ContentLength,
				"formvalues",
				values.FormValues,
				"host",
				values.Host,
				"method",
				values.Method,
				"uri",
				values.URI,
				"startTime",
				values.StartTime,
				"status",
				values.Status,
				"path",
				values.RoutePath,
			)
			return nil
		},
	}))
}
