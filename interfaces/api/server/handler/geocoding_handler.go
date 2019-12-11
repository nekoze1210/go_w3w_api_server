package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/nekoze1210/echo-api-w3w/application/use_case"
	"github.com/nekoze1210/echo-api-w3w/infrastructure/what3words/geocoding"
)

func HandleGocoding(c echo.Context) error {
	word := c.Param("three_word_address")
	r := geocoding.Request{
		Words:    word,
		Language: "ja",
		Format:   "json",
	}
	result, err := usecase.Geocoding(&r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}
