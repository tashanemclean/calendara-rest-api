package router

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"

	"github.com/tashanemclean/calendara-rest-api-api/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	// TODO config validatior for server

	// Health check route
	e.GET("/health", handlers.Healthcheck)
	e.GET("status", echoprometheus.NewHandler())
	e.GET("/v1", handlers.Default)

	v1 := e.Group("/v1")

	// Classification routes
	v1.POST("/processText", handlers.PromptText)
}