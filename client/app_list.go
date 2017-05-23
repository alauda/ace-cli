package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// ListAppsParams defines the query parameters for the ListApps API.
type ListAppsParams struct {
	Cluster string
}

// ListAppsResult defines the response body for the ListApps API.
type ListAppsResult struct {
	Apps []App
}

// ListApps returns all apps in a cluster.
func (client *Client) ListApps(params *ListAppsParams) (*ListAppsResult, error) {
	url := client.buildListAppsURL()
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

func (client *Client) buildListAppsURL() string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/applications/%s", server, client.Namespace())
}

func (client *Client) buildListAppsRequest(params *ListAppsParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.Cluster != "" {
		request.SetQueryParam("region", params.Cluster)
	}

	return request
}

func parseListAppsResult(response *rest.Response) (*ListAppsResult, error) {
	result := ListAppsResult{}

	err := json.Unmarshal(response.Body(), &result.Apps)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
