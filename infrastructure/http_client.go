package infrastructure

import (
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

func NewClient(rawurl string) (*Client, error) {
	parsedurl, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse ur: %s", rawurl)
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	return &Client{
		URL:        parsedurl,
		HTTPClient: client,
	}, nil
}
