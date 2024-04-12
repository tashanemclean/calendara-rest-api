package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tashanemclean/genai-rest-api/args"
)

func GetQueryParams[T args.Args](c *echo.Context) (*T, error) {
	params := new(T)
	err := (*c).Bind(params)
	if err != nil {
		return nil, err
	}

	err = (&echo.DefaultBinder{}).BindHeaders(*c, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}