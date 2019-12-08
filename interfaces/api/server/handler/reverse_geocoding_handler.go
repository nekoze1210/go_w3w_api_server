package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/nekoze1210/echo-api-w3w/application/use_case"
)

func HandleReverseGeocoding(c echo.Context) error {
	coordinates := c.Param("coordinates")
	result, err := usecase.ReverseGeocoding(coordinates)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}
