package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListRegistryProjectsResult defines the response body for the ListRegistryProjects API.
type ListRegistryProjectsResult struct {
	Projects []RegistryProject
}

// ListRegistryProjects returns the list of projects for a registry.
func (client *Client) ListRegistryProjects(registryName string) (*ListRegistryProjectsResult, error) {
	url := client.buildURL("registries", "%s/projects", registryName)
	request := client.buildListRegistryProjectsRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListRegistryProjectsResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListRegistryProjectsRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListRegistryProjectsResult(response *rest.Response) (*ListRegistryProjectsResult, error) {
	result := ListRegistryProjectsResult{}

	err := json.Unmarshal(response.Body(), &result.Projects)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
