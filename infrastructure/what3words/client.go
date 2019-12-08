package what3words

import infra "github.com/nekoze1210/echo-api-w3w/infrastructure"

func NewClient() (*infra.Client, error) {
	const apiEndpoint = "https://api.what3words.com/v3"
	client, err := infra.NewClient(apiEndpoint)
	if err != nil {
		return nil, err
	}
	return client, nil
}
