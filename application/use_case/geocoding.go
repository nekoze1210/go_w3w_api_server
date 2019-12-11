package usecase

import (
	w3w "github.com/nekoze1210/echo-api-w3w/infrastructure/what3words"
	"github.com/nekoze1210/echo-api-w3w/infrastructure/what3words/geocoding"
)

func Geocoding(r *geocoding.Request) (*w3w.Response, error) {
	result, err := r.Execute()
	if err != nil {
		return nil, err
	}

	return result, nil
}
