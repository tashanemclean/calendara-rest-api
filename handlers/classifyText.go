package handlers

import (
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"github.com/tashanemclean/calendara-rest-api-api/args"
	"github.com/tashanemclean/calendara-rest-api-api/internal/interactor"
)

func ClassifyText(c echo.Context) error {
	params, err := GetQueryParams[args.ClassifyText](&c)
	if err != nil {
		log.Error("Classify Text GetQueryParams error", err)
		return err
	}

	result := interactor.ClassifyText(interactor.ClassifyTextArgs{
		Activity: params.Activity,
		Categories: params.Categories,
		City: params.City,
		State: params.State,
		Days: params.Days,
	}).Execute()

	if result.IsError() {
		log.Error("ClassifyText error", result.AsError())
		return echo.NewHTTPError(http.StatusInternalServerError, result.AsError())
	}

	log.Info("ClassifyText request", err)
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	return c.JSON(http.StatusOK, result)
}