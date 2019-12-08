package usecase

import (
	w3w "github.com/nekoze1210/echo-api-w3w/infrastructure/what3words"
	"github.com/nekoze1210/echo-api-w3w/infrastructure/what3words/geocoding"
)

func Geocoding(word string) (*w3w.Response, error) {
	r := geocoding.Request{
		Words:    word,
		Language: "ja",
		Format:   "json",
	}
	result, err := r.Execute()
	if err != nil {
		return nil, err
	}

	return result, nil
}
