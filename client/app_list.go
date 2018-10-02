package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListAppsParams defines the query parameters for the ListApps API.
type ListAppsParams struct {
	Project string
}

// ListAppsResult defines the response body for the ListApps API.
type ListAppsResult []App

// ListApps returns all applications in a Kubernetes namespace
func (client *Client) ListApps(cluster string, namespace string, params *ListAppsParams) (*ListAppsResult, error) {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/applications/%s/", cluster, namespace)
	request := client.buildListAppsRequest(params)

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListAppsResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListAppsRequest(params *ListAppsParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	request.SetQueryParam("project_name", params.Project)

	return request
}

func parseListAppsResult(response *rest.Response) (*ListAppsResult, error) {
	result := ListAppsResult{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
