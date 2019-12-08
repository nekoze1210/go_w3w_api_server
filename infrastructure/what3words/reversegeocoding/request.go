package reversegeocoding

import (
	"net/url"

	w3w "github.com/nekoze1210/echo-api-w3w/infrastructure/what3words"
)

const spath = "/convert-to-3wa"

type Request struct {
	Coordinates string `json:"key"`
	Language    string `json:"language"`
	Format      string `json:"format"`
}

func (r *Request) Execute() (*w3w.Response, error) {
	client, err := w3w.NewClient()
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Set("coordinates", r.Coordinates)
	v.Set("language", r.Language)
	v.Set("Format", r.Format)

	request, err := w3w.NewGetRequest(client, v, spath)
	if err != nil {
		return nil, err
	}

	res, err := client.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	var response w3w.Response
	if err := w3w.DecodeBody(res, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
