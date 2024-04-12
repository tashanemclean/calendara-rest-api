package handlers

import (
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"github.com/tashanemclean/genai-rest-api/args"
	"github.com/tashanemclean/genai-rest-api/internal/interactor"
)

func ClassifyText(c echo.Context) error {
	params, err := GetQueryParams[args.ClassifyText](&c)
	if err != nil {
		log.Error("Classify Text GetQueryParams error", err)
		return err
	}

	result := interactor.ClassifyText(interactor.ClassifyTextArgs{
		Text: params.Text,
	}).Execute()

	if result.IsError() {
		log.Error("ClassifyText error", result.AsError())
		return echo.NewHTTPError(http.StatusInternalServerError, result.AsError())
	}

	log.Info("ClassifyText request", err)

	return c.JSON(http.StatusOK, result)
}