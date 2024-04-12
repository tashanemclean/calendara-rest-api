package handlers 

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type defaultMessage struct {
	App 	string `json:"app"`
	Version string `json:"version"`
	Message string `json:"message"`
}

var responseMessage = defaultMessage{App: "GenAI Rest API", Version: "0.1", Message: "OK"}

func Default(c echo.Context) error {
	return c.JSON(http.StatusOK, responseMessage)
}