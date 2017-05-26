package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectApp retrieves details about a specific application.
func (client *Client) InspectApp(name string) (*App, error) {
	url := client.buildURL("applications", name)
	request := client.buildInspectAppRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectAppResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectAppResult(response *rest.Response) (*App, error) {
	result := App{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
