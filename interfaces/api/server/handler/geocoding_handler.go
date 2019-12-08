package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/nekoze1210/echo-api-w3w/application/use_case"
)

func HandleGocoding(c echo.Context) error {
	word := c.Param("three_word_address")
	result, err := usecase.Geocoding(word)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}
