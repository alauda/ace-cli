package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectProject retrieves details about a specific project.
func (client *Client) InspectProject(name string) (*Project, error) {
	url := client.buildURL("v1", "projects", name)
	request := client.buildInspectProjectRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectProjectResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectProjectRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectProjectResult(response *rest.Response) (*Project, error) {
	result := Project{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
