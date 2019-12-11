package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/nekoze1210/echo-api-w3w/application/use_case"
	"github.com/nekoze1210/echo-api-w3w/infrastructure/what3words/reversegeocoding"
)

func HandleReverseGeocoding(c echo.Context) error {
	coordinates := c.Param("coordinates")
	r := reversegeocoding.Request{
		Coordinates: coordinates,
		Language:    "ja",
		Format:      "json",
	}
	result, err := usecase.ReverseGeocoding(&r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}
