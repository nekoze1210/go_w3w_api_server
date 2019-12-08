package what3words

import (
	"net/http"
	"net/url"
	"os"
	"path"

	infra "github.com/nekoze1210/echo-api-w3w/infrastructure"
	errors "github.com/pkg/errors"
)

const apiEndpoint = "https://api.what3words.com/v3"

func NewGetRequest(c *infra.Client, v url.Values, spath string) (*http.Request, error) {
	apiKey := os.Getenv("W3W_API_KEY")
	url := *c.URL
	url.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize HTTP Request")
	}

	req.Header.Set("Content-Type", "Application/json;charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)

	req.URL.RawQuery = v.Encode()

	return req, nil
}
