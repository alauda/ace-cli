package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListConfigsResult defines the response body for the ListConfigs API.
type ListConfigsResult struct {
	Count   int      `json:"count"`
	Configs []Config `json:"results"`
}

// ListConfigs returns all configurations.
func (client *Client) ListConfigs() (*ListConfigsResult, error) {
	url := client.buildURL("configs", "")
	request := client.buildListConfigsRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListConfigsResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListConfigsRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListConfigsResult(response *rest.Response) (*ListConfigsResult, error) {
	result := ListConfigsResult{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
