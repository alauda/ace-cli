package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListProjectsResult defines the response body for the ListProjects API.
type ListProjectsResult struct {
	Projects []Project
}

// ListProjects returns all projects in an account.
func (client *Client) ListProjects() (*ListProjectsResult, error) {
	url := client.buildURL("v1", "projects", "")
	request := client.buildListProjectsRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListProjectsResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListProjectsRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListProjectsResult(response *rest.Response) (*ListProjectsResult, error) {
	result := ListProjectsResult{}

	err := json.Unmarshal(response.Body(), &result.Projects)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
