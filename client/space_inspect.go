package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectSpace retrieves details about a specific space.
func (client *Client) InspectSpace(name string) (*Space, error) {
	url := client.buildURL("spaces", "space/%s", name)
	request := client.buildInspectSpaceRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectSpaceResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectSpaceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectSpaceResult(response *rest.Response) (*Space, error) {
	result := Space{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
