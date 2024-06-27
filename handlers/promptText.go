package handlers

import (
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"github.com/tashanemclean/calendara-rest-api-api/internal/interactor"
)

func PromptText(c echo.Context) error {
	params, err := GetQueryParams(&c)
	if err != nil {
		log.Error("Prompt Text GetQueryParams error", err)
		return err
	}

	result := interactor.PromptText(interactor.PromptTextArgs{
		Activity: params.Activity,
		Categories: params.Categories,
		City: params.City,
		State: params.State,
		Days: params.Days,
	}).Execute()

	if result.IsError() {
		log.Error("PromptText error", result.AsError())
		return echo.NewHTTPError(http.StatusInternalServerError, result.AsError())
	}

	log.Info("PromptText request", err)

	return c.JSON(http.StatusOK, result)
}