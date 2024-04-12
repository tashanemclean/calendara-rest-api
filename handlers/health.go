package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthMessage struct {
	App 	string `json:"app"`
	Version string `json:"version"`
	Message string `json:"message"`
}

var healthmessage = healthMessage{App: "App Name", Version: "0.1", Message: "OK"}

func Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, healthmessage)
}