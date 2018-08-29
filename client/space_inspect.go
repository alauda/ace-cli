package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectSpaceParams defines the query parameters for the InspectSpace API.
type InspectSpaceParams struct {
	Project string
}

// InspectSpace retrieves details about a specific space.
func (client *Client) InspectSpace(name string, params *InspectSpaceParams) (*Space, error) {
	url := client.buildURL("v1", "spaces", "space/%s", name)
	request := client.buildInspectSpaceRequest(params)

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

func (client *Client) buildInspectSpaceRequest(params *InspectSpaceParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	request.SetQueryParam("project_name", params.Project)

	return request
}

func parseInspectSpaceResult(response *rest.Response) (*Space, error) {
	result := Space{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
