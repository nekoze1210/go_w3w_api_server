package usecase

import (
	w3w "github.com/nekoze1210/echo-api-w3w/infrastructure/what3words"
	"github.com/nekoze1210/echo-api-w3w/infrastructure/what3words/reversegeocoding"
)

func ReverseGeocoding(coordinatesString string) (*w3w.Response, error) {
	r := reversegeocoding.Request{
		Coordinates: coordinatesString,
		Language:    "ja",
		Format:      "json",
	}
	result, err := r.Execute()
	if err != nil {
		return nil, err
	}

	return result, nil
}
