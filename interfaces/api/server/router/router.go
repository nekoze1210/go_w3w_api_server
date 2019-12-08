package router

import (
	"github.com/labstack/echo/v4"
	h "github.com/nekoze1210/echo-api-w3w/interfaces/api/server/handler"
)

func NewRouter(e *echo.Echo) {
	e.GET("/geocoding/:three_word_address", h.HandleGocoding)
	e.GET("/reverse_geocoding/:coordinates", h.HandleReverseGeocoding)
}
