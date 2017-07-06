package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectConfig retrieves details about the specified configuration.
func (client *Client) InspectConfig(name string) (*Config, error) {
	url := client.buildURL("configs", name)
	request := client.buildInspectConfigRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectConfigResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectConfigRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectConfigResult(response *rest.Response) (*Config, error) {
	result := Config{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
