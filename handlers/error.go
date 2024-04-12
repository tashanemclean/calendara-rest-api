package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrInvalidArgs = echo.NewHTTPError(http.StatusBadRequest, "invalid args")